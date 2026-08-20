// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightlimitsprofile


type QuicksightLimitsProfileResourceLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_limits_profile#max_value QuicksightLimitsProfile#max_value}.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_limits_profile#unit QuicksightLimitsProfile#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

