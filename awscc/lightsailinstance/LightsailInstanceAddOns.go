// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance


type LightsailInstanceAddOns struct {
	// The add-on type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_instance#add_on_type LightsailInstance#add_on_type}
	AddOnType *string `field:"optional" json:"addOnType" yaml:"addOnType"`
	// An object that represents additional parameters when enabling or modifying the automatic snapshot add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_instance#auto_snapshot_add_on_request LightsailInstance#auto_snapshot_add_on_request}
	AutoSnapshotAddOnRequest *LightsailInstanceAddOnsAutoSnapshotAddOnRequest `field:"optional" json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// Status of the Addon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_instance#status LightsailInstance#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

