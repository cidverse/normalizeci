package envstruct

import (
	"reflect"
	"testing"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

func TestEnvMapToStruct(t *testing.T) {
	envMap := map[string]string{
		"NCI_MERGE_REQUEST_ID":                 "123",
		"NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME": "feat/new-feature",
		"NCI_MERGE_REQUEST_SOURCE_HASH":        "abc",
		"NCI_MERGE_REQUEST_TARGET_BRANCH_NAME": "main",
		"NCI_MERGE_REQUEST_TARGET_HASH":        "def",
	}

	expectedStruct := v1.MergeRequest{
		Id:               "123",
		SourceBranchName: "feat/new-feature",
		SourceHash:       "abc",
		TargetBranchName: "main",
		TargetHash:       "def",
	}

	myStruct := v1.MergeRequest{}
	err := EnvMapToStruct(&myStruct, envMap)
	if err != nil {
		t.Errorf("Error populating struct from env map: %v", err)
	}

	if !reflect.DeepEqual(myStruct, expectedStruct) {
		t.Errorf("Struct mismatch.\nExpected: %+v\nGot: %+v", expectedStruct, myStruct)
	}
}
