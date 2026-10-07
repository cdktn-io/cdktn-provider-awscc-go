// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheglobalreplicationgroup


type ElasticacheGlobalReplicationGroupTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_global_replication_group#key ElasticacheGlobalReplicationGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_global_replication_group#value ElasticacheGlobalReplicationGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

