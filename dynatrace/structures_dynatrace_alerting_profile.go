package dynatrace

import (
	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

func expandAlertingProfileRules(rules []interface{}) []dynatraceConfigV1.AlertingProfileSeverityRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.AlertingProfileSeverityRule{}
	}

	ars := []dynatraceConfigV1.AlertingProfileSeverityRule{}

	for _, rule := range rules {
		i := rule.(map[string]interface{})
		tf := i["tag_filters"].([]interface{})[0]
		tagFilter := tf.(map[string]interface{})

		ar := dynatraceConfigV1.AlertingProfileSeverityRule{
			SeverityLevel: i["severity_level"].(string),
			TagFilter: dynatraceConfigV1.AlertingProfileTagFilter{
				IncludeMode: tagFilter["include_mode"].(string),
				TagFilters:  expandAlertingProfileTagFilters(tagFilter["tag_filter"].([]interface{})),
			},
			DelayInMinutes: i["delay_in_minutes"].(int),
		}
		ars = append(ars, ar)
	}

	return ars

}

func expandAlertingProfileTagFilters(tagFilters []interface{}) []dynatraceConfigV1.TagFilter {
	if len(tagFilters) == 0 || tagFilters[0] == nil {
		return []dynatraceConfigV1.TagFilter{}
	}

	tfs := make([]dynatraceConfigV1.TagFilter, len(tagFilters))

	for i, filter := range tagFilters {

		m := filter.(map[string]interface{})

		tfs[i] = dynatraceConfigV1.TagFilter{
			Context: m["context"].(string),
			Key:     m["key"].(string),
			Value:   m["value"].(string),
		}

	}

	return tfs

}

func expandEventTypeFilters(typeFilters []interface{}) []dynatraceConfigV1.AlertingEventTypeFilter {
	if len(typeFilters) == 0 || typeFilters[0] == nil {
		return []dynatraceConfigV1.AlertingEventTypeFilter{}
	}

	etfs := make([]dynatraceConfigV1.AlertingEventTypeFilter, len(typeFilters))

	for i, typeFilter := range typeFilters {

		m := typeFilter.(map[string]interface{})

		if v, ok := m["predefined_event_filter"]; ok && len(m["predefined_event_filter"].([]interface{})) != 0 {
			etfs[i] = dynatraceConfigV1.AlertingEventTypeFilter{
				PredefinedEventFilter: expandPredefinedEventFilter(v.([]interface{})),
			}
		}

		if v, ok := m["custom_event_filter"]; ok && len(m["custom_event_filter"].([]interface{})) != 0 {
			cef := v.([]interface{})[0]
			customEventFilter := cef.(map[string]interface{})

			etfs[i] = dynatraceConfigV1.AlertingEventTypeFilter{
				CustomEventFilter: &dynatraceConfigV1.AlertingCustomEventFilter{
					CustomTitleFilter:       expandCustomTextFilter(customEventFilter["custom_title_filter"].([]interface{})),
					CustomDescriptionFilter: expandCustomTextFilter(customEventFilter["custom_description_filter"].([]interface{})),
				},
			}
		}
	}

	return etfs
}

func expandPredefinedEventFilter(predefinedEventFilter []interface{}) *dynatraceConfigV1.AlertingPredefinedEventFilter {
	if len(predefinedEventFilter) == 0 || predefinedEventFilter[0] == nil {
		return nil
	}

	m := predefinedEventFilter[0].(map[string]interface{})

	pef := &dynatraceConfigV1.AlertingPredefinedEventFilter{}

	if negate, ok := m["negate"]; ok {
		pef.Negate = negate.(bool)
	}

	if eventType, ok := m["event_type"]; ok {
		pef.EventType = eventType.(string)
	}

	return pef

}

func expandCustomTextFilter(customTextFilter []interface{}) *dynatraceConfigV1.AlertingCustomTextFilter {
	if len(customTextFilter) == 0 || customTextFilter[0] == nil {
		return nil
	}

	m := customTextFilter[0].(map[string]interface{})

	ctf := &dynatraceConfigV1.AlertingCustomTextFilter{}

	if enabled, ok := m["enabled"]; ok {
		ctf.Enabled = enabled.(bool)
	}

	if value, ok := m["value"]; ok {
		ctf.Value = value.(string)
	}

	if operator, ok := m["operator"]; ok {
		ctf.Operator = operator.(string)
	}

	if negate, ok := m["negate"]; ok {
		ctf.Negate = negate.(bool)
	}

	if caseInsensitive, ok := m["case_insensitive"]; ok {
		ctf.CaseInsensitive = caseInsensitive.(bool)
	}

	return ctf

}

func flattenAlertingProfileRulesData(alertingProfileRules *[]dynatraceConfigV1.AlertingProfileSeverityRule) []interface{} {
	if alertingProfileRules != nil {
		ars := make([]interface{}, len(*alertingProfileRules), len(*alertingProfileRules))

		for i, alertingProfileRules := range *alertingProfileRules {
			ar := make(map[string]interface{})

			ar["severity_level"] = alertingProfileRules.SeverityLevel
			ar["delay_in_minutes"] = alertingProfileRules.DelayInMinutes
			ar["tag_filters"] = flattenAlertingProfileTagFilter(&alertingProfileRules.TagFilter)
			ars[i] = ar
		}

		return ars
	}

	return make([]interface{}, 0)
}

func flattenAlertingProfileTagFilter(alertingProfileTagFilter *dynatraceConfigV1.AlertingProfileTagFilter) []interface{} {
	if alertingProfileTagFilter == nil {
		return []interface{}{alertingProfileTagFilter}
	}
	t := make(map[string]interface{})

	t["include_mode"] = alertingProfileTagFilter.IncludeMode
	t["tag_filter"] = flattenAlertingProfileTagFilters(&alertingProfileTagFilter.TagFilters)

	return []interface{}{t}

}

func flattenAlertingProfileTagFilters(alertingProfileTagFilters *[]dynatraceConfigV1.TagFilter) []interface{} {
	if alertingProfileTagFilters != nil {
		tfs := make([]interface{}, len(*alertingProfileTagFilters), len(*alertingProfileTagFilters))

		for i, alertingProfileTagFilters := range *alertingProfileTagFilters {
			tf := make(map[string]interface{})

			tf["context"] = alertingProfileTagFilters.Context
			tf["key"] = alertingProfileTagFilters.Key
			tf["value"] = alertingProfileTagFilters.Value
			tfs[i] = tf
		}

		return tfs
	}

	return make([]interface{}, 0)
}

func flattenAlertingProfileEventTypeFiltersData(alertingProfileEventTypeFilters *[]dynatraceConfigV1.AlertingEventTypeFilter) []interface{} {
	if alertingProfileEventTypeFilters != nil {

		efs := make([]interface{}, len(*alertingProfileEventTypeFilters), len(*alertingProfileEventTypeFilters))

		for i, alertingProfileEventTypeFilters := range *alertingProfileEventTypeFilters {
			ef := make(map[string]interface{})

			ef["predefined_event_filter"] = flattenPredefinedEventFilter(alertingProfileEventTypeFilters.PredefinedEventFilter)
			ef["custom_event_filter"] = flattenCustomEventFilter(alertingProfileEventTypeFilters.CustomEventFilter)
			efs[i] = ef

		}
		return efs
	}

	return make([]interface{}, 0)
}

func flattenPredefinedEventFilter(alertingProfilePredefinedEventFilters *dynatraceConfigV1.AlertingPredefinedEventFilter) []interface{} {
	if alertingProfilePredefinedEventFilters == nil {
		return nil
	}

	pef := make(map[string]interface{})

	pef["event_type"] = alertingProfilePredefinedEventFilters.EventType
	pef["negate"] = alertingProfilePredefinedEventFilters.Negate

	return []interface{}{pef}

}

func flattenCustomEventFilter(alertingProfileCustomEventFilters *dynatraceConfigV1.AlertingCustomEventFilter) []interface{} {
	if alertingProfileCustomEventFilters == nil {
		return nil
	}

	cef := make(map[string]interface{})

	cef["custom_title_filter"] = flattenCustomTextFilter(alertingProfileCustomEventFilters.CustomTitleFilter)
	cef["custom_description_filter"] = flattenCustomTextFilter(alertingProfileCustomEventFilters.CustomDescriptionFilter)

	return []interface{}{cef}

}

func flattenCustomTextFilter(alertingProfileCustomTextFilters *dynatraceConfigV1.AlertingCustomTextFilter) []interface{} {
	if alertingProfileCustomTextFilters == nil {
		return nil
	}

	ctf := make(map[string]interface{})

	ctf["enabled"] = alertingProfileCustomTextFilters.Enabled
	ctf["value"] = alertingProfileCustomTextFilters.Value
	ctf["operator"] = alertingProfileCustomTextFilters.Operator
	ctf["negate"] = alertingProfileCustomTextFilters.Negate
	ctf["case_insensitive"] = alertingProfileCustomTextFilters.CaseInsensitive

	return []interface{}{ctf}
}
