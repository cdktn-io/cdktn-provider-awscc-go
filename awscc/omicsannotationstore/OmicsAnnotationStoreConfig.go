// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsannotationstore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsAnnotationStoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#name OmicsAnnotationStore#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#store_format OmicsAnnotationStore#store_format}.
	StoreFormat *string `field:"required" json:"storeFormat" yaml:"storeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#description OmicsAnnotationStore#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#reference OmicsAnnotationStore#reference}.
	Reference *OmicsAnnotationStoreReference `field:"optional" json:"reference" yaml:"reference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#sse_config OmicsAnnotationStore#sse_config}.
	SseConfig *OmicsAnnotationStoreSseConfig `field:"optional" json:"sseConfig" yaml:"sseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#store_options OmicsAnnotationStore#store_options}.
	StoreOptions *OmicsAnnotationStoreStoreOptions `field:"optional" json:"storeOptions" yaml:"storeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_annotation_store#tags OmicsAnnotationStore#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

