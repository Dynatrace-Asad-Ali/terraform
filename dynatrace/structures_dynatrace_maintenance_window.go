package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandMaintenanceWindowScope(scope []interface{}) dynatraceConfigV1.Scope {
	if len(scope) == 0 || scope[0] == nil {
		return dynatraceConfigV1.Scope{}
	}

	m := scope[0].(map[string]interface{})

	mws := dynatraceConfigV1.Scope{}

	if entities, ok := m["entities"]; ok {
		mws.Entities = expandEntities(entities.(*schema.Set).List())
	}

	if matches, ok := m["match"]; ok {
		mws.Matches = expandMatches(matches.([]interface{}))
	}

	return mws

}

func expandEntities(entities []interface{}) []string {
	mwe := make([]string, len(entities))

	for i, v := range entities {
		mwe[i] = v.(string)
	}

	return mwe

}

func expandMatches(matches []interface{}) []dynatraceConfigV1.MonitoredEntityFilter {
	if len(matches) < 1 {
		return []dynatraceConfigV1.MonitoredEntityFilter{}
	}

	mwm := make([]dynatraceConfigV1.MonitoredEntityFilter, len(matches))

	for i, match := range matches {

		m := match.(map[string]interface{})

		mwm[i] = dynatraceConfigV1.MonitoredEntityFilter{
			Type:           m["type"].(string),
			MzId:           m["mz_id"].(string),
			Tags:           expandTags(m["tags"].([]interface{})),
			TagCombination: m["tag_combination"].(string),
		}
	}

	return mwm
}

func expandTags(tags []interface{}) []dynatraceConfigV1.TagInfo {
	if len(tags) < 1 {
		return []dynatraceConfigV1.TagInfo{}
	}

	mwt := make([]dynatraceConfigV1.TagInfo, len(tags))

	for i, tag := range tags {

		m := tag.(map[string]interface{})

		mwt[i] = dynatraceConfigV1.TagInfo{
			Context: m["context"].(string),
			Key:     m["key"].(string),
			Value:   m["value"].(string),
		}
	}

	return mwt
}

func expandMaintenanceWindowSchedule(schedule []interface{}) dynatraceConfigV1.Schedule {
	if len(schedule) == 0 || schedule[0] == nil {
		return dynatraceConfigV1.Schedule{}
	}

	m := schedule[0].(map[string]interface{})

	mws := dynatraceConfigV1.Schedule{}

	if recurrenceType, ok := m["recurrence_type"]; ok {
		mws.RecurrenceType = recurrenceType.(string)
	}

	if recurrence, ok := m["recurrence"]; ok && len(m["recurrence"].([]interface{})) != 0 {
		mws.Recurrence = expandRecurrence(recurrence.([]interface{}))
	}

	if start, ok := m["start"]; ok {
		mws.Start = start.(string)
	}

	if end, ok := m["end"]; ok {
		mws.End = end.(string)
	}

	if zoneID, ok := m["zone_id"]; ok {
		mws.ZoneId = zoneID.(string)
	}

	return mws

}

func expandRecurrence(recurrence []interface{}) *dynatraceConfigV1.Recurrence {
	if len(recurrence) == 0 || recurrence[0] == nil {
		return &dynatraceConfigV1.Recurrence{}
	}

	m := recurrence[0].(map[string]interface{})

	mwr := dynatraceConfigV1.Recurrence{}

	if dayOfWeek, ok := m["day_of_week"]; ok {
		mwr.DayOfWeek = dayOfWeek.(string)
	}

	if dayOfMonth, ok := m["day_of_month"]; ok {
		mwr.DayOfMonth = dayOfMonth.(int)
	}

	if startTime, ok := m["start_time"]; ok {
		mwr.StartTime = startTime.(string)
	}

	if durationMinutes, ok := m["duration_minutes"]; ok {
		mwr.DurationMinutes = durationMinutes.(int)
	}

	return &mwr

}

func flattenMaintenanceWindowScopeData(maintenanceWindowScope *dynatraceConfigV1.Scope) []interface{} {
	if maintenanceWindowScope == nil {
		return []interface{}{maintenanceWindowScope}
	}

	m := make(map[string]interface{})

	m["entities"] = maintenanceWindowScope.Entities
	m["match"] = flattenMaintenanceWindowMatches(&maintenanceWindowScope.Matches)

	return []interface{}{m}
}

func flattenMaintenanceWindowMatches(maintenanceWindowMatches *[]dynatraceConfigV1.MonitoredEntityFilter) []interface{} {
	if maintenanceWindowMatches != nil {
		mwm := make([]interface{}, len(*maintenanceWindowMatches), len(*maintenanceWindowMatches))

		for i, maintenanceWindowMatches := range *maintenanceWindowMatches {
			mm := make(map[string]interface{})

			mm["type"] = maintenanceWindowMatches.Type
			mm["mz_id"] = maintenanceWindowMatches.MzId
			mm["tags"] = flattenMaintenanceWindowTags(&maintenanceWindowMatches.Tags)
			mm["tag_combination"] = maintenanceWindowMatches.TagCombination
			mwm[i] = mm
		}

		return mwm
	}

	return make([]interface{}, 0)

}

func flattenMaintenanceWindowTags(maintenanceWindowTags *[]dynatraceConfigV1.TagInfo) []interface{} {
	if maintenanceWindowTags != nil {
		mwt := make([]interface{}, len(*maintenanceWindowTags), len(*maintenanceWindowTags))

		for i, maintenanceWindowTags := range *maintenanceWindowTags {
			mt := make(map[string]interface{})

			mt["context"] = maintenanceWindowTags.Context
			mt["key"] = maintenanceWindowTags.Key
			mt["value"] = maintenanceWindowTags.Value
			mwt[i] = mt
		}

		return mwt
	}

	return make([]interface{}, 0)

}

func flattenMaintenanceWindowScheduleData(maintenanceWindowSchedule *dynatraceConfigV1.Schedule) []interface{} {
	if maintenanceWindowSchedule == nil {
		return []interface{}{maintenanceWindowSchedule}
	}

	m := make(map[string]interface{})

	m["recurrence_type"] = maintenanceWindowSchedule.RecurrenceType
	m["recurrence"] = flattenMaintenanceWindowRecurrence(maintenanceWindowSchedule.Recurrence)
	m["start"] = maintenanceWindowSchedule.Start
	m["end"] = maintenanceWindowSchedule.End
	m["zone_id"] = maintenanceWindowSchedule.ZoneId

	return []interface{}{m}
}

func flattenMaintenanceWindowRecurrence(maintenanceWindowRecurrence *dynatraceConfigV1.Recurrence) []interface{} {
	if maintenanceWindowRecurrence == nil {
		return nil
	}

	m := make(map[string]interface{})

	m["day_of_week"] = maintenanceWindowRecurrence.DayOfWeek
	m["day_of_month"] = maintenanceWindowRecurrence.DayOfMonth
	m["start_time"] = maintenanceWindowRecurrence.StartTime
	m["duration_minutes"] = maintenanceWindowRecurrence.DurationMinutes

	return []interface{}{m}
}
