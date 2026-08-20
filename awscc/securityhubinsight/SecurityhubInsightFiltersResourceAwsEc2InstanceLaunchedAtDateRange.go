// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange struct {
	// A date range unit for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_insight#unit SecurityhubInsight#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// A date range value for the date filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityhub_insight#value SecurityhubInsight#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

