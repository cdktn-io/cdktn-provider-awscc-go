// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationproduct

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogCloudformationProductConfig struct {
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
	// The name of the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#name ServicecatalogCloudformationProduct#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The owner of the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#owner ServicecatalogCloudformationProduct#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#accept_language ServicecatalogCloudformationProduct#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#description ServicecatalogCloudformationProduct#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#distributor ServicecatalogCloudformationProduct#distributor}
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// The type of product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#product_type ServicecatalogCloudformationProduct#product_type}
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
	// The configuration of the provisioning artifact (also known as a version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#provisioning_artifact_parameters ServicecatalogCloudformationProduct#provisioning_artifact_parameters}
	ProvisioningArtifactParameters interface{} `field:"optional" json:"provisioningArtifactParameters" yaml:"provisioningArtifactParameters"`
	// This property is turned off by default.
	//
	// If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#replace_provisioning_artifacts ServicecatalogCloudformationProduct#replace_provisioning_artifacts}
	ReplaceProvisioningArtifacts interface{} `field:"optional" json:"replaceProvisioningArtifacts" yaml:"replaceProvisioningArtifacts"`
	// A top level ProductViewDetail response containing details about the product's connection.
	//
	// AWS Service Catalog returns this field for the CreateProduct, UpdateProduct, DescribeProductAsAdmin, and SearchProductAsAdmin APIs. This response contains the same fields as the ConnectionParameters request, with the addition of the LastSync response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#source_connection ServicecatalogCloudformationProduct#source_connection}
	SourceConnection *ServicecatalogCloudformationProductSourceConnection `field:"optional" json:"sourceConnection" yaml:"sourceConnection"`
	// The support information about the product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#support_description ServicecatalogCloudformationProduct#support_description}
	SupportDescription *string `field:"optional" json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#support_email ServicecatalogCloudformationProduct#support_email}
	SupportEmail *string `field:"optional" json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#support_url ServicecatalogCloudformationProduct#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_cloudformation_product#tags ServicecatalogCloudformationProduct#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

