// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster


type OdbCloudVmClusterIamRoles struct {
	// The AWS integration configuration settings for the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_cloud_vm_cluster#aws_integration OdbCloudVmCluster#aws_integration}
	AwsIntegration *string `field:"optional" json:"awsIntegration" yaml:"awsIntegration"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_cloud_vm_cluster#iam_role_arn OdbCloudVmCluster#iam_role_arn}
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The current status of the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_cloud_vm_cluster#status OdbCloudVmCluster#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

