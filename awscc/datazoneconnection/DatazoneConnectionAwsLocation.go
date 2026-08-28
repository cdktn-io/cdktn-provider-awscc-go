// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionAwsLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_connection#access_role DatazoneConnection#access_role}.
	AccessRole *string `field:"optional" json:"accessRole" yaml:"accessRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_connection#aws_account_id DatazoneConnection#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_connection#aws_region DatazoneConnection#aws_region}.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_connection#iam_connection_id DatazoneConnection#iam_connection_id}.
	IamConnectionId *string `field:"optional" json:"iamConnectionId" yaml:"iamConnectionId"`
}

