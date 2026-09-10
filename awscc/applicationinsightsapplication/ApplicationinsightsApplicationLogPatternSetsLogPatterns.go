// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationLogPatternSetsLogPatterns struct {
	// The log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationinsights_application#pattern ApplicationinsightsApplication#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationinsights_application#pattern_name ApplicationinsightsApplication#pattern_name}
	PatternName *string `field:"optional" json:"patternName" yaml:"patternName"`
	// Rank of the log pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationinsights_application#rank ApplicationinsightsApplication#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

