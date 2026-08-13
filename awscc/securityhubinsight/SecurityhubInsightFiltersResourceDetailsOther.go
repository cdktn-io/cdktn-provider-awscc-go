// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersResourceDetailsOther struct {
	// The condition to apply to the key value when filtering Security Hub findings with a map filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#comparison SecurityhubInsight#comparison}
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Non-empty string definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#key SecurityhubInsight#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Non-empty string definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#value SecurityhubInsight#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

