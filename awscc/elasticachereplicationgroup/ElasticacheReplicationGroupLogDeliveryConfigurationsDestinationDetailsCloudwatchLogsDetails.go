// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetailsCloudwatchLogsDetails struct {
	// The name of the CloudWatch Logs log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticache_replication_group#log_group ElasticacheReplicationGroup#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

