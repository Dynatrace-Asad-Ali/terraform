# dynatrace_alerting_profile Data Source

Use this data source to get information about a specific alerting profile in dynatrace that already exists

## Example Usage

```hcl
data "dynatrace_alerting_profile" "error_alerts" {
   display_name = "easyTravel availability alerts"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The id of the alerting profile.
* `display_name` (Optional) The name of the alerting profile.
