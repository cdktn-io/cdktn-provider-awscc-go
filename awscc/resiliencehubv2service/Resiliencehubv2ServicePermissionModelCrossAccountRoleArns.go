// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServicePermissionModelCrossAccountRoleArns struct {
	// ARN of the cross-account IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/resiliencehubv2_service#cross_account_role_arn Resiliencehubv2Service#cross_account_role_arn}
	CrossAccountRoleArn *string `field:"optional" json:"crossAccountRoleArn" yaml:"crossAccountRoleArn"`
	// External ID for cross-account access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/resiliencehubv2_service#external_id Resiliencehubv2Service#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

