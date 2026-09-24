// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionResourceActionM2ManagedApplicationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apptest_test_case#action_type ApptestTestCase#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apptest_test_case#properties ApptestTestCase#properties}.
	Properties *ApptestTestCaseStepsActionResourceActionM2ManagedApplicationActionProperties `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/apptest_test_case#resource ApptestTestCase#resource}.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

