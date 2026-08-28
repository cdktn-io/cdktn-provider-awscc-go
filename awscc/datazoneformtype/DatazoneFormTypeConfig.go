// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneformtype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneFormTypeConfig struct {
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
	// The ID of the Amazon DataZone domain in which this metadata form type is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#domain_identifier DatazoneFormType#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The model of this Amazon DataZone metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#model DatazoneFormType#model}
	Model *DatazoneFormTypeModel `field:"required" json:"model" yaml:"model"`
	// The name of this Amazon DataZone metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#name DatazoneFormType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Amazon DataZone project that owns this metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#owning_project_identifier DatazoneFormType#owning_project_identifier}
	OwningProjectIdentifier *string `field:"required" json:"owningProjectIdentifier" yaml:"owningProjectIdentifier"`
	// The description of this Amazon DataZone metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#description DatazoneFormType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The status of this Amazon DataZone metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_form_type#status DatazoneFormType#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

