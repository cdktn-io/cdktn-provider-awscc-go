// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2filter


type Inspectorv2FilterFilterCriteriaVulnerablePackagesFilePath struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_filter#comparison Inspectorv2Filter#comparison}.
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/inspectorv2_filter#value Inspectorv2Filter#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

