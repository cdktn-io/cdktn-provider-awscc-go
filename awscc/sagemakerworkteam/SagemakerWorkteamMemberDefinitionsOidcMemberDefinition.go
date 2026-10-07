// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamMemberDefinitionsOidcMemberDefinition struct {
	// A list of OIDC group names whose members will be part of this workteam.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#oidc_groups SagemakerWorkteam#oidc_groups}
	OidcGroups *[]*string `field:"optional" json:"oidcGroups" yaml:"oidcGroups"`
}

