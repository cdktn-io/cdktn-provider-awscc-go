// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimits struct {
	// The cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_serverless_cache#data_storage ElasticacheServerlessCache#data_storage}
	DataStorage *ElasticacheServerlessCacheCacheUsageLimitsDataStorage `field:"optional" json:"dataStorage" yaml:"dataStorage"`
	// The ECPU per second of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_serverless_cache#ecpu_per_second ElasticacheServerlessCache#ecpu_per_second}
	EcpuPerSecond *ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond `field:"optional" json:"ecpuPerSecond" yaml:"ecpuPerSecond"`
}

