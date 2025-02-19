package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nikhilsbhat/gocd-sdk-go"
	"github.com/nikhilsbhat/terraform-provider-gocd/pkg/utils"
)

func dataSourcePluginsSetting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePluginsSettingRead,
		Schema: map[string]*schema.Schema{
			"plugin_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Required:    false,
				Description: "The plugin identifier of the cluster profile.",
			},
			"configuration": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "List of configuration required to configure the plugin settings.",
				Elem:        propertiesSchemaData(),
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Etag used to track the cluster profile",
			},
		},
	}
}

func dataSourcePluginsSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defaultConfig := meta.(gocd.GoCd)

	id := d.Id()

	if len(id) == 0 {
		newID, err := utils.GetRandomID()
		if err != nil {
			d.SetId("")

			return diag.Errorf("errored while fetching randomID %v", err)
		}
		id = newID
	}

	pluginID := utils.String(d.Get(utils.TerraformPluginID))

	response, err := defaultConfig.GetPluginSettings(pluginID)
	if err != nil {
		return diag.Errorf("getting cluster profile %s errored with: %v", pluginID, err)
	}

	flattenedConfiguration, err := utils.MapSlice(response.Configuration)
	if err != nil {
		d.SetId("")

		return diag.Errorf("errored while flattening Configuration obtained: %v", err)
	}

	if err = d.Set(utils.TerraformResourceConfiguration, flattenedConfiguration); err != nil {
		return diag.Errorf(settingAttrErrorTmp, err, utils.TerraformResourceConfiguration)
	}

	d.SetId(id)

	return nil
}
