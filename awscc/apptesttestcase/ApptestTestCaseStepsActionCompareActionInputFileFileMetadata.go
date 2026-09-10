// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase


type ApptestTestCaseStepsActionCompareActionInputFileFileMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apptest_test_case#database_cdc ApptestTestCase#database_cdc}.
	DatabaseCdc *ApptestTestCaseStepsActionCompareActionInputFileFileMetadataDatabaseCdc `field:"optional" json:"databaseCdc" yaml:"databaseCdc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apptest_test_case#data_sets ApptestTestCase#data_sets}.
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
}

