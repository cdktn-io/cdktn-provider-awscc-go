// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsreferencestore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsReferenceStoreConfig struct {
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
	// A name for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_reference_store#name OmicsReferenceStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_reference_store#description OmicsReferenceStore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Server-side encryption (SSE) settings for a store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_reference_store#sse_config OmicsReferenceStore#sse_config}
	SseConfig *OmicsReferenceStoreSseConfig `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_reference_store#tags OmicsReferenceStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

