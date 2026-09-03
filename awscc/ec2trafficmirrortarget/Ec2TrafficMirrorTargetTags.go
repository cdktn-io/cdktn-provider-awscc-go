// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2trafficmirrortarget


type Ec2TrafficMirrorTargetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_traffic_mirror_target#key Ec2TrafficMirrorTarget#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_traffic_mirror_target#value Ec2TrafficMirrorTarget#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

