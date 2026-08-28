// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecataloglaunchtemplateconstraint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogLaunchTemplateConstraintConfig struct {
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
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_launch_template_constraint#portfolio_id ServicecatalogLaunchTemplateConstraint#portfolio_id}
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_launch_template_constraint#product_id ServicecatalogLaunchTemplateConstraint#product_id}
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// A json encoded string of the template constraint rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_launch_template_constraint#rules ServicecatalogLaunchTemplateConstraint#rules}
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_launch_template_constraint#accept_language ServicecatalogLaunchTemplateConstraint#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_launch_template_constraint#description ServicecatalogLaunchTemplateConstraint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

