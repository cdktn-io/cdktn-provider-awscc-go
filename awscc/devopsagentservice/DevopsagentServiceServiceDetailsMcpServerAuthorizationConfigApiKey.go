// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerAuthorizationConfigApiKey struct {
	// HTTP header name to send the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#api_key_header DevopsagentService#api_key_header}
	ApiKeyHeader *string `field:"optional" json:"apiKeyHeader" yaml:"apiKeyHeader"`
	// User friendly API key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#api_key_name DevopsagentService#api_key_name}
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// API key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#api_key_value DevopsagentService#api_key_value}
	ApiKeyValue *string `field:"optional" json:"apiKeyValue" yaml:"apiKeyValue"`
}

