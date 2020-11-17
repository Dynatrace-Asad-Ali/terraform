# Resource: dynatrace_maintenance_window

Provides a dynatrace management resource. It allows to create, update, delete maintenance windows in a dynatrace environment. [Maintenance Windows API]

## Example Usage

```hcl
resource "dynatrace_maintenance_window" "windows_patching" {
  name = "windows patches"
  description = "Weekly udpate of windows servers"
  type = "PLANNED" 
  suppression = "DETECT_PROBLEMS_DONT_ALERT"
  scope {
    match {
      type = "HOST"
      tags {
        context = "CONTEXTLESS"
        key = "OS"  
        value = "windows"
      }
    }
    match {
      type = "HOST"
      tags {
        context = "CONTEXTLESS"
        key = "OS"  
        value = "WIN32"
      }
    }
  }
  schedule {
    recurrence_type = "WEEKLY"
    recurrence {
      day_of_week = "THURSDAY"
      start_time = "19:21"
      duration_minutes = 60
    }
    start = "2020-10-20 15:38"
    end = "2020-10-25 15:38"
    zone_id = "America/Chicago"
  }
  
}
```

## Argument Reference

* `name` - (Required) The name of the maintenance window, displayed in the UI.
* `description` - (Required) A short description of the maintenance purpose.
* `type` - (Required) The type of the maintenance: planned or unplanned.
* `suppression` - (Required) The type of suppression of alerting and problem detection during the maintenance.
* `scope` - (Optional) The scope of the maintenance window.
* `schedule` - (Required) The schedule of the maintenance window.

## Attribute Reference

* `id` - (Optional) The ID of the maintenance window.

## Nested Blocks

`scope`

### Arguments

* `entities` - (Required) A list of Dynatrace entities (for example, hosts or services) to be included in the scope.
* `matches` - (Required) A list of matching rules for dynamic scope formation.

`matches`

### Arguments

* `type` - (Optional) The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching.
* `mz_id` - (Optional) The ID of a management zone to which the matched entities must belong.
* `tags` - (Required) The tag you want to use for matching.
* `tag_combination` (Optional) The logic that applies when several tags are specified: AND/OR.

`tags`

### Arguments

* `context` - (Required) The origin of the tag, such as AWS or Cloud Foundry.
* `key` - (Required) The key of the tag.
* `value` - (Optional) The value of the tag.

`schedule`

### Arguments

* `recurrence_type` - (Required) The type of the schedule recurrence.
* `recurrence` - (Optional) The recurrence of the maintenance window.
* `start` - (Required) The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
* `end` - (Required) The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
* `zone_id` - (Required) The time zone of the start and end time. Default time zone is UTC.

`recurrence`

### Arguments

* `day_of_week` - (Optional) The day of the week for weekly maintenance.
* `day_of_month` - (Optional) The day of the month for monthly maintenance.
* `start_time` - (Required) The start time of the maintenance window in HH:mm format.
* `duration_minutes` - (Required) The duration of the maintenance window in minutes.

## Import

Dynatrace maintenance windows can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_maintenance_window.linux_updates -4638826838889583423
```

[Maintenance Windows API]: (https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/maintenance-windows-api/)
