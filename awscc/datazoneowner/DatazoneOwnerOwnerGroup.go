// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneowner


type DatazoneOwnerOwnerGroup struct {
	// The ID of the domain unit owners group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_owner#group_identifier DatazoneOwner#group_identifier}
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
}

