// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceAlternateDataSourceParametersS3Parameters struct {
	// <p>Amazon S3 manifest file location.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#manifest_file_location QuicksightDataSource#manifest_file_location}
	ManifestFileLocation *QuicksightDataSourceAlternateDataSourceParametersS3ParametersManifestFileLocation `field:"optional" json:"manifestFileLocation" yaml:"manifestFileLocation"`
	// <p>Use the <code>RoleArn</code> structure to override an account-wide role for a specific S3 data source.
	//
	// For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use <code>RoleArn</code> to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_source#role_arn QuicksightDataSource#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

