// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment


type EvsEnvironmentLicenseInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evs_environment#solution_key EvsEnvironment#solution_key}.
	SolutionKey *string `field:"optional" json:"solutionKey" yaml:"solutionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/evs_environment#vsan_key EvsEnvironment#vsan_key}.
	VsanKey *string `field:"optional" json:"vsanKey" yaml:"vsanKey"`
}

