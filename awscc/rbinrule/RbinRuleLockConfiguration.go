// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rbinrule


type RbinRuleLockConfiguration struct {
	// The unit of time in which to measure the unlock delay.
	//
	// Currently, the unlock delay can be measure only in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rbin_rule#unlock_delay_unit RbinRule#unlock_delay_unit}
	UnlockDelayUnit *string `field:"optional" json:"unlockDelayUnit" yaml:"unlockDelayUnit"`
	// The unlock delay period, measured in the unit specified for UnlockDelayUnit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rbin_rule#unlock_delay_value RbinRule#unlock_delay_value}
	UnlockDelayValue *float64 `field:"optional" json:"unlockDelayValue" yaml:"unlockDelayValue"`
}

