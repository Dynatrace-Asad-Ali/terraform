package dynatrace

import (
	"encoding/json"
	"log"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandManagementZoneRules(rules []interface{}) []dynatraceConfigV1.ManagementZoneRule {
	if len(rules) < 1 {
		return []dynatraceConfigV1.ManagementZoneRule{}
	}

	mrs := make([]dynatraceConfigV1.ManagementZoneRule, len(rules))

	for i, rule := range rules {

		m := rule.(map[string]interface{})

		mrs[i] = dynatraceConfigV1.ManagementZoneRule{
			Type:             m["type"].(string),
			Enabled:          m["enabled"].(bool),
			PropagationTypes: expandPropagationTypes(m["propagation_types"].(*schema.Set).List()),
			Conditions:       expandConditions(m["condition"].([]interface{})),
		}
	}

	return mrs
}

func expandPropagationTypes(propagationTypes []interface{}) []string {
	pts := make([]string, len(propagationTypes))

	for i, v := range propagationTypes {
		pts[i] = v.(string)
	}

	return pts

}

func expandConditions(conditions []interface{}) []dynatraceConfigV1.EntityRuleEngineCondition {
	if len(conditions) < 1 {
		return []dynatraceConfigV1.EntityRuleEngineCondition{}
	}

	mcs := make([]dynatraceConfigV1.EntityRuleEngineCondition, len(conditions))

	for i, condition := range conditions {

		m := condition.(map[string]interface{})

		mcs[i] = dynatraceConfigV1.EntityRuleEngineCondition{
			Key:            expandConditionKey(m["key"].([]interface{})),
			ComparisonInfo: expandConditionComparisonInfo(m["comparison_info"].([]interface{})),
		}
	}

	return mcs
}

func expandConditionKey(conditionKey []interface{}) dynatraceConfigV1.ConditionKey {
	if len(conditionKey) == 0 || conditionKey[0] == nil {
		return dynatraceConfigV1.ConditionKey{}
	}

	m := conditionKey[0].(map[string]interface{})

	mck := dynatraceConfigV1.ConditionKey{}

	if attribute, ok := m["attribute"]; ok {
		mck.Attribute = attribute.(string)
	}

	if dynamicKey, ok := m["dynamic_key"]; ok {
		mck.DynamicKey = dynamicKey.(string)
	}

	if ckType, ok := m["type"]; ok {
		mck.Type = ckType.(string)
	}

	return mck

}

func expandConditionComparisonInfo(comparisonInfo []interface{}) dynatraceConfigV1.ComparisonBasic {
	if len(comparisonInfo) == 0 || comparisonInfo[0] == nil {
		return dynatraceConfigV1.ComparisonBasic{}
	}

	m := comparisonInfo[0].(map[string]interface{})

	mci := dynatraceConfigV1.ComparisonBasic{}

	if operator, ok := m["operator"]; ok {
		mci.Operator = operator.(string)
	}

	if value, ok := m["value"]; ok {
		mci.Value = expandComparisonInfoValue(value.(string))
	}

	if negate, ok := m["negate"]; ok {
		mci.Negate = negate.(bool)
	}

	if ciType, ok := m["type"]; ok {
		mci.Type = ciType.(string)
	}

	if caseSensitive, ok := m["case_sensitive"]; ok {
		mci.CaseSensitive = caseSensitive.(bool)
	}

	return mci

}

func expandComparisonInfoValue(value interface{}) interface{} {
	var val interface{}
	if err := json.Unmarshal([]byte(value.(string)), &val); err != nil {
		log.Printf("[ERROR] Could not unmarshal comparison info value %s: %v", value.(string), err)
		return nil
	}

	return val
}

func flattenManagementZoneRulesData(managementZoneRules *[]dynatraceConfigV1.ManagementZoneRule) []interface{} {
	if managementZoneRules != nil {
		mrs := make([]interface{}, len(*managementZoneRules), len(*managementZoneRules))

		for i, managementZoneRules := range *managementZoneRules {
			mr := make(map[string]interface{})

			mr["type"] = managementZoneRules.Type
			mr["enabled"] = managementZoneRules.Enabled
			mr["propagation_types"] = managementZoneRules.PropagationTypes
			mr["condition"] = flattenManagementZoneConditionsData(&managementZoneRules.Conditions)
			mrs[i] = mr

		}
		return mrs
	}

	return make([]interface{}, 0)
}

func flattenManagementZoneConditionsData(managementZoneConditions *[]dynatraceConfigV1.EntityRuleEngineCondition) []interface{} {
	if managementZoneConditions != nil {
		mcs := make([]interface{}, len(*managementZoneConditions), len(*managementZoneConditions))

		for i, managementZoneConditions := range *managementZoneConditions {
			mc := make(map[string]interface{})

			mc["key"] = flattenManagementZoneKey(&managementZoneConditions.Key)
			mc["comparison_info"] = flattenManagementZoneComparisonInfo(&managementZoneConditions.ComparisonInfo)
			mcs[i] = mc
		}

		return mcs
	}

	return make([]interface{}, 0)

}

func flattenManagementZoneKey(managementZoneConditionKey *dynatraceConfigV1.ConditionKey) []interface{} {
	if managementZoneConditionKey == nil {
		return []interface{}{managementZoneConditionKey}
	}

	k := make(map[string]interface{})

	k["attribute"] = managementZoneConditionKey.Attribute
	k["type"] = managementZoneConditionKey.Type
	k["dynamic_key"] = managementZoneConditionKey.DynamicKey

	return []interface{}{k}
}

func flattenManagementZoneComparisonInfo(managementZoneComparisonInfo *dynatraceConfigV1.ComparisonBasic) []interface{} {
	if managementZoneComparisonInfo == nil {
		return []interface{}{managementZoneComparisonInfo}
	}

	c := make(map[string]interface{})

	c["operator"] = managementZoneComparisonInfo.Operator
	c["value"] = flattenComparisonInfoValue(&managementZoneComparisonInfo.Value)
	c["negate"] = managementZoneComparisonInfo.Negate
	c["type"] = managementZoneComparisonInfo.Type

	return []interface{}{c}
}

func flattenComparisonInfoValue(value interface{}) interface{} {
	json, err := json.Marshal(value)
	if err != nil {
		log.Printf("[ERROR] Could not marshal comparison info value %s: %v", value.(string), err)
		return nil
	}
	return string(json)

}
