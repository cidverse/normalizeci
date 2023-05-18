package envstruct

import (
	"testing"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

func TestStructToEnvMap(t *testing.T) {
	myStruct := v1.MergeRequest{
		Id:               "123",
		SourceBranchName: "feat/new-feature",
		SourceHash:       "abc",
		TargetBranchName: "main",
		TargetHash:       "def",
	}

	expected := map[string]interface{}{
		"NCI_MERGE_REQUEST_ID":                 "123",
		"NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME": "feat/new-feature",
		"NCI_MERGE_REQUEST_SOURCE_HASH":        "abc",
		"NCI_MERGE_REQUEST_TARGET_BRANCH_NAME": "main",
		"NCI_MERGE_REQUEST_TARGET_HASH":        "def",
	}

	result := StructToEnvMap(myStruct)

	for key, expectedValue := range expected {
		resultValue, exists := result[key]
		if !exists {
			t.Errorf("Key %q is missing in the result map", key)
			continue
		}

		if resultValue != expectedValue {
			t.Errorf("Unexpected value for key %q. Expected: %v, Got: %v", key, expectedValue, resultValue)
		}
	}
}

type withMap struct {
	PipelineId    string            `env:"NCI_PIPELINE_ID"`
	PipelineInput map[string]string `env-prefix:"NCI_INPUT_"`
}

func TestStructToEnvMapWithSubMap(t *testing.T) {
	myStruct := withMap{
		PipelineId: "123",
		PipelineInput: map[string]string{
			"foo": "bar",
			"bar": "foo",
		},
	}

	expected := map[string]interface{}{
		"NCI_PIPELINE_ID": "123",
		"NCI_INPUT_FOO":   "bar",
		"NCI_INPUT_BAR":   "foo",
	}

	result := StructToEnvMap(myStruct)

	for key, expectedValue := range expected {
		resultValue, exists := result[key]
		if !exists {
			t.Errorf("Key %q is missing in the result map", key)
			continue
		}

		if resultValue != expectedValue {
			t.Errorf("Unexpected value for key %q. Expected: %v, Got: %v", key, expectedValue, resultValue)
		}
	}
}

type withSubStruct struct {
	Hello    string `env:"NCI_HELLO"`
	Pipeline withMap
}

func TestStructToEnvMapWithSubStruct(t *testing.T) {
	myStruct := withSubStruct{
		Hello: "World",
		Pipeline: withMap{
			PipelineId: "123",
			PipelineInput: map[string]string{
				"foo": "bar",
				"bar": "foo",
			},
		},
	}

	expected := map[string]interface{}{
		"NCI_HELLO":       "World",
		"NCI_PIPELINE_ID": "123",
		"NCI_INPUT_FOO":   "bar",
		"NCI_INPUT_BAR":   "foo",
	}

	result := StructToEnvMap(myStruct)

	for key, expectedValue := range expected {
		resultValue, exists := result[key]
		if !exists {
			t.Errorf("Key %q is missing in the result map", key)
			continue
		}

		if resultValue != expectedValue {
			t.Errorf("Unexpected value for key %q. Expected: %v, Got: %v", key, expectedValue, resultValue)
		}
	}
}
