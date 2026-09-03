// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesCaseConfig struct {
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
	// The full customer profile ARN for the case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_case#customer_id CasesCase#customer_id}
	CustomerId *string `field:"required" json:"customerId" yaml:"customerId"`
	// The unique identifier of the Cases domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_case#domain_id CasesCase#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// A unique identifier of a template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_case#template_id CasesCase#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// The title of the case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_case#title CasesCase#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of tags for the case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cases_case#tags CasesCase#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

