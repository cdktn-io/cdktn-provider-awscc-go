// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesssecurityconfig


type OpensearchserverlessSecurityConfigIamIdentityCenterOptions struct {
	// Group attribute for this IAM Identity Center integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_security_config#group_attribute OpensearchserverlessSecurityConfig#group_attribute}
	GroupAttribute *string `field:"optional" json:"groupAttribute" yaml:"groupAttribute"`
	// The ARN of the IAM Identity Center instance used to integrate with OpenSearch Serverless.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_security_config#instance_arn OpensearchserverlessSecurityConfig#instance_arn}
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// User attribute for this IAM Identity Center integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_security_config#user_attribute OpensearchserverlessSecurityConfig#user_attribute}
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

