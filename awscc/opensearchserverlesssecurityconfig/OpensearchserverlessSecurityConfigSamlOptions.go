// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesssecurityconfig


type OpensearchserverlessSecurityConfigSamlOptions struct {
	// Group attribute for this saml integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchserverless_security_config#group_attribute OpensearchserverlessSecurityConfig#group_attribute}
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// The XML saml provider metadata document that you want to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchserverless_security_config#metadata OpensearchserverlessSecurityConfig#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Custom entity id attribute to override default entity id for this saml integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchserverless_security_config#open_search_serverless_entity_id OpensearchserverlessSecurityConfig#open_search_serverless_entity_id}
	OpenSearchServerlessEntityId *string `field:"optional" json:"openSearchServerlessEntityId" yaml:"openSearchServerlessEntityId"`
	// Defines the session timeout in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchserverless_security_config#session_timeout OpensearchserverlessSecurityConfig#session_timeout}
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// Custom attribute for this saml integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchserverless_security_config#user_attribute OpensearchserverlessSecurityConfig#user_attribute}
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

