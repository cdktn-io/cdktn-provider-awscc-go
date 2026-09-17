// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebidentityprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceswebIdentityProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesweb_identity_provider#identity_provider_details WorkspaceswebIdentityProvider#identity_provider_details}.
	IdentityProviderDetails *map[string]*string `field:"required" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesweb_identity_provider#identity_provider_name WorkspaceswebIdentityProvider#identity_provider_name}.
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesweb_identity_provider#identity_provider_type WorkspaceswebIdentityProvider#identity_provider_type}.
	IdentityProviderType *string `field:"required" json:"identityProviderType" yaml:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesweb_identity_provider#portal_arn WorkspaceswebIdentityProvider#portal_arn}.
	PortalArn *string `field:"optional" json:"portalArn" yaml:"portalArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesweb_identity_provider#tags WorkspaceswebIdentityProvider#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

