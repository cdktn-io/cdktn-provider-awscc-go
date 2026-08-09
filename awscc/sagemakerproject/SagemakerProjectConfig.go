// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerProjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#project_name SagemakerProject#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The description of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#project_description SagemakerProject#project_description}
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// Provisioned ServiceCatalog  Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#service_catalog_provisioned_product_details SagemakerProject#service_catalog_provisioned_product_details}
	ServiceCatalogProvisionedProductDetails *SagemakerProjectServiceCatalogProvisionedProductDetails `field:"optional" json:"serviceCatalogProvisionedProductDetails" yaml:"serviceCatalogProvisionedProductDetails"`
	// Input ServiceCatalog Provisioning Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#service_catalog_provisioning_details SagemakerProject#service_catalog_provisioning_details}
	ServiceCatalogProvisioningDetails *SagemakerProjectServiceCatalogProvisioningDetails `field:"optional" json:"serviceCatalogProvisioningDetails" yaml:"serviceCatalogProvisioningDetails"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#tags SagemakerProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An array of template providers associated with the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_project#template_provider_details SagemakerProject#template_provider_details}
	TemplateProviderDetails interface{} `field:"optional" json:"templateProviderDetails" yaml:"templateProviderDetails"`
}

