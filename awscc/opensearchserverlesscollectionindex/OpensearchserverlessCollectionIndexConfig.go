// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollectionindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserverlessCollectionIndexConfig struct {
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
	// The identifier of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchserverless_collection_index#collection_index_id OpensearchserverlessCollectionIndex#collection_index_id}
	CollectionIndexId *string `field:"required" json:"collectionIndexId" yaml:"collectionIndexId"`
	// The name of the collection index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchserverless_collection_index#index_name OpensearchserverlessCollectionIndex#index_name}
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The Mappings for the collection index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/opensearchserverless_collection_index#index_schema OpensearchserverlessCollectionIndex#index_schema}
	IndexSchema *string `field:"optional" json:"indexSchema" yaml:"indexSchema"`
}

