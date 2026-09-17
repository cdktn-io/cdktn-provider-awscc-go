// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment


type EvsEnvironmentHosts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment#dedicated_host_id EvsEnvironment#dedicated_host_id}.
	DedicatedHostId *string `field:"optional" json:"dedicatedHostId" yaml:"dedicatedHostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment#host_name EvsEnvironment#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment#instance_type EvsEnvironment#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment#key_name EvsEnvironment#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment#placement_group_id EvsEnvironment#placement_group_id}.
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
}

