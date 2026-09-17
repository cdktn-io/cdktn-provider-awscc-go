// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraysamplingrule


type XraySamplingRuleSamplingRuleRecord struct {
	// When the rule was created, in Unix time seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/xray_sampling_rule#created_at XraySamplingRule#created_at}
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// When the rule was modified, in Unix time seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/xray_sampling_rule#modified_at XraySamplingRule#modified_at}
	ModifiedAt *string `field:"optional" json:"modifiedAt" yaml:"modifiedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/xray_sampling_rule#sampling_rule XraySamplingRule#sampling_rule}.
	SamplingRule *XraySamplingRuleSamplingRuleRecordSamplingRule `field:"optional" json:"samplingRule" yaml:"samplingRule"`
}

