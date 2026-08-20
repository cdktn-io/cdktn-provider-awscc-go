// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionMainframeActionActionTypeBatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#batch_job_name ApptestTestCase#batch_job_name}.
	BatchJobName *string `field:"optional" json:"batchJobName" yaml:"batchJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#batch_job_parameters ApptestTestCase#batch_job_parameters}.
	BatchJobParameters *map[string]*string `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apptest_test_case#export_data_set_names ApptestTestCase#export_data_set_names}.
	ExportDataSetNames *[]*string `field:"optional" json:"exportDataSetNames" yaml:"exportDataSetNames"`
}

