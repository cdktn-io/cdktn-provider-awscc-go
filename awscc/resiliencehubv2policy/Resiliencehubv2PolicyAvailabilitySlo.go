// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2policy


type Resiliencehubv2PolicyAvailabilitySlo struct {
	// Availability target percentage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_policy#target Resiliencehubv2Policy#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
}

