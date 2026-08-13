// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2filter


type Inspectorv2FilterFilterCriteriaLastObservedAt struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/inspectorv2_filter#end_inclusive Inspectorv2Filter#end_inclusive}.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/inspectorv2_filter#start_inclusive Inspectorv2Filter#start_inclusive}.
	StartInclusive *float64 `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

