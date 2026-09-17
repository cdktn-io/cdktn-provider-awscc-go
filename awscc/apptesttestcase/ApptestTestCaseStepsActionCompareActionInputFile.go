// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionCompareActionInputFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apptest_test_case#file_metadata ApptestTestCase#file_metadata}.
	FileMetadata *ApptestTestCaseStepsActionCompareActionInputFileFileMetadata `field:"optional" json:"fileMetadata" yaml:"fileMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apptest_test_case#source_location ApptestTestCase#source_location}.
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apptest_test_case#target_location ApptestTestCase#target_location}.
	TargetLocation *string `field:"optional" json:"targetLocation" yaml:"targetLocation"`
}

