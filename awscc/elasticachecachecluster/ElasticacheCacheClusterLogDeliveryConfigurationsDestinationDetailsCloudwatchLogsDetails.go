// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachecachecluster


type ElasticacheCacheClusterLogDeliveryConfigurationsDestinationDetailsCloudwatchLogsDetails struct {
	// The name of the CloudWatch Logs log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_cache_cluster#log_group ElasticacheCacheCluster#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

