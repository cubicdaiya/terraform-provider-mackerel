package mackerel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mackerelio/mackerel-client-go"
)

var dashboardRangeDataResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"relative": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"period": {
						Type:     schema.TypeInt,
						Computed: true,
					},
					"offset": {
						Type:     schema.TypeInt,
						Computed: true,
					},
				},
			},
		},
		"absolute": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"start": {
						Type:     schema.TypeInt,
						Computed: true,
					},
					"end": {
						Type:     schema.TypeInt,
						Computed: true,
					},
				},
			},
		},
	},
}

var dashboardLayoutDataResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"x": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"y": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"width": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"height": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	},
}

var dashboardMetricDataResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"host": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"host_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"service": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"service_name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"expression": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"expression": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"query": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"query": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"legend": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	},
}

func dataSourceMackerelDashboard() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMackerelDashboardRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memo": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"graph": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"role": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"role_fullname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_stacked": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"service": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"expression": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"expression": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"query": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"legend": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"range": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardRangeDataResource,
						},
						"layout": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardLayoutDataResource,
						},
					},
				},
			},
			"value": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardMetricDataResource,
						},
						"fraction_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"suffix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"layout": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardLayoutDataResource,
						},
					},
				},
			},
			"markdown": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"markdown": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"layout": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardLayoutDataResource,
						},
					},
				},
			},
			"alert_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_fullname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"layout": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     dashboardLayoutDataResource,
						},
					},
				},
			},
		},
	}
}

func dataSourceMackerelDashboardRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Get("id").(string)

	client := m.(*mackerel.Client)
	dashboard, err := client.FindDashboard(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dashboard.ID)

	return flattenDashboard(dashboard, d)
}
