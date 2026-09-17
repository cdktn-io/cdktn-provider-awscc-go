// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessinstance


type Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogs struct {
	// Indicates whether logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_instance#enabled Ec2VerifiedAccessInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the CloudWatch Logs log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_instance#log_group Ec2VerifiedAccessInstance#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

