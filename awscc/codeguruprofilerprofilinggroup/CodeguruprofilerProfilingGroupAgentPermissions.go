// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeguruprofilerprofilinggroup


type CodeguruprofilerProfilingGroupAgentPermissions struct {
	// The principals for the agent permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codeguruprofiler_profiling_group#principals CodeguruprofilerProfilingGroup#principals}
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

