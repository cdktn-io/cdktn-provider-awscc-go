// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneprojectmembership


type DatazoneProjectMembershipMember struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_project_membership#group_identifier DatazoneProjectMembership#group_identifier}.
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_project_membership#user_identifier DatazoneProjectMembership#user_identifier}.
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

