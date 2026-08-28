// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlessaccesspolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserverlessAccessPolicyConfig struct {
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
	// The name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchserverless_access_policy#name OpensearchserverlessAccessPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The JSON policy document that is the content for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchserverless_access_policy#policy OpensearchserverlessAccessPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The possible types for the access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchserverless_access_policy#type OpensearchserverlessAccessPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchserverless_access_policy#description OpensearchserverlessAccessPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

