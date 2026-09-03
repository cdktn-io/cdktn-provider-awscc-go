// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserverlesscollectiongroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserverlessCollectionGroupConfig struct {
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
	// The name of the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#name OpensearchserverlessCollectionGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether standby replicas are used for the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#standby_replicas OpensearchserverlessCollectionGroup#standby_replicas}
	StandbyReplicas *string `field:"required" json:"standbyReplicas" yaml:"standbyReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#capacity_limits OpensearchserverlessCollectionGroup#capacity_limits}.
	CapacityLimits *OpensearchserverlessCollectionGroupCapacityLimits `field:"optional" json:"capacityLimits" yaml:"capacityLimits"`
	// The description of the collection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#description OpensearchserverlessCollectionGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The generation of Amazon OpenSearch Serverless for the collection group. Valid values are CLASSIC and NEXTGEN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#generation OpensearchserverlessCollectionGroup#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchserverless_collection_group#tags OpensearchserverlessCollectionGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

