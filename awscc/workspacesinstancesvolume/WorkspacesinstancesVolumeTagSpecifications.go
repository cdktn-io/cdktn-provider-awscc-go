// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesvolume


type WorkspacesinstancesVolumeTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/workspacesinstances_volume#resource_type WorkspacesinstancesVolume#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/workspacesinstances_volume#tags WorkspacesinstancesVolume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

