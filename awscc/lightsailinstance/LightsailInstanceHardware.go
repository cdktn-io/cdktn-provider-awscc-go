// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance


type LightsailInstanceHardware struct {
	// Disks attached to the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lightsail_instance#disks LightsailInstance#disks}
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
}

