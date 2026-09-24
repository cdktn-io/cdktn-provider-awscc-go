// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationproduct


type ServicecatalogCloudformationProductSourceConnection struct {
	// The connection details based on the connection Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_cloudformation_product#connection_parameters ServicecatalogCloudformationProduct#connection_parameters}
	ConnectionParameters *ServicecatalogCloudformationProductSourceConnectionConnectionParameters `field:"optional" json:"connectionParameters" yaml:"connectionParameters"`
	// The only supported SourceConnection type is Codestar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_cloudformation_product#type ServicecatalogCloudformationProduct#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

