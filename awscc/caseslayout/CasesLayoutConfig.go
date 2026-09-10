// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package caseslayout

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesLayoutConfig struct {
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
	// Defines the layout structure and field organization for the case interface.
	//
	// Specifies which fields appear in the top panel and More Info tab, and their display order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_layout#content CasesLayout#content}
	Content *CasesLayoutContent `field:"required" json:"content" yaml:"content"`
	// A descriptive name for the layout.
	//
	// Must be unique within the Cases domain and should clearly indicate the layout's purpose and field organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_layout#name CasesLayout#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier of the Cases domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_layout#domain_id CasesLayout#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The tags that you attach to this layout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cases_layout#tags CasesLayout#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

