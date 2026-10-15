// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace


type DevopsagentAgentSpaceOperatorAppIdp struct {
	// The OIDC client identifier for the external identity provider application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_agent_space#client_id DevopsagentAgentSpace#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OIDC client secret for the external identity provider application.
	//
	// Use a Secrets Manager dynamic reference to avoid storing the secret in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_agent_space#client_secret DevopsagentAgentSpace#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer URL for the external identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_agent_space#issuer_url DevopsagentAgentSpace#issuer_url}
	IssuerUrl *string `field:"optional" json:"issuerUrl" yaml:"issuerUrl"`
	// The IAM role that Operator App users assume to access DevOps Agent APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_agent_space#operator_app_role_arn DevopsagentAgentSpace#operator_app_role_arn}
	OperatorAppRoleArn *string `field:"optional" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// The external identity provider type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_agent_space#provider DevopsagentAgentSpace#provider}
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

