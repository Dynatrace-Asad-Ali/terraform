package dynatrace

import (
	"reflect"
	"testing"

	dynatraceConfigV1 "github.com/dynatrace-ace/dynatrace-go-api-client/api/v1/config/dynatrace"
)

func TestFlattenConditionKey(t *testing.T) {
	cases := []struct {
		Input          *dynatraceConfigV1.ConditionKey
		ExpectedOutput []interface{}
	}{
		{
			&dynatraceConfigV1.ConditionKey{
				Attribute:  "SERVICE_TAGS",
				DynamicKey: "ENVIRONMENT",
				Type:       "PROCESS_CUSTOM_METADATA_KEY",
			},
			[]interface{}{
				map[string]interface{}{
					"attribute":   "SERVICE_TAGS",
					"dynamic_key": "ENVIRONMENT",
					"type":        "PROCESS_CUSTOM_METADATA_KEY",
				},
			},
		},
	}
	for _, tc := range cases {
		output := flattenManagementZoneKey(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandConditionKey(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput dynatraceConfigV1.ConditionKey
	}{
		{
			[]interface{}{
				map[string]interface{}{
					"attribute":   "SERVICE_TAGS",
					"dynamic_key": "ENVIRONMENT",
					"type":        "PROCESS_CUSTOM_METADATA_KEY",
				},
			},
			dynatraceConfigV1.ConditionKey{
				Attribute:  "SERVICE_TAGS",
				DynamicKey: "ENVIRONMENT",
				Type:       "PROCESS_CUSTOM_METADATA_KEY",
			},
		},
	}
	for _, tc := range cases {
		output := expandConditionKey(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
