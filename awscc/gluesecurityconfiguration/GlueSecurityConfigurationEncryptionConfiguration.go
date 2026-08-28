// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesecurityconfiguration


type GlueSecurityConfigurationEncryptionConfiguration struct {
	// The encryption configuration for Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_security_configuration#cloudwatch_encryption GlueSecurityConfiguration#cloudwatch_encryption}
	CloudwatchEncryption *GlueSecurityConfigurationEncryptionConfigurationCloudwatchEncryption `field:"optional" json:"cloudwatchEncryption" yaml:"cloudwatchEncryption"`
	// The encryption configuration for job bookmarks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_security_configuration#job_bookmarks_encryption GlueSecurityConfiguration#job_bookmarks_encryption}
	JobBookmarksEncryption *GlueSecurityConfigurationEncryptionConfigurationJobBookmarksEncryption `field:"optional" json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_security_configuration#s3_encryptions GlueSecurityConfiguration#s3_encryptions}
	S3Encryptions interface{} `field:"optional" json:"s3Encryptions" yaml:"s3Encryptions"`
}

