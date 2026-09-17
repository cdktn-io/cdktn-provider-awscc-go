// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelLoggingInfoCloudwatchLogs struct {
	// Whether CloudWatch Logs logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#enabled MskChannel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#log_group MskChannel#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

