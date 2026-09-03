// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package m2environment


type M2EnvironmentHighAvailabilityConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/m2_environment#desired_capacity M2Environment#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
}

