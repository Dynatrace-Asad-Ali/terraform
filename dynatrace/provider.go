package dynatrace

import (
	"context"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function for Dynatrace API
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dt_env_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_ENV_URL", nil),
			},
			"dt_api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DYNATRACE_API_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dynatrace_alerting_profile":   resourceDynatraceAlertingProfile(),
			"dynatrace_management_zone":    resourceDynatraceManagementZone(),
			"dynatrace_maintenance_window": resourceDynatraceMaintenanceWindow(),
			"dynatrace_dashboard":          resourceDynatraceDashboard(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dynatrace_alerting_profile": dataSourceDynatraceAlertingProfile(),
			"dynatrace_management_zone":  dataSourceDynatraceManagementZone(),
		},
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return providerConfigure(ctx, d, p.TerraformVersion)
	}

	return p
}

// ProviderConfiguration contains the initialized API clients to communicate with the Datadog API
type ProviderConfiguration struct {
	DynatraceConfigClientV1 *dynatraceConfigV1.APIClient
	AuthConfigV1            context.Context
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	dtEnvURL := d.Get("dt_env_url").(string)
	apiToken := d.Get("dt_api_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Initialize the official Dynatrace Config V1 API client
	authConfigV1 := context.WithValue(
		context.Background(),
		dynatraceConfigV1.ContextAPIKey,
		dynatraceConfigV1.APIKey{
			Prefix: "Api-token",
			Key:    apiToken,
		},
	)

	configV1 := dynatraceConfigV1.NewConfiguration()
	configV1.BasePath = dtEnvURL + "/api/config/v1"

	dynatraceConfigClientV1 := dynatraceConfigV1.NewAPIClient(configV1)

	return &ProviderConfiguration{
		DynatraceConfigClientV1: dynatraceConfigClientV1,
		AuthConfigV1:            authConfigV1,
	}, diags

}
