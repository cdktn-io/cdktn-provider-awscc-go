// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsDataStorage struct {
	// The maximum cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The minimum cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_serverless_cache#minimum ElasticacheServerlessCache#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// The unit of cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_serverless_cache#unit ElasticacheServerlessCache#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

