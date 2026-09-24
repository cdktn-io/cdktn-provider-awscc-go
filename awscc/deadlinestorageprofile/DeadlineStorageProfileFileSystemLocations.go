// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinestorageprofile


type DeadlineStorageProfileFileSystemLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/deadline_storage_profile#name DeadlineStorageProfile#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/deadline_storage_profile#path DeadlineStorageProfile#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/deadline_storage_profile#type DeadlineStorageProfile#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

