// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2filter


type Inspectorv2FilterFilterCriteriaInspectorScore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_filter#lower_inclusive Inspectorv2Filter#lower_inclusive}.
	LowerInclusive *float64 `field:"optional" json:"lowerInclusive" yaml:"lowerInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/inspectorv2_filter#upper_inclusive Inspectorv2Filter#upper_inclusive}.
	UpperInclusive *float64 `field:"optional" json:"upperInclusive" yaml:"upperInclusive"`
}

