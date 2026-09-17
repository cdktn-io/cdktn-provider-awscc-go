// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheglobalreplicationgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticacheGlobalReplicationGroupConfig struct {
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
	// The replication groups that comprise the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#members ElasticacheGlobalReplicationGroup#members}
	Members interface{} `field:"required" json:"members" yaml:"members"`
	// AutomaticFailoverEnabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#automatic_failover_enabled ElasticacheGlobalReplicationGroup#automatic_failover_enabled}
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// The cache node type of the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#cache_node_type ElasticacheGlobalReplicationGroup#cache_node_type}
	CacheNodeType *string `field:"optional" json:"cacheNodeType" yaml:"cacheNodeType"`
	// Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#cache_parameter_group_name ElasticacheGlobalReplicationGroup#cache_parameter_group_name}
	CacheParameterGroupName *string `field:"optional" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
	// The engine of the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#engine ElasticacheGlobalReplicationGroup#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The engine version of the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#engine_version ElasticacheGlobalReplicationGroup#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Indicates the number of node groups in the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#global_node_group_count ElasticacheGlobalReplicationGroup#global_node_group_count}
	GlobalNodeGroupCount *float64 `field:"optional" json:"globalNodeGroupCount" yaml:"globalNodeGroupCount"`
	// The optional description of the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#global_replication_group_description ElasticacheGlobalReplicationGroup#global_replication_group_description}
	GlobalReplicationGroupDescription *string `field:"optional" json:"globalReplicationGroupDescription" yaml:"globalReplicationGroupDescription"`
	// The suffix name of a Global Datastore.
	//
	// Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#global_replication_group_id_suffix ElasticacheGlobalReplicationGroup#global_replication_group_id_suffix}
	GlobalReplicationGroupIdSuffix *string `field:"optional" json:"globalReplicationGroupIdSuffix" yaml:"globalReplicationGroupIdSuffix"`
	// Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_global_replication_group#regional_configurations ElasticacheGlobalReplicationGroup#regional_configurations}
	RegionalConfigurations interface{} `field:"optional" json:"regionalConfigurations" yaml:"regionalConfigurations"`
}

