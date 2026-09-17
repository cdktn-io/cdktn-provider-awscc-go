// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationsresourcepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrganizationsResourcePolicyConfig struct {
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
	// The policy document.
	//
	// For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/organizations_resource_policy#content OrganizationsResourcePolicy#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// A list of tags that you want to attach to the resource policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/organizations_resource_policy#tags OrganizationsResourcePolicy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

