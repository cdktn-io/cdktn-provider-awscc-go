// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbmultiregioncluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MemorydbMultiRegionClusterConfig struct {
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
	// The compute and memory capacity of the nodes in the multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#node_type MemorydbMultiRegionCluster#node_type}
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Description of the multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#description MemorydbMultiRegionCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The engine type used by the multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#engine MemorydbMultiRegionCluster#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The Redis engine version used by the multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#engine_version MemorydbMultiRegionCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The name of the Multi Region cluster.
	//
	// This value must be unique as it also serves as the multi region cluster identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#multi_region_cluster_name_suffix MemorydbMultiRegionCluster#multi_region_cluster_name_suffix}
	MultiRegionClusterNameSuffix *string `field:"optional" json:"multiRegionClusterNameSuffix" yaml:"multiRegionClusterNameSuffix"`
	// The name of the parameter group associated with the multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#multi_region_parameter_group_name MemorydbMultiRegionCluster#multi_region_parameter_group_name}
	MultiRegionParameterGroupName *string `field:"optional" json:"multiRegionParameterGroupName" yaml:"multiRegionParameterGroupName"`
	// The number of shards the multi region cluster will contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#num_shards MemorydbMultiRegionCluster#num_shards}
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// An array of key-value pairs to apply to this multi region cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#tags MemorydbMultiRegionCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#tls_enabled MemorydbMultiRegionCluster#tls_enabled}
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
	// An enum string value that determines the update strategy for scaling.
	//
	// Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/memorydb_multi_region_cluster#update_strategy MemorydbMultiRegionCluster#update_strategy}
	UpdateStrategy *string `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
}

