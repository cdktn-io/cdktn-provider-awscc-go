// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationgeneratedtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationGeneratedTemplateConfig struct {
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
	// The name assigned to the generated template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_generated_template#generated_template_name CloudformationGeneratedTemplate#generated_template_name}
	GeneratedTemplateName *string `field:"required" json:"generatedTemplateName" yaml:"generatedTemplateName"`
	// The configuration details of the generated template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_generated_template#template_configuration CloudformationGeneratedTemplate#template_configuration}
	TemplateConfiguration *CloudformationGeneratedTemplateTemplateConfiguration `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

