// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakedatalake


type SecuritylakeDataLakeReplicationConfiguration struct {
	// Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets.
	//
	// Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securitylake_data_lake#regions SecuritylakeDataLake#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Replication settings for the Amazon S3 buckets.
	//
	// This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securitylake_data_lake#role_arn SecuritylakeDataLake#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

