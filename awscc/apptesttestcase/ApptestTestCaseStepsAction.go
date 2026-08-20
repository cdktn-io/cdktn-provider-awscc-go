// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#compare_action ApptestTestCase#compare_action}.
	CompareAction *ApptestTestCaseStepsActionCompareAction `field:"optional" json:"compareAction" yaml:"compareAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#mainframe_action ApptestTestCase#mainframe_action}.
	MainframeAction *ApptestTestCaseStepsActionMainframeAction `field:"optional" json:"mainframeAction" yaml:"mainframeAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#resource_action ApptestTestCase#resource_action}.
	ResourceAction *ApptestTestCaseStepsActionResourceAction `field:"optional" json:"resourceAction" yaml:"resourceAction"`
}

