// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamoidcprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamOidcProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_oidc_provider#client_id_list IamOidcProvider#client_id_list}.
	ClientIdList *[]*string `field:"optional" json:"clientIdList" yaml:"clientIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_oidc_provider#tags IamOidcProvider#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_oidc_provider#thumbprint_list IamOidcProvider#thumbprint_list}.
	ThumbprintList *[]*string `field:"optional" json:"thumbprintList" yaml:"thumbprintList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_oidc_provider#url IamOidcProvider#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

