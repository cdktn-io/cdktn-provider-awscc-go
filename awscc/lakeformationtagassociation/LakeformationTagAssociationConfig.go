// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lakeformationtagassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LakeformationTagAssociationConfig struct {
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
	// List of Lake Formation Tags to associate with the Lake Formation Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_tag_association#lf_tags LakeformationTagAssociation#lf_tags}
	LfTags interface{} `field:"required" json:"lfTags" yaml:"lfTags"`
	// Resource to tag with the Lake Formation Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lakeformation_tag_association#resource LakeformationTagAssociation#resource}
	Resource *LakeformationTagAssociationResource `field:"required" json:"resource" yaml:"resource"`
}

