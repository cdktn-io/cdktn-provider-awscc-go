// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesfield

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesFieldConfig struct {
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
	// The display name of the field as it appears to agents in the case interface.
	//
	// Should be descriptive and user-friendly (e.g., 'Customer Priority Level', 'Issue Category').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#name CasesField#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the field, which determines validation rules, input constraints, and display format.
	//
	// Each type has specific constraints: Text (string input), Number (numeric values), Boolean (true/false), DateTime (date/time picker), SingleSelect (dropdown options), Url (URL validation), User (Amazon Connect user selection).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#type CasesField#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Field-type specific attributes that control rendering and validation behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#attributes CasesField#attributes}
	Attributes *CasesFieldAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// A description explaining the purpose and usage of this field in cases.
	//
	// Helps agents and administrators understand what information should be captured in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#description CasesField#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the Cases domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#domain_id CasesField#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The tags that you attach to this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cases_field#tags CasesField#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

