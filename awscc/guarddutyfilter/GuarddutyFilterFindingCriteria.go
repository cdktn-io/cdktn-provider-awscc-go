// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutyfilter


type GuarddutyFilterFindingCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/guardduty_filter#criterion GuarddutyFilter#criterion}.
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

