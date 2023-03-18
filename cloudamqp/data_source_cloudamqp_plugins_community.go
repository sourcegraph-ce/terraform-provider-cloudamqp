package cloudamqp

import (
	"fmt"
	log "github.com/sourcegraph-ce/logrus"

	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourcePluginsCommunity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePluginsCommunityRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Instance identifier",
			},
			"plugins": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"require": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePluginsCommunityRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*api.API)
	log.Printf("[DEBUG] cloudamqp::data_source::plugin_comminity::read instance id: %v", d.Get("instance_id"))
	data, err := api.ReadPluginsCommunity(d.Get("instance_id").(int))
	d.SetId(fmt.Sprintf("%v.plugins_community", d.Get("instance_id").(int)))
	if err != nil {
		return err
	}
	if err = d.Set("plugins", data); err != nil {
		return fmt.Errorf("error setting community plugins for resource %s: %s", d.Id(), err)
	}
	return nil
}
