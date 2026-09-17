// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdates struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#acl_configuration AthenaWorkGroup#acl_configuration}
	AclConfiguration *AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesAclConfiguration `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#encryption_configuration AthenaWorkGroup#encryption_configuration}
	EncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationUpdatesResultConfigurationUpdatesEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The AWS account ID of the owner of S3 bucket where query results are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#expected_bucket_owner AthenaWorkGroup#expected_bucket_owner}
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/.
	//
	// To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#output_location AthenaWorkGroup#output_location}
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#remove_acl_configuration AthenaWorkGroup#remove_acl_configuration}.
	RemoveAclConfiguration interface{} `field:"optional" json:"removeAclConfiguration" yaml:"removeAclConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#remove_encryption_configuration AthenaWorkGroup#remove_encryption_configuration}.
	RemoveEncryptionConfiguration interface{} `field:"optional" json:"removeEncryptionConfiguration" yaml:"removeEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#remove_expected_bucket_owner AthenaWorkGroup#remove_expected_bucket_owner}.
	RemoveExpectedBucketOwner interface{} `field:"optional" json:"removeExpectedBucketOwner" yaml:"removeExpectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#remove_output_location AthenaWorkGroup#remove_output_location}.
	RemoveOutputLocation interface{} `field:"optional" json:"removeOutputLocation" yaml:"removeOutputLocation"`
}

