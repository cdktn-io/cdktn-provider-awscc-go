// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceVolumes struct {
	// The device name (for example, /dev/sdh or xvdh).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_instance#device Ec2Instance#device}
	Device *string `field:"optional" json:"device" yaml:"device"`
	// The ID of the EBS volume. The volume and instance must be within the same Availability Zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_instance#volume_id Ec2Instance#volume_id}
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

