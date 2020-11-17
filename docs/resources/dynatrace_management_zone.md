# dynatrace_management_zone Resource

Provides a dynatrace management resource. It allows to create, update, delete management zones in a dynatrace environment. [Management Zones API]

## Example Usage

```hcl

resource "dynatrace_management_zone" "sockshop_prod" {

  name = "sockshop_prod"
  rule{
    type = "SERVICE"
    enabled = true
    propagation_types = [
      "SERVICE_TO_HOST_LIKE",
      "SERVICE_TO_PROCESS_GROUP_LIKE"
    ]
    condition {
      key {
        attribute = "SERVICE_TAGS"
      }
      comparison_info {
        type = "TAG"
        operator = "EQUALS"
        value = jsonencode(
            {
                context = "CONTEXTLESS"
                key     = "env"
                value   = "prod"
            }
        )
        negate = false
      }
    }

    condition {
      key {
        attribute = "HOST_GROUP_NAME"
      }
      comparison_info {
        type = "STRING"
        operator = "BEGINS_WITH"
        value = jsonencode("simpleapp")
        negate = false
        case_sensitive = false
      }
    }
  }
}

```

## Argument Reference

* `name` - (Required) The name of the management zone.
* `rule` - (Optional) Rules for management zone usage. Each rule is evaluated independently of all other rules.

## Attribute Reference

* `id` - The ID of the management zone.

## Nested Blocks

`rule`

### Arguments

* `type` - (Required) The type of Dynatrace entities the management zone can be applied to.
* `enabled` - (Required) The rule is enabled (true) or disabled (false).
* `propagation_types` - (Optional) How to apply the management zone to underlying entities.
* `condition` - (Required) A list of matching rules for the management zone. The management zone applies only if all conditions are fulfilled.

`condition`

### Arguments

* `key` - (Required) The key to identify the data we're matching.
* `comparison_info` (Required) Defines how the matching is actually performed: what and how are we comparing.

`key`

### Arguments

* `attribute` - (Required) The attribute to be used for comparison.
* `type` - (Optional) Defines the actual set of fields depending on the value.
* `dynamic_key` - (Optional) Dynamic key generated based on selected type/attribute.

`comparison_info`

### Arguments

* `operator` - (Required) Operator of the comparison. You can reverse it by setting negate to true. Possible values depend on the type of the comparison. Find the list of actual models in the description of the type field and check the description of the model you need.
* `value` - (Optional) The value to compare to as string containing plain JSON-encoded data.
* `negate` - (Required) Reverses the comparison operator. For example it turns the begins with into does not begin with.
* `type` - (Required) Defines the actual set of fields depending on the value.
* `case_sensitive` - (Optional) Defines if value to compare to is case sensitive

### Arguments

## Import

Dynatrace management zones can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_management_zone.keptn_carts -4638826838889583423
```

[Management Zones API]: (https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/management-zones-api/)
