// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomcontent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomContentConfig struct {
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
	// The identifier of the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#knowledge_base_id WisdomContent#knowledge_base_id}
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#name WisdomContent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A key/value map to store attributes without affecting tagging or recommendations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#metadata WisdomContent#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The URI you want to use for the article.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#override_link_out_uri WisdomContent#override_link_out_uri}
	OverrideLinkOutUri *string `field:"optional" json:"overrideLinkOutUri" yaml:"overrideLinkOutUri"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#tags WisdomContent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The title of the content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#title WisdomContent#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A pointer to the uploaded asset. This value is returned by StartContentUpload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/wisdom_content#upload_id WisdomContent#upload_id}
	UploadId *string `field:"optional" json:"uploadId" yaml:"uploadId"`
}

