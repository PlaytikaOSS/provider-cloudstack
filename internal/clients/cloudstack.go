/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/playtika/provider-cloudstack/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal cloudstack credentials as JSON"
	keyAPIKey               = "api_key"
	keySecretKey            = "secret_key"
	keyAPIURL               = "api_url"
	timeout                 = "timeout"
	httpGetOnly             = "http_get_only"
	configKey               = "config"
	profile                 = "profile"
)

func buildProviderConfig(creds map[string]string) map[string]any {
	config := map[string]any{}

	if value, ok := creds[keyAPIKey]; ok {
		config[keyAPIKey] = value
	}
	if value, ok := creds[keySecretKey]; ok {
		config[keySecretKey] = value
	}
	if value, ok := creds[keyAPIURL]; ok {
		config[keyAPIURL] = value
	}
	if value, ok := creds[timeout]; ok {
		config[timeout] = value
	}
	if value, ok := creds[httpGetOnly]; ok {
		config[httpGetOnly] = value
	}
	if value, ok := creds[configKey]; ok {
		config[configKey] = value
	}
	if value, ok := creds[profile]; ok {
		config[profile] = value
	}

	return config
}

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
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		
		ps.Configuration = buildProviderConfig(creds)
		// Set credentials in Terraform provider configuration.
		/*ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}*/
		return ps, nil
	}
}
