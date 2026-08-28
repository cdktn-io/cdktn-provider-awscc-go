// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubapp


type ResiliencehubAppResourceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#mapping_type ResiliencehubApp#mapping_type}.
	MappingType *string `field:"required" json:"mappingType" yaml:"mappingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#physical_resource_id ResiliencehubApp#physical_resource_id}.
	PhysicalResourceId *ResiliencehubAppResourceMappingsPhysicalResourceId `field:"required" json:"physicalResourceId" yaml:"physicalResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#eks_source_name ResiliencehubApp#eks_source_name}.
	EksSourceName *string `field:"optional" json:"eksSourceName" yaml:"eksSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#logical_stack_name ResiliencehubApp#logical_stack_name}.
	LogicalStackName *string `field:"optional" json:"logicalStackName" yaml:"logicalStackName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#resource_name ResiliencehubApp#resource_name}.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehub_app#terraform_source_name ResiliencehubApp#terraform_source_name}.
	TerraformSourceName *string `field:"optional" json:"terraformSourceName" yaml:"terraformSourceName"`
}

