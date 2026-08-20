// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationproduct


type ServicecatalogCloudformationProductSourceConnectionConnectionParametersCodeStar struct {
	// The absolute path where the artifact resides within the repo and branch, formatted as "folder/file.json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#artifact_path ServicecatalogCloudformationProduct#artifact_path}
	ArtifactPath *string `field:"optional" json:"artifactPath" yaml:"artifactPath"`
	// The specific branch where the artifact resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#branch ServicecatalogCloudformationProduct#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The CodeStar ARN, which is the connection between AWS Service Catalog and the external repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#connection_arn ServicecatalogCloudformationProduct#connection_arn}
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
	// The specific repository where the product's artifact-to-be-synced resides, formatted as "Account/Repo.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#repository ServicecatalogCloudformationProduct#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

