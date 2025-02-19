package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nikhilsbhat/terraform-provider-gocd/pkg/client"
)

func init() { //nolint:gochecknoinits
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("GOCD_BASE_URL", "www.gocd.com"),
				Description: "base url of GoCD server, which this terraform provider can interact with",
			},
			"ca_file": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("GOCD_CAFILE_CONTENT", "some_ca_context"),
				Description: "CA file contents, to be used while connecting to GoCD server when CA based auth is enabled",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("GOCD_USERNAME", "username"),
				Description: "username to be used while connecting with GoCD",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("GOCD_PASSWORD", "password"),
				Description: "password to be used while connecting with GoCD",
			},
			"loglevel": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("GOCD_PASSWORD", "password"),
				Description: "loglevel to be set for the api calls made to GoCD",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"gocd_plugin_setting":        resourcePluginsSetting(),
			"gocd_auth_config":           resourceAuthConfig(),
			"gocd_cluster_profile":       resourceClusterProfile(),
			"gocd_elastic_agent_profile": resourceElasticAgentProfile(),
			"gocd_config_repository":     resourceConfigRepository(),
			"gocd_environment":           resourceEnvironment(),
			"gocd_encrypt_value":         resourceEncryptValue(),
			"gocd_secret_config":         resourceSecretConfig(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"gocd_plugin_setting":        dataSourcePluginsSetting(),
			"gocd_auth_config":           dataSourceAuthConfig(),
			"gocd_cluster_profile":       dataSourceClusterProfile(),
			"gocd_elastic_agent_profile": dataSourceElasticAgentProfile(),
			"gocd_config_repository":     dataSourceConfigRepository(),
			"gocd_environment":           dataSourceEnvironment(),
			"gocd_secret_config":         dataSourceSecretConfig(),
		},

		ConfigureContextFunc: client.GetGoCDClient,
	}
}
