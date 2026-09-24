// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerSigV4AuthorizationConfig struct {
	// Custom headers for the SigV4 MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#custom_headers DevopsagentService#custom_headers}
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// IAM role ARN to assume for SigV4 signing.
	//
	// Optional - when omitted, credentials are resolved at runtime via a monitor account association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#mcp_role_arn DevopsagentService#mcp_role_arn}
	McpRoleArn *string `field:"optional" json:"mcpRoleArn" yaml:"mcpRoleArn"`
	// AWS region for SigV4 signing. Use '*' for SigV4a multi-region signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#region DevopsagentService#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Deprecated - use McpRoleArn instead. IAM role ARN to assume for SigV4 signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#role_arn DevopsagentService#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// AWS service name for SigV4 signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#service DevopsagentService#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

