// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamMemberDefinitions struct {
	// The Amazon Cognito user group that is part of the work team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/sagemaker_workteam#cognito_member_definition SagemakerWorkteam#cognito_member_definition}
	CognitoMemberDefinition *SagemakerWorkteamMemberDefinitionsCognitoMemberDefinition `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// A list user groups that exist in your OIDC Identity Provider (IdP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/sagemaker_workteam#oidc_member_definition SagemakerWorkteam#oidc_member_definition}
	OidcMemberDefinition *SagemakerWorkteamMemberDefinitionsOidcMemberDefinition `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

