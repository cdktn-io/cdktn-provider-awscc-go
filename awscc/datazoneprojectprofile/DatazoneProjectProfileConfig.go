// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneprojectprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneProjectProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#name DatazoneProjectProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#allow_custom_project_resource_tags DatazoneProjectProfile#allow_custom_project_resource_tags}.
	AllowCustomProjectResourceTags interface{} `field:"optional" json:"allowCustomProjectResourceTags" yaml:"allowCustomProjectResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#description DatazoneProjectProfile#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#domain_identifier DatazoneProjectProfile#domain_identifier}.
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#domain_unit_identifier DatazoneProjectProfile#domain_unit_identifier}.
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#environment_configurations DatazoneProjectProfile#environment_configurations}.
	EnvironmentConfigurations interface{} `field:"optional" json:"environmentConfigurations" yaml:"environmentConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#project_resource_tags DatazoneProjectProfile#project_resource_tags}.
	ProjectResourceTags interface{} `field:"optional" json:"projectResourceTags" yaml:"projectResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#project_resource_tags_description DatazoneProjectProfile#project_resource_tags_description}.
	ProjectResourceTagsDescription *string `field:"optional" json:"projectResourceTagsDescription" yaml:"projectResourceTagsDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#status DatazoneProjectProfile#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datazone_project_profile#use_default_configurations DatazoneProjectProfile#use_default_configurations}.
	UseDefaultConfigurations interface{} `field:"optional" json:"useDefaultConfigurations" yaml:"useDefaultConfigurations"`
}

