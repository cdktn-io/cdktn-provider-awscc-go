// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationprovisionedproduct

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogCloudformationProvisionedProductConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#accept_language ServicecatalogCloudformationProvisionedProduct#accept_language}.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#notification_arns ServicecatalogCloudformationProvisionedProduct#notification_arns}.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#path_id ServicecatalogCloudformationProvisionedProduct#path_id}.
	PathId *string `field:"optional" json:"pathId" yaml:"pathId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#path_name ServicecatalogCloudformationProvisionedProduct#path_name}.
	PathName *string `field:"optional" json:"pathName" yaml:"pathName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#product_id ServicecatalogCloudformationProvisionedProduct#product_id}.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#product_name ServicecatalogCloudformationProvisionedProduct#product_name}.
	ProductName *string `field:"optional" json:"productName" yaml:"productName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#provisioned_product_name ServicecatalogCloudformationProvisionedProduct#provisioned_product_name}.
	ProvisionedProductName *string `field:"optional" json:"provisionedProductName" yaml:"provisionedProductName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#provisioning_artifact_id ServicecatalogCloudformationProvisionedProduct#provisioning_artifact_id}.
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#provisioning_artifact_name ServicecatalogCloudformationProvisionedProduct#provisioning_artifact_name}.
	ProvisioningArtifactName *string `field:"optional" json:"provisioningArtifactName" yaml:"provisioningArtifactName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#provisioning_parameters ServicecatalogCloudformationProvisionedProduct#provisioning_parameters}.
	ProvisioningParameters interface{} `field:"optional" json:"provisioningParameters" yaml:"provisioningParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#provisioning_preferences ServicecatalogCloudformationProvisionedProduct#provisioning_preferences}.
	ProvisioningPreferences *ServicecatalogCloudformationProvisionedProductProvisioningPreferences `field:"optional" json:"provisioningPreferences" yaml:"provisioningPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_cloudformation_provisioned_product#tags ServicecatalogCloudformationProvisionedProduct#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

