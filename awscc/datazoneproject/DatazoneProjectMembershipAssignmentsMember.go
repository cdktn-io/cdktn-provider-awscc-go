// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneproject


type DatazoneProjectMembershipAssignmentsMember struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_project#group_identifier DatazoneProject#group_identifier}.
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_project#user_identifier DatazoneProject#user_identifier}.
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

