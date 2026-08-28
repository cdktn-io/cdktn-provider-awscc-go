// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2filter


type Inspectorv2FilterFilterCriteriaPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_filter#begin_inclusive Inspectorv2Filter#begin_inclusive}.
	BeginInclusive *float64 `field:"optional" json:"beginInclusive" yaml:"beginInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_filter#end_inclusive Inspectorv2Filter#end_inclusive}.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
}

