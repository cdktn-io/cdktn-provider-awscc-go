// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionMainframeActionActionType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apptest_test_case#batch ApptestTestCase#batch}.
	Batch *ApptestTestCaseStepsActionMainframeActionActionTypeBatch `field:"optional" json:"batch" yaml:"batch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apptest_test_case#tn_3270 ApptestTestCase#tn_3270}.
	Tn3270 *ApptestTestCaseStepsActionMainframeActionActionTypeTn3270 `field:"optional" json:"tn3270" yaml:"tn3270"`
}

