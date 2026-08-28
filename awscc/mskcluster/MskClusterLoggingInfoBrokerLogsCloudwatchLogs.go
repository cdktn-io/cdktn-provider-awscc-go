// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterLoggingInfoBrokerLogsCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#enabled MskCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_cluster#log_group MskCluster#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

