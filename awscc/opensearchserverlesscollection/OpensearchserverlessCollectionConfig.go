// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserverlessCollectionConfig struct {
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
	// The name of the collection.
	//
	// The name must meet the following criteria:
	// Unique to your account and AWS Region
	// Starts with a lowercase letter
	// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
	// Contains between 3 and 64 characters
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#name OpensearchserverlessCollection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the collection group to associate with the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#collection_group_name OpensearchserverlessCollection#collection_group_name}
	CollectionGroupName *string `field:"optional" json:"collectionGroupName" yaml:"collectionGroupName"`
	// The deletion protection state of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#deletion_protection OpensearchserverlessCollection#deletion_protection}
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The description of the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#description OpensearchserverlessCollection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption settings for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#encryption_config OpensearchserverlessCollection#encryption_config}
	EncryptionConfig *OpensearchserverlessCollectionEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The possible standby replicas for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#standby_replicas OpensearchserverlessCollection#standby_replicas}
	StandbyReplicas *string `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
	// List of tags to be added to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#tags OpensearchserverlessCollection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The possible types for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#type OpensearchserverlessCollection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Vector search configuration options for the collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection#vector_options OpensearchserverlessCollection#vector_options}
	VectorOptions *OpensearchserverlessCollectionVectorOptions `field:"optional" json:"vectorOptions" yaml:"vectorOptions"`
}

