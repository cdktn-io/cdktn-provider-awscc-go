// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesssecurityconfig


type OpensearchserverlessSecurityConfigIamFederationOptions struct {
	// Group attribute for this IAM federation integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchserverless_security_config#group_attribute OpensearchserverlessSecurityConfig#group_attribute}
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// User attribute for this IAM federation integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchserverless_security_config#user_attribute OpensearchserverlessSecurityConfig#user_attribute}
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

