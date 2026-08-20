// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionMainframeActionActionTypeTn3270 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#export_data_set_names ApptestTestCase#export_data_set_names}.
	ExportDataSetNames *[]*string `field:"optional" json:"exportDataSetNames" yaml:"exportDataSetNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#script ApptestTestCase#script}.
	Script *ApptestTestCaseStepsActionMainframeActionActionTypeTn3270Script `field:"optional" json:"script" yaml:"script"`
}

