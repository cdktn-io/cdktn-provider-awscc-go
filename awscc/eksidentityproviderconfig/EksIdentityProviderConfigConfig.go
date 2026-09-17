// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksidentityproviderconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksIdentityProviderConfigConfig struct {
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
	// The name of the identity provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_identity_provider_config#cluster_name EksIdentityProviderConfig#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The type of the identity provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_identity_provider_config#type EksIdentityProviderConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the OIDC provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_identity_provider_config#identity_provider_config_name EksIdentityProviderConfig#identity_provider_config_name}
	IdentityProviderConfigName *string `field:"optional" json:"identityProviderConfigName" yaml:"identityProviderConfigName"`
	// An object representing an OpenID Connect (OIDC) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_identity_provider_config#oidc EksIdentityProviderConfig#oidc}
	Oidc *EksIdentityProviderConfigOidc `field:"optional" json:"oidc" yaml:"oidc"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/eks_identity_provider_config#tags EksIdentityProviderConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

