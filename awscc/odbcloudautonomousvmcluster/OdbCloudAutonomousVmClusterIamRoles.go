// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster


type OdbCloudAutonomousVmClusterIamRoles struct {
	// The AWS integration configuration settings for the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/odb_cloud_autonomous_vm_cluster#aws_integration OdbCloudAutonomousVmCluster#aws_integration}
	AwsIntegration *string `field:"optional" json:"awsIntegration" yaml:"awsIntegration"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/odb_cloud_autonomous_vm_cluster#iam_role_arn OdbCloudAutonomousVmCluster#iam_role_arn}
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The current status of the AWS Identity and Access Management (IAM) service role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/odb_cloud_autonomous_vm_cluster#status OdbCloudAutonomousVmCluster#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

