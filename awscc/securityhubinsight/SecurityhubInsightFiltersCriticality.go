// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersCriticality struct {
	// The equal-to condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#eq SecurityhubInsight#eq}
	Eq *float64 `field:"optional" json:"eq" yaml:"eq"`
	// The greater-than-equal condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#gte SecurityhubInsight#gte}
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// The less-than-equal condition to be applied to a single field when querying for findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityhub_insight#lte SecurityhubInsight#lte}
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
}

