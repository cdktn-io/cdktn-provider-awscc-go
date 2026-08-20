// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionResourceAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#cloudformation_action ApptestTestCase#cloudformation_action}.
	CloudformationAction *ApptestTestCaseStepsActionResourceActionCloudformationAction `field:"optional" json:"cloudformationAction" yaml:"cloudformationAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#m2_managed_application_action ApptestTestCase#m2_managed_application_action}.
	M2ManagedApplicationAction *ApptestTestCaseStepsActionResourceActionM2ManagedApplicationAction `field:"optional" json:"m2ManagedApplicationAction" yaml:"m2ManagedApplicationAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#m2_non_managed_application_action ApptestTestCase#m2_non_managed_application_action}.
	M2NonManagedApplicationAction *ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationAction `field:"optional" json:"m2NonManagedApplicationAction" yaml:"m2NonManagedApplicationAction"`
}

