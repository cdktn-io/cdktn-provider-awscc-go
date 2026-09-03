// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscachesnapshot


type ElasticacheServerlessCacheSnapshotTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticache_serverless_cache_snapshot#key ElasticacheServerlessCacheSnapshot#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticache_serverless_cache_snapshot#value ElasticacheServerlessCacheSnapshot#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

