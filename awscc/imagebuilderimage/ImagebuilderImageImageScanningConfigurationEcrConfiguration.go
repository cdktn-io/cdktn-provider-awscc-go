// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimage


type ImagebuilderImageImageScanningConfigurationEcrConfiguration struct {
	// Tags for Image Builder to apply the output container image that is scanned.
	//
	// Tags can help you identify and manage your scanned images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_image#container_tags ImagebuilderImage#container_tags}
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// The name of the container repository that Amazon Inspector scans to identify findings for your container images.
	//
	// The name includes the path for the repository location. If you don’t provide this information, Image Builder creates a repository in your account named image-builder-image-scanning-repository to use for vulnerability scans for your output container images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_image#repository_name ImagebuilderImage#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

