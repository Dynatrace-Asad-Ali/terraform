# Resource: dynatrace_dashboard

Provides a dynatrace dashboard resource. It allows to create, update, delete dashboards in a dynatrace environment. [Dashboards API]

## Example Usage

```hcl
resource "dynatrace_dashboard" "test" {
  dashboard_metadata {
    dashboard_filter {
      timeframe = "l_7_DAYS"
    }
    name = "My Dashboard"
    preset = false
    shared = true
    tags = ["tag1", "tag2"]
    sharing_details {
      link_shared = true
      published = false
    }
  }

  tile {
    chart_visible = false
    configured = true
    custom_name = ""
    exclude_maintenance_windows = false
    limit = 0
    name = "Infrastructure"
    tile_type = "HEADER"
    tile_filter {}
    bounds {
      top = 0
      left = 0
      width = 304
      height = 38
    }
  }

  tile {
    name = "Host CPU Load"
    tile_filter {}
    tile_type = "CUSTOM_CHARTING"
    limit = 0
    bounds {
      top = 38
      left = 0
      width = 456
      height = 152
    }
    chart_visible = false
    configured = true
    exclude_maintenance_windows = false
    filter_config {
      chart_config {
        legend_shown = true
        series {
          aggregation = "AVG"
          aggregation_rate = "TOTAL"
          entity_type = "HOST"
          metric = "builtin:host.cpu.load"
          percentile = 0
          sort_ascending = false
          sort_column = true
          type = "LINE"
          dimensions {
            entity_dimension = true
            id = "0"
            name = "dt.entity.host"
          }
        }
        type = "TIMESERIES"
      }
      custom_name = "CPU"
      default_name = "Custom Chart"
      type = "MIXED"
      filters_per_entity_type = jsonencode(
        {
          HOST = {
            AUTO_TAGS = [
              "keptn_stage:production"
              ]
          }
        }
      )
    }
  }

  tile {
    name = "hosts"
    tile_filter {}
    chart_visible = false
    configured = true
    limit = 0
    tile_type = "HOSTS"
    bounds {
      top = 190
      left = 608
      width = 304
      height = 304
    }
    exclude_maintenance_windows = false
    filter_config {
      chart_config {
        legend_shown = true
        type = "TIMESERIES"
      }
      custom_name = "Hosts"
      default_name = "Hosts"
      filters_per_entity_type = "{}"
      type = "HOST"
    }
  }

  tile {
    name = "Network Status"
    chart_visible = false
    configured = true
    exclude_maintenance_windows = false
    limit = 0
    tile_filter {}
    tile_type = "NETWORK_MEDIUM"
    bounds {
      top = 38
      left = 456
      width = 456
      height = 152
    }
  }

  tile {
    chart_visible = false
    configured = true
    exclude_maintenance_windows = false
    limit = 0
    name = "Service health"
    tile_type = "SERVICES"
    bounds {
      top = 190
      left = 0
      width = 304
      height = 304
    }

    tile_filter {
      management_zone {
        id   = "-8012204502378741258"
        name = "Frontend Services"
      }
    }
  }

  tile {
    chart_visible = false
    configured = true 
    exclude_maintenance_windows = false
    limit = 0
    name = "Top web applications"
    tile_type = "APPLICATIONS_MOST_ACTIVE"
    bounds {
      top = 190
      left = 304
      width = 304
      height = 304
      }
    tile_filter {}
  }
}
```

## Argument Reference

* `dashboard_metadata` - (Required) Parameters of a dashboard.
* `tile` - (Required) The list of tiles on the dashboard.

## Attribute Reference

* `id` - (Optional) The ID of the dynatrace dashboard.

## Nested Blocks

`dashboard_metadata`

### Arguments

* `name` - (Required) The name of the dashboard.
* `shared` - (Optional) The dashboard is shared (true) or private (false).
* `owner` - (Optional) The owner of the dashboard.
* `tags` - (Optional) A set of tags assigned to the dashboard.
* `preset`- (optional) The dashboard is a preset (true)
* `valid_filter_keys`- (optional) A set of all possible global dashboard filters that can be applied to dashboard.
* `sharing_details`- (optional) Sharing configuration of a dashboard.
* `dashboard_filter` (optional) Filters, applied to a dashboard.

`sharing_details`

### Arguments

* `link_shared` - (optional) The default timeframe of the dashboard.
* `published` - (optional) If true, the dashboard is published to anyone on this environment.

`dashboard_filter`

* `timeframe` - (optional) If true, the dashboard is shared via link and authenticated users with the link can view.
* `management_zone` - (optional) The short representation of a Dynatrace entity.

`management_zone`

### Arguments

* `id` - (optional) The ID of the Dynatrace entity.
* `name` - (optional) The name of the Dynatrace entity.

`tile`

### Arguments

* `name` - (Required) The name of the tile.
* `tile_type` - (Required) Defines the actual set of fields depending on the value.
* `configured` - (Optional) The tile is configured and ready to use (true) or just placed on the dashboard (false).
* `bounds` - (Required) The position and size of a tile.
* `tileFilter` - (Optional) A filter applied to a tile. It overrides dashboard's filter.
* `assigned_entites` - (Optional) The list of Dynatrace entities, assigned to the tile.
* `metric` - (Optional) The tile is visible and ready to use (true) or just placed on the dashboard (false).
* `filter_config` - (Optional) Configuration of the custom filter of a tile.
* `chart_visible` - (Optional) The tile is visible and ready to use (true) or just placed on the dashboard (false).
* `markdown` - (Optional) The markdown-formatted content of the tile.
* `exclude_maintenance_windows` - (Optional) Include (`false') or exclude (`true`) maintenance windows from availability calculations.
* `custom_name` - (Optional) The markdown-formatted content of the tile.
* `query` - (Optional) A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile.
* `type` - (Optional) The visualization of the tile.
* `timeframe_shift` - (Optional) The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift.
* `visualization_config` - (Optional) A filter applied to a tile, it overrides the dashboard's filter.
* `limit` - (Optional) The limit of the results, if not set will use the default value of the system.

`bounds`

### Arguments

* `top` - (Optional) The vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.
* `left` - (Optional) The horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels.
* `width` - (Optional) The width of the tile, in pixels.
* `height` - (Optional) The height of the tile, in pixels.

`filter_config`

### Arguments

* `type` - (Required) The type of the filter.
* `custom_name` - (Required) The name of the tile, set by user.
* `default_name` - (Required) The default name of the tile.
* `filters_per_entity_type` - (Optional) A list of filters, applied to specific entity types containing plain JSON-encoded data.
* `chart_config` - (Optional) Configuration of a custom chart.

`chart_config`

### Arguments

* `legend_shown"` - (Optional) Defines if a legend should be shown.
* `type` - (Required) The type of the chart.
* `result_metadata` - (Optional) Additional information about charted metric.
* `axis_limits` - (Optional) The optional custom y-axis limits containing plain JSON-encoded data.
* `left_axis_custom_unit` - (Optional) The custom unit for the left Y-axis.
* `right_axis_custom_unit` - (Optional) The custom unit for the right Y-axis.
* `series` - (Optional) A list of charted metrics.

`result_metadata`

### Arguments

* `last_modified` - (Optional) The timestamp of the last metadata modification, in UTC milliseconds.
* `custom_color` - (Optional) The color of the metric in the chart, hex format.

`series`

### Arguments

* `metric` - (Required) The name of the charted metric.
* `aggregation` - (Required) The charted aggregation of the metric.
* `percentile` - (Optional) The charted percentile.
* `type` - (Required) The visualization of the timeseries chart.
* `entity_type` - (Required) The type of the Dynatrace entity that delivered the charted metric.
* `sort_ascending` - (Optional) Sort ascending (true) or descending (false).
* `sort_column` - (Optional) Sort column (true) or (false).
* `aggregation_rate` - (Optional) The aggregation rate.
* `dimensions` - (Optional) Configuration of the charted metric splitting.

`dimensions`

### Arguments

* `id` - (Required) The ID of the dimension by which the metric is split.
* `name` - (Optional) The name of the dimension by which the metric is split.
* `values` - (Optional) The splitting value.
* `entity_dimension` - (Optional) The name of the entity dimension by which the metric is split.


`visualization_config`

### Arguments

* `has_axis_bucketing` - (Optional) The axis bucketing when enabled groups similar series in the same virtual axis.



## Import

Dynatrace dashboards can be imported using their ID, e.g.

```hcl
$ terraform import dynatrace_dashboard.my_dashboard 4aa520d-c9e1-4a5a-9a5a-3fd27f3a79e8
```

[Dashboards API]: https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/dashboards-api/
