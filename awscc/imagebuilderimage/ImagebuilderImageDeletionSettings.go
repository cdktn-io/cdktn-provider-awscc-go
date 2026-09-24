// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageDeletionSettings struct {
	// The execution role to use for deleting the image, as well as underlying resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/imagebuilder_image#execution_role ImagebuilderImage#execution_role}
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
}

