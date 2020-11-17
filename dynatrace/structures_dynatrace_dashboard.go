package dynatrace

import (
	"encoding/json"
	"log"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandDashboardMetadata(dashboardMetadata []interface{}) dynatraceConfigV1.DashboardMetadata {
	if len(dashboardMetadata) == 0 || dashboardMetadata[0] == nil {
		return dynatraceConfigV1.DashboardMetadata{}
	}

	m := dashboardMetadata[0].(map[string]interface{})

	ddm := dynatraceConfigV1.DashboardMetadata{}

	if name, ok := m["name"]; ok {
		ddm.Name = name.(string)
	}

	if shared, ok := m["shared"]; ok {
		ddm.Shared = shared.(bool)
	}

	if owner, ok := m["owner"]; ok {
		ddm.Owner = owner.(string)
	}

	if sharingDetails, ok := m["sharing_details"]; ok {
		ddm.SharingDetails = expandDashboardSharingDetails(sharingDetails.([]interface{}))
	}

	if dashboardFilter, ok := m["dashboard_filter"]; ok {
		ddm.DashboardFilter = expandDashboardFilter(dashboardFilter.([]interface{}))
	}

	if tags, ok := m["tags"]; ok {
		ddm.Tags = expandDashboardTags(tags.(*schema.Set).List())
	}

	if validFilterKeys, ok := m["valid_filter_keys"]; ok {
		ddm.ValidFilterKeys = expandDashboardFilterKeys(validFilterKeys.(*schema.Set).List())
	}

	return ddm

}

func expandDashboardSharingDetails(sharingDetails []interface{}) dynatraceConfigV1.SharingInfo {
	if len(sharingDetails) == 0 || sharingDetails[0] == nil {
		return dynatraceConfigV1.SharingInfo{}
	}

	m := sharingDetails[0].(map[string]interface{})

	dsd := dynatraceConfigV1.SharingInfo{}

	if linkShared, ok := m["link_shared"]; ok {
		dsd.LinkShared = linkShared.(bool)
	}

	if published, ok := m["published"]; ok {
		dsd.Published = published.(bool)
	}

	return dsd
}

func expandDashboardFilter(filters []interface{}) dynatraceConfigV1.DashboardFilter {
	if len(filters) == 0 || filters[0] == nil {
		return dynatraceConfigV1.DashboardFilter{}
	}

	m := filters[0].(map[string]interface{})

	ddf := dynatraceConfigV1.DashboardFilter{}

	if timeframe, ok := m["timeframe"]; ok {
		ddf.Timeframe = timeframe.(string)
	}

	if managementZone, ok := m["management_zone"]; ok {
		ddf.ManagementZone = expandDashboardManagementZone(managementZone.([]interface{}))
	}

	return ddf

}

func expandDashboardTags(tags []interface{}) []string {
	ddt := make([]string, len(tags))

	for i, v := range tags {
		ddt[i] = v.(string)
	}

	return ddt

}

func expandDashboardFilterKeys(filterKeys []interface{}) []string {
	dfk := make([]string, len(filterKeys))

	for i, v := range filterKeys {
		dfk[i] = v.(string)
	}

	return dfk

}

func expandDashboardManagementZone(managementZone []interface{}) *dynatraceConfigV1.EntityShortRepresentation {
	if len(managementZone) == 0 || managementZone[0] == nil {
		return &dynatraceConfigV1.EntityShortRepresentation{}
	}

	m := managementZone[0].(map[string]interface{})

	dmz := dynatraceConfigV1.EntityShortRepresentation{}

	if id, ok := m["id"]; ok {
		dmz.Id = id.(string)
	}

	if name, ok := m["name"]; ok {
		dmz.Name = name.(string)
	}

	return &dmz

}

func expandDashboardTiles(dashboardTiles []interface{}) []dynatraceConfigV1.Tile {
	if len(dashboardTiles) < 1 {
		return []dynatraceConfigV1.Tile{}
	}

	dts := make([]dynatraceConfigV1.Tile, len(dashboardTiles))

	for i, tile := range dashboardTiles {

		m := tile.(map[string]interface{})

		dts[i] = dynatraceConfigV1.Tile{
			Name:             m["name"].(string),
			TileType:         m["tile_type"].(string),
			Configured:       m["configured"].(bool),
			Bounds:           expandTileBounds(m["bounds"].([]interface{})),
			TileFilter:       expandTileFilter(m["tile_filter"].([]interface{})),
			AssignedEntities: expandAssignedEntities(m["assigned_entities"].(*schema.Set).List()),
			Metric:           m["metric"].(string),
			FilterConfig:     expandFilterConfig(m["filter_config"].([]interface{})),
		}
	}

	return dts
}

func expandTileBounds(bounds []interface{}) dynatraceConfigV1.TileBounds {
	if len(bounds) == 0 || bounds[0] == nil {
		return dynatraceConfigV1.TileBounds{}
	}

	m := bounds[0].(map[string]interface{})

	dtb := dynatraceConfigV1.TileBounds{}

	if top, ok := m["top"]; ok {
		dtb.Top = top.(int)
	}

	if left, ok := m["left"]; ok {
		dtb.Left = left.(int)
	}

	if width, ok := m["width"]; ok {
		dtb.Width = width.(int)
	}

	if height, ok := m["height"]; ok {
		dtb.Height = height.(int)
	}

	return dtb

}

func expandTileFilter(filter []interface{}) dynatraceConfigV1.TileFilter {
	if len(filter) == 0 || filter[0] == nil {
		return dynatraceConfigV1.TileFilter{}
	}

	m := filter[0].(map[string]interface{})

	dtf := dynatraceConfigV1.TileFilter{}

	if timeframe, ok := m["timeframe"]; ok {
		dtf.Timeframe = timeframe.(string)
	}

	if managementZone, ok := m["management_zone"]; ok {
		dtf.ManagementZone = expandDashboardManagementZone(managementZone.([]interface{}))
	}

	return dtf

}

func expandAssignedEntities(assignedEntities []interface{}) []string {
	dae := make([]string, len(assignedEntities))

	for i, v := range assignedEntities {
		dae[i] = v.(string)
	}

	return dae

}

func expandFilterConfig(filterConfig []interface{}) *dynatraceConfigV1.CustomFilterConfig {
	if len(filterConfig) == 0 || filterConfig[0] == nil {
		return nil
	}

	m := filterConfig[0].(map[string]interface{})

	dfc := dynatraceConfigV1.CustomFilterConfig{}

	if dfType, ok := m["type"]; ok {
		dfc.Type = dfType.(string)
	}

	if customName, ok := m["custom_name"]; ok {
		dfc.CustomName = customName.(string)
	}

	if defaultName, ok := m["default_name"]; ok {
		dfc.DefaultName = defaultName.(string)
	}

	if chartConfig, ok := m["chart_config"]; ok {
		dfc.ChartConfig = expandDashboardChartConfig(chartConfig.([]interface{}))
	}

	if filtersPerEntityType, ok := m["filters_per_entity_type"]; ok {
		dfc.FiltersPerEntityType = expandFiltersPerEntityType(filtersPerEntityType.(string))
	}

	return &dfc

}

func expandDashboardChartConfig(chartConfig []interface{}) dynatraceConfigV1.CustomFilterChartConfig {
	if len(chartConfig) == 0 || chartConfig[0] == nil {
		return dynatraceConfigV1.CustomFilterChartConfig{}
	}

	m := chartConfig[0].(map[string]interface{})

	dcc := dynatraceConfigV1.CustomFilterChartConfig{}

	if legendShown, ok := m["legend_shown"]; ok {
		dcc.LegendShown = legendShown.(bool)
	}

	if dcType, ok := m["type"]; ok {
		dcc.Type = dcType.(string)
	}

	if series, ok := m["series"]; ok {
		dcc.Series = expandCustomChartSeries(series.([]interface{}))
	}

	if resultMetadata, ok := m["result_metadata"]; ok {
		dcc.ResultMetadata = expandResultMetadata(resultMetadata.((string)))
	}

	if axisLimits, ok := m["axis_limits"]; ok {
		dcc.AxisLimits = expandAxisLimits(axisLimits.((string)))
	}

	if leftAxisCustomLimit, ok := m["left_axis_custom_limit"]; ok {
		dcc.LeftAxisCustomUnit = leftAxisCustomLimit.(string)
	}

	if rightAxisCustomLimit, ok := m["left_axis_custom_limit"]; ok {
		dcc.RightAxisCustomUnit = rightAxisCustomLimit.(string)
	}

	return dcc

}

func expandCustomChartSeries(chartSeries []interface{}) []dynatraceConfigV1.CustomFilterChartSeriesConfig {
	if len(chartSeries) < 1 {
		return []dynatraceConfigV1.CustomFilterChartSeriesConfig{}
	}

	ccs := make([]dynatraceConfigV1.CustomFilterChartSeriesConfig, len(chartSeries))

	for i, series := range chartSeries {

		m := series.(map[string]interface{})

		ccs[i] = dynatraceConfigV1.CustomFilterChartSeriesConfig{
			Metric:          m["metric"].(string),
			Aggregation:     m["aggregation"].(string),
			Percentile:      m["percentile"].(int),
			Type:            m["type"].(string),
			EntityType:      m["entity_type"].(string),
			Dimensions:      expandSeriesDimensions(m["dimensions"].([]interface{})),
			SortAscending:   m["sort_ascending"].(bool),
			SortColumn:      m["sort_column"].(bool),
			AggregationRate: m["aggregation_rate"].(string),
		}
	}

	return ccs
}

func expandSeriesDimensions(dimensions []interface{}) []dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig {
	if len(dimensions) < 1 {
		return []dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig{}
	}

	csd := make([]dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig, len(dimensions))

	for i, dimension := range dimensions {

		m := dimension.(map[string]interface{})

		csd[i] = dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig{
			Id:              m["id"].(string),
			Name:            m["name"].(string),
			Values:          expandDimensionsValues(m["values"].(*schema.Set).List()),
			EntityDimension: m["entity_dimension"].(bool),
		}

	}

	return csd

}

func expandDimensionsValues(values []interface{}) []string {
	ddv := make([]string, len(values))

	for i, v := range values {
		ddv[i] = v.(string)
	}

	return ddv

}

func expandResultMetadata(metadata interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(metadata.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal result metadata values %s: %v", metadata.(string), err)
		return nil
	}

	return val
}

func expandAxisLimits(limits interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(limits.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal axis limits values %s: %v", limits.(string), err)
		return nil
	}

	return val
}

func expandFiltersPerEntityType(filters interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(filters.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal filters per entity type values %s: %v", filters.(string), err)
		return nil
	}

	return val
}

func flattenDashboardMetadata(dashboardMetadata *dynatraceConfigV1.DashboardMetadata) []interface{} {
	if dashboardMetadata == nil {
		return []interface{}{dashboardMetadata}
	}

	m := make(map[string]interface{})

	m["name"] = dashboardMetadata.Name
	m["shared"] = dashboardMetadata.Shared
	m["owner"] = dashboardMetadata.Owner
	m["sharing_details"] = flattenSharingDetails(&dashboardMetadata.SharingDetails)
	m["dashboard_filter"] = flattenDashboardFilter(&dashboardMetadata.DashboardFilter)
	m["tags"] = dashboardMetadata.Tags

	return []interface{}{m}

}

func flattenSharingDetails(sharingDetails *dynatraceConfigV1.SharingInfo) []interface{} {
	if sharingDetails == nil {
		return []interface{}{sharingDetails}
	}

	s := make(map[string]interface{})

	s["link_shared"] = sharingDetails.LinkShared
	s["published"] = sharingDetails.Published

	return []interface{}{s}
}

func flattenDashboardFilter(dashboardFilter *dynatraceConfigV1.DashboardFilter) []interface{} {
	if dashboardFilter == nil {
		return []interface{}{dashboardFilter}
	}

	f := make(map[string]interface{})

	f["timeframe"] = dashboardFilter.Timeframe
	f["management_zone"] = flattenManagementZone(dashboardFilter.ManagementZone)

	return []interface{}{f}

}

func flattenManagementZone(dashboardManagementZone *dynatraceConfigV1.EntityShortRepresentation) []interface{} {
	if dashboardManagementZone == nil {
		return nil
	}

	m := make(map[string]interface{})

	m["id"] = dashboardManagementZone.Id
	m["name"] = dashboardManagementZone.Name

	return []interface{}{m}

}

func flattenDashboardTilesData(dashboardTiles []dynatraceConfigV1.Tile) []interface{} {
	if dashboardTiles != nil {
		dts := make([]interface{}, len(dashboardTiles), len(dashboardTiles))

		for i, dashboardTile := range dashboardTiles {
			dt := make(map[string]interface{})

			dt["name"] = dashboardTile.Name
			dt["tile_type"] = dashboardTile.TileType
			dt["configured"] = dashboardTile.Configured
			dt["bounds"] = flattenTileBounds(&dashboardTile.Bounds)
			dt["tile_filter"] = flattenTileFilter(&dashboardTile.TileFilter)
			dt["filter_config"] = flattenFilterConfig(dashboardTile.FilterConfig)
			dts[i] = dt

		}
		return dts

	}

	return make([]interface{}, 0)
}

func flattenTileBounds(tileBounds *dynatraceConfigV1.TileBounds) []interface{} {
	if tileBounds == nil {
		return []interface{}{tileBounds}
	}

	b := make(map[string]interface{})

	b["top"] = tileBounds.Top
	b["left"] = tileBounds.Left
	b["width"] = tileBounds.Width
	b["height"] = tileBounds.Height

	return []interface{}{b}
}

func flattenTileFilter(tileFilter *dynatraceConfigV1.TileFilter) []interface{} {
	if tileFilter == nil {
		return []interface{}{tileFilter}
	}

	f := make(map[string]interface{})

	f["timeframe"] = tileFilter.Timeframe
	f["management_zone"] = flattenManagementZone(tileFilter.ManagementZone)

	return []interface{}{f}

}

func flattenFilterConfig(filterConfig *dynatraceConfigV1.CustomFilterConfig) []interface{} {
	if filterConfig == nil {
		return nil
	}

	f := make(map[string]interface{})

	f["type"] = filterConfig.Type
	f["custom_name"] = filterConfig.CustomName
	f["default_name"] = filterConfig.DefaultName
	f["chart_config"] = flattenChartConfig(&filterConfig.ChartConfig)
	f["filters_per_entity_type"] = flattenFiltersPerEntityType(&filterConfig.FiltersPerEntityType)

	return []interface{}{f}

}

func flattenChartConfig(chartConfig *dynatraceConfigV1.CustomFilterChartConfig) []interface{} {
	if chartConfig == nil {
		return []interface{}{chartConfig}
	}

	c := make(map[string]interface{})

	c["legend_shown"] = chartConfig.LegendShown
	c["type"] = chartConfig.Type
	c["series"] = flattenFilterSeries(&chartConfig.Series)
	c["axis_limits"] = flattenAxisLimits(&chartConfig.AxisLimits)
	c["result_metadata"] = flattenResultMetadata(&chartConfig.ResultMetadata)
	c["left_axis_custom_unit"] = chartConfig.LeftAxisCustomUnit
	c["right_axis_custom_unit"] = chartConfig.RightAxisCustomUnit
	c["type"] = chartConfig.Type

	return []interface{}{c}

}

func flattenFilterSeries(filterSeries *[]dynatraceConfigV1.CustomFilterChartSeriesConfig) []interface{} {
	if filterSeries != nil {
		csc := make([]interface{}, len(*filterSeries), len(*filterSeries))

		for i, filterSeries := range *filterSeries {
			cs := make(map[string]interface{})

			cs["metric"] = filterSeries.Metric
			cs["aggregation"] = filterSeries.Aggregation
			cs["percentile"] = filterSeries.Percentile
			cs["type"] = filterSeries.Type
			cs["entity_type"] = filterSeries.EntityType
			cs["dimensions"] = flattenChartDimensions(&filterSeries.Dimensions)
			cs["sort_ascending"] = filterSeries.SortAscending
			cs["sort_column"] = filterSeries.SortColumn
			cs["aggregation_rate"] = filterSeries.AggregationRate
			csc[i] = cs

		}
		return csc

	}

	return make([]interface{}, 0)
}

func flattenChartDimensions(chartDimensions *[]dynatraceConfigV1.CustomFilterChartSeriesDimensionConfig) []interface{} {
	if chartDimensions != nil {
		csd := make([]interface{}, len(*chartDimensions), len(*chartDimensions))

		for i, chartDimension := range *chartDimensions {
			cd := make(map[string]interface{})

			cd["id"] = chartDimension.Id
			cd["name"] = chartDimension.Name
			cd["values"] = chartDimension.Values
			cd["entity_dimension"] = chartDimension.EntityDimension
			csd[i] = cd

		}
		return csd

	}

	return make([]interface{}, 0)
}

func flattenFiltersPerEntityType(filters interface{}) interface{} {
	json, err := json.Marshal(filters)
	if err != nil {
		log.Printf("[ERROR] Could not marshal filters per entity type value %s: %v", filters.(string), err)
		return nil
	}

	return string(json)
}

func flattenAxisLimits(limits interface{}) interface{} {
	if limits != nil {
		json, err := json.Marshal(limits)
		if err != nil {
			log.Printf("[ERROR] Could not marshal axis limits values %s: %v", limits.(string), err)
			return nil
		}
		return string(json)
	}
	return make([]interface{}, 0)
}

func flattenResultMetadata(metadata interface{}) interface{} {
	json, err := json.Marshal(metadata)
	if err != nil {
		log.Printf("[ERROR] Could not marshal result metadata values %s: %v", metadata.(string), err)
		return nil
	}

	return string(json)
}
