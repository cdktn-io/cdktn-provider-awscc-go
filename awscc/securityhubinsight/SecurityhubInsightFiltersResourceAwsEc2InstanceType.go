// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersResourceAwsEc2InstanceType struct {
	// The condition to apply to a string value when filtering Security Hub findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#comparison SecurityhubInsight#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Non-empty string definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#value SecurityhubInsight#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

