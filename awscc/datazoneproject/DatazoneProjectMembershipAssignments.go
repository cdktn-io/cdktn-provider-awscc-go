// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneproject


type DatazoneProjectMembershipAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_project#designation DatazoneProject#designation}.
	Designation *string `field:"optional" json:"designation" yaml:"designation"`
	// The member of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_project#member DatazoneProject#member}
	Member *DatazoneProjectMembershipAssignmentsMember `field:"optional" json:"member" yaml:"member"`
}

