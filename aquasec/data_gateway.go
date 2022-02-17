package aquasec

import (
	"context"
	"log"

	"github.com/aquasecurity/terraform-provider-aquasec/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGateways() *schema.Resource {
	return &schema.Resource{
		Description: "The data source `aquasec_gateways` provides a method to query all gateways within the Aqua ",
		ReadContext: dataGatewayRead,
		Schema: map[string]*schema.Schema{
			"gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logical_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssh_add": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"grpc_add": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataGatewayRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG]  inside dataGateway")
	c := m.(*client.Client)
	result, err := c.GetGateways()
	if err == nil {
		gateways, id := flattenGatewaysData(&result)
		d.SetId(id)
		if err := d.Set("gateways", gateways); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diag.FromErr(err)
	}

	return nil
}
