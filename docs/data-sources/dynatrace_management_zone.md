# dynatrace_management_zone Data Source

Use this data source to get information about a specific management zone in dynatrace that already exists

## Example Usage

```hcl
data "dynatrace_management_zone" "prod"{
    name = "Production"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) The id of the management zone.
* `name` (Optional) The name of the management zone.
