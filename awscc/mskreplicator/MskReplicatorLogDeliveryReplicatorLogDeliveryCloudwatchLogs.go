// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorLogDeliveryReplicatorLogDeliveryCloudwatchLogs struct {
	// Whether log delivery to CloudWatch Logs is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#enabled MskReplicator#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/msk_replicator#log_group MskReplicator#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

