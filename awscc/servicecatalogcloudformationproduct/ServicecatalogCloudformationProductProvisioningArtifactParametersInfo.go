// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationproduct


type ServicecatalogCloudformationProductProvisioningArtifactParametersInfo struct {
	// The physical id of the resource that contains the template. Currently only supports AWS CloudFormation stack arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#import_from_physical_id ServicecatalogCloudformationProduct#import_from_physical_id}
	ImportFromPhysicalId *string `field:"optional" json:"importFromPhysicalId" yaml:"importFromPhysicalId"`
	// The URL of the AWS CloudFormation template in Amazon S3 in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_cloudformation_product#load_template_from_url ServicecatalogCloudformationProduct#load_template_from_url}
	LoadTemplateFromUrl *string `field:"optional" json:"loadTemplateFromUrl" yaml:"loadTemplateFromUrl"`
}

