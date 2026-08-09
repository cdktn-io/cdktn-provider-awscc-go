// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserviceapplication


type OpensearchserviceApplicationIamIdentityCenterOptions struct {
	// Whether IAM Identity Center is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchservice_application#enabled OpensearchserviceApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the IAM Identity Center instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchservice_application#iam_identity_center_instance_arn OpensearchserviceApplication#iam_identity_center_instance_arn}
	IamIdentityCenterInstanceArn *string `field:"optional" json:"iamIdentityCenterInstanceArn" yaml:"iamIdentityCenterInstanceArn"`
	// The ARN of the IAM role for Identity Center application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchservice_application#iam_role_for_identity_center_application_arn OpensearchserviceApplication#iam_role_for_identity_center_application_arn}
	IamRoleForIdentityCenterApplicationArn *string `field:"optional" json:"iamRoleForIdentityCenterApplicationArn" yaml:"iamRoleForIdentityCenterApplicationArn"`
}

