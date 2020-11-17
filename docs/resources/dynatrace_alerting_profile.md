# Resource: dynatrace_alerting_profile

Provides a dynatrace alerting profile resource. It allows to create, update, delete alerting profiles in a dynatrace environment. [Alerting profiles API]

## Example Usage

```hcl
resource "dynatrace_alerting_profile" "sockshop_errors" {

  display_name = "sockshop_errors"
  mz_id = dynatrace_management_zone.sockshop_prod.id

  rule{
    severity_level = "AVAILABILITY"
    tag_filters {
      include_mode = "INCLUDE_ALL"
      tag_filter {
        context = "CONTEXTLESS"
        key = "app"
        value = "carts"
      }
      tag_filter {
        context = "CONTEXTLESS"
        key = "env"
        value = "prod"
      }
    }
    delay_in_minutes = 2
  }

  event_type_filter{
    predefined_event_filter{
      negate = true
      event_type = "EC2_HIGH_CPU"
    }
  }

  event_type_filter{
    custom_event_filter{
      custom_title_filter{
        enabled = true
        value = "sockshop"
        operator = "CONTAINS"
        negate = true
        case_insensitive = false
      }
    }
  }
}
```

## Argument Reference

* `display_name` - (Required) The name of the alerting profile, displayed in the UI.
* `mz_id` - (Optional) The ID of the management zone to which the alerting profile applies.
* `rule` - (Optional) Contains a list of severity rules. The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops. If you specify both severity rule and event filter, the AND logic applies.
* `event_type_filter` - (Optional) Describes the configuration of the event filter for the alerting profile.

## Attribute Reference

* `id` - The ID of the alerting profile.

## Nested blocks

`rule`

### Arguments

* `severity_level` - (Required) The severity level to trigger the alert.
* `tag_filters` - (Required) Configuration of the tag filtering of the alerting profile.
* `delay_in_minutes` - (Required) Send a notification if a problem remains open longer than X minutes.

`tag_filters`

### Arguments

* `include_mode` - (Optional) The filtering mode.
* `tag_filter` - (Optional) A tag-based filter of monitored entities.

`tag_filter`

### Arguments

* `context` - (Required) The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the CONTEXTLESS value.
* `key` - (Required) The key of the tag. Custom tags have the tag value here.
* `value` - (Required) The value of the tag. Not applicable to custom tags.

`event_rype_filter`

### Arguments

* `predefined_event_filter` - (Optional) Configuration of a predefined event filter.
* `custom_event_filter` - (Optional) Configuration of a custom event filter. See nested custom event filter below for details.

`predefined_event_filter`

### Arguments

* `event_type` - (Required) The type of the predefined event.
* `negate` - (Required) The alert triggers when the problem of specified severity arises while the specified event is happening (false) or while the specified event is not happening (true).

`custom_event_filter`

### Arguments

* `custom_title_filter` - (Optional) Configuration of a matching filter.
* `custom_description_filter` - (Optional) Configuration of a matching filter.

`custom_title_filter` / `custom_description_filter`

### Arguments

* `enabled` - (Required) - The filter is enabled (true) or disabled (false).
* `value` - (Required) The value to compare to.
* `operator` - (Required) Operator of the comparison. You can reverse it by setting negate to true.
* `negate` - (Required) Reverses the comparison operator. For example it turns the begins with into does not begin with.
* `case_insensitive` - (Optional) The condition is case sensitive (false) or case insensitive (true). If not set, then false is used, making the condition case sensitive.

## Import

Dynatrace alerting profiles can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_alerting_profile.keptn dc228252-2b3d-43ec-b6c5-7bd231adeb6e
```

[Alerting profiles API]: (https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/alerting-profiles-api/post-profile/)
