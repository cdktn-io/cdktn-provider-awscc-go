// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesssecurityconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserverlessSecurityConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Security config description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#description OpensearchserverlessSecurityConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Describe IAM federation options in form of key value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#iam_federation_options OpensearchserverlessSecurityConfig#iam_federation_options}
	IamFederationOptions *OpensearchserverlessSecurityConfigIamFederationOptions `field:"optional" json:"iamFederationOptions" yaml:"iamFederationOptions"`
	// Describes IAM Identity Center options for an OpenSearch Serverless security configuration in the form of a key-value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#iam_identity_center_options OpensearchserverlessSecurityConfig#iam_identity_center_options}
	IamIdentityCenterOptions *OpensearchserverlessSecurityConfigIamIdentityCenterOptions `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// The friendly name of the security config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#name OpensearchserverlessSecurityConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Describes saml options in form of key value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#saml_options OpensearchserverlessSecurityConfig#saml_options}
	SamlOptions *OpensearchserverlessSecurityConfigSamlOptions `field:"optional" json:"samlOptions" yaml:"samlOptions"`
	// Config type for security config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/opensearchserverless_security_config#type OpensearchserverlessSecurityConfig#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

