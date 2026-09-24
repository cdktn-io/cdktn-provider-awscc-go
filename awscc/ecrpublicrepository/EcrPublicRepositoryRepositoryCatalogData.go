// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrpublicrepository


type EcrPublicRepositoryRepositoryCatalogData struct {
	// Provide a detailed description of the repository.
	//
	// Identify what is included in the repository, any licensing details, or other relevant information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecr_public_repository#about_text EcrPublicRepository#about_text}
	AboutText *string `field:"optional" json:"aboutText" yaml:"aboutText"`
	// Select the system architectures that the images in your repository are compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecr_public_repository#architectures EcrPublicRepository#architectures}
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Select the operating systems that the images in your repository are compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecr_public_repository#operating_systems EcrPublicRepository#operating_systems}
	OperatingSystems *[]*string `field:"optional" json:"operatingSystems" yaml:"operatingSystems"`
	// The description of the public repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecr_public_repository#repository_description EcrPublicRepository#repository_description}
	RepositoryDescription *string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
	// Provide detailed information about how to use the images in the repository.
	//
	// This provides context, support information, and additional usage details for users of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecr_public_repository#usage_text EcrPublicRepository#usage_text}
	UsageText *string `field:"optional" json:"usageText" yaml:"usageText"`
}

