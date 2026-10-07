// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomcontentassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomContentAssociationConfig struct {
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
	// The identifier of the associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content_association#association WisdomContentAssociation#association}
	Association *WisdomContentAssociationAssociation `field:"required" json:"association" yaml:"association"`
	// The type of association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content_association#association_type WisdomContentAssociation#association_type}
	AssociationType *string `field:"required" json:"associationType" yaml:"associationType"`
	// The identifier of the content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content_association#content_id WisdomContentAssociation#content_id}
	ContentId *string `field:"required" json:"contentId" yaml:"contentId"`
	// The identifier of the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content_association#knowledge_base_id WisdomContentAssociation#knowledge_base_id}
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content_association#tags WisdomContentAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

