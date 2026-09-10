// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypoolroleattachment


type CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfigurationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool_role_attachment#claim CognitoIdentityPoolRoleAttachment#claim}.
	Claim *string `field:"optional" json:"claim" yaml:"claim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool_role_attachment#match_type CognitoIdentityPoolRoleAttachment#match_type}.
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool_role_attachment#role_arn CognitoIdentityPoolRoleAttachment#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cognito_identity_pool_role_attachment#value CognitoIdentityPoolRoleAttachment#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

