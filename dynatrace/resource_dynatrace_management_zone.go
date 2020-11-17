package dynatrace

import (
	"context"

	"github.com/antihax/optional"
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDynatraceManagementZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceManagementZoneCreate,
		ReadContext:   resourceDynatraceManagementZoneRead,
		UpdateContext: resourceDynatraceManagementZoneUpdate,
		DeleteContext: resourceDynatraceManagementZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the management zone.",
				Required:    true,
			},
			"rule": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of rules for management zone usage. Each rule is evaluated independently of all other rules.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The type of Dynatrace entities the management zone can be applied to.",
							Required:    true,
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "The rule is enabled (true) or disabled (false).",
							Required:    true,
						},
						"propagation_types": &schema.Schema{
							Type:        schema.TypeSet,
							Description: "How to apply the management zone to underlying entities.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"condition": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of matching rules for the management zone. The management zone applies only if all conditions are fulfilled.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The key to identify the data we're matching.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The attribute to be used for comparision.",
													Required:    true,
												},
												"dynamic_key": {
													Type:        schema.TypeString,
													Description: "Dynamic key generated based on selected type/attribute.",
													Optional:    true,
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Defines the actual set of fields depending on the value.",
													Optional:    true,
												},
											},
										},
									},
									"comparison_info": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Defines how the matching is actually performed: what and how are we comparing.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operator": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Operator of the comparison. You can reverse it by setting negate to true. Possible values depend on the type of the comparison. Find the list of actual models in the description of the type field and check the description of the model you need.",
													Required:    true,
												},
												"value": {
													Type:         schema.TypeString,
													Description:  "The value to compare to.",
													Optional:     true,
													ValidateFunc: validation.StringIsJSON,
												},
												"negate": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Reverses the comparison operator. For example it turns the begins with into does not begin with.",
													Required:    true,
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Defines the actual set of fields depending on the value.",
													Required:    true,
												},
												"case_sensitive": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Defines if value to compare to is case sensitive",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceDynatraceManagementZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	mz := dynatraceConfigV1.ManagementZone{
		Name:  d.Get("name").(string),
		Rules: expandManagementZoneRules(d.Get("rule").([]interface{})),
	}

	mzBody := dynatraceConfigV1.CreateManagementZoneOpts{
		ManagementZone: optional.NewInterface(mz),
	}

	managementZone, _, err := dynatraceConfigClientV1.ManagementZonesApi.CreateManagementZone(authConfigV1, &mzBody)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace client",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(managementZone.Id)

	resourceDynatraceManagementZoneRead(ctx, d, m)

	return diags
}

func resourceDynatraceManagementZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	managementZone, _, err := dynatraceConfigClientV1.ManagementZonesApi.GetSingleManagementZoneConfig(authConfigV1, managementZoneID, &dynatraceConfigV1.GetSingleManagementZoneConfigOpts{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace client",
			Detail:   err.Error(),
		})
		return diags
	}

	managementZoneRules := flattenManagementZoneRulesData(&managementZone.Rules)
	if err := d.Set("rule", managementZoneRules); err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", &managementZone.Name)

	return diags
}

func resourceDynatraceManagementZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	if d.HasChange("name") || d.HasChange("rule") {

		mz := dynatraceConfigV1.ManagementZone{
			Name:  d.Get("name").(string),
			Rules: expandManagementZoneRules(d.Get("rule").([]interface{})),
		}

		mzBody := dynatraceConfigV1.CreateOrUpdateManagementZoneOpts{
			ManagementZone: optional.NewInterface(mz),
		}

		_, _, err := dynatraceConfigClientV1.ManagementZonesApi.CreateOrUpdateManagementZone(authConfigV1, managementZoneID, &mzBody)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create dynatrace client",
				Detail:   err.Error(),
			})
			return diags
		}

	}

	return resourceDynatraceManagementZoneRead(ctx, d, m)
}

func resourceDynatraceManagementZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	managementZoneID := d.Id()

	_, err := dynatraceConfigClientV1.ManagementZonesApi.DeleteManagementZone(authConfigV1, managementZoneID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace client",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
