// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert


type CloudwatchAlertTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#key CloudwatchAlert#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_alert#value CloudwatchAlert#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

