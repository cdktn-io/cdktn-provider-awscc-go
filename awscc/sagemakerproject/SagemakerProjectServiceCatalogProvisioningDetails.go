// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject


type SagemakerProjectServiceCatalogProvisioningDetails struct {
	// The path identifier of the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#path_id SagemakerProject#path_id}
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// Service Catalog product identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#product_id SagemakerProject#product_id}
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The identifier of the provisioning artifact (also known as a version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#provisioning_artifact_id SagemakerProject#provisioning_artifact_id}
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// Parameters specified by the administrator that are required for provisioning the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#provisioning_parameters SagemakerProject#provisioning_parameters}
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
}

