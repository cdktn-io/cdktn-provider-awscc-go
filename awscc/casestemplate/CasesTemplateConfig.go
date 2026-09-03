// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casestemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesTemplateConfig struct {
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
	// A name for the template. It must be unique per domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#name CasesTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description explaining the purpose and use case for this template.
	//
	// Should indicate what types of cases this template is designed for and any specific workflow it supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#description CasesTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the Cases domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#domain_id CasesTemplate#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Specifies the default layout to use when displaying cases created from this template.
	//
	// The layout determines which fields are visible and their arrangement in the agent interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#layout_configuration CasesTemplate#layout_configuration}
	LayoutConfiguration *CasesTemplateLayoutConfiguration `field:"optional" json:"layoutConfiguration" yaml:"layoutConfiguration"`
	// A list of fields that must contain a value for a case to be successfully created with this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#required_fields CasesTemplate#required_fields}
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
	// A list of case rules (also known as case field conditions) on a template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#rules CasesTemplate#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The current status of the template.
	//
	// Active templates can be used to create new cases, while Inactive templates are disabled but preserved for existing cases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#status CasesTemplate#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags that you attach to this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_template#tags CasesTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

