/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/maltejk/provider-jet-myrasec/apis/v1alpha1"
)

const (
	apiKey     = "api_key"
	apiSecret  = "secret"
	apiBaseURL = "api_base_url"

	// MyraSec credentials environment variable names
	envAPIKey    = "MYRASEC_API_KEY"
	envAPISecret = "MYRASEC_API_SECRET"
)

const (
	fmtEnvVar = "%s=%s"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal myrasec credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		myrasecCreds := map[string]string{}
		if err := json.Unmarshal(data, &myrasecCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// provide api_base_url, as it's not accepted from env var
		// https://github.com/Myra-Security-GmbH/terraform-provider-myrasec/blob/v1.20.0/myrasec/provider.go#L34
		// untils this https://github.com/Myra-Security-GmbH/terraform-provider-myrasec/issues/28
		// gets fixed
		if myrasecCreds[apiBaseURL] != "" {
			tfCfg := map[string]interface{}{}
			tfCfg["api_base_url"] = myrasecCreds[apiBaseURL]
			ps.Configuration = tfCfg
		}

		// set environment variables for sensitive provider configuration
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envAPIKey, myrasecCreds[apiKey]),
			fmt.Sprintf(fmtEnvVar, envAPISecret, myrasecCreds[apiSecret]),
		}
		return ps, nil
	}
}
