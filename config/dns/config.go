package dns

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Customize configures individual resources by adding custom ResourceConfigurators.
func Customize(p *config.Provider) {
	p.AddResourceConfigurator("record_id", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
