// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogstacksetconstraint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogStackSetConstraintConfig struct {
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
	// One or more AWS accounts that will have access to the provisioned product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#account_list ServicecatalogStackSetConstraint#account_list}
	AccountList *[]*string `field:"required" json:"accountList" yaml:"accountList"`
	// AdminRole ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#admin_role ServicecatalogStackSetConstraint#admin_role}
	AdminRole *string `field:"required" json:"adminRole" yaml:"adminRole"`
	// The description of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#description ServicecatalogStackSetConstraint#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// ExecutionRole name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#execution_role ServicecatalogStackSetConstraint#execution_role}
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#portfolio_id ServicecatalogStackSetConstraint#portfolio_id}
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#product_id ServicecatalogStackSetConstraint#product_id}
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#region_list ServicecatalogStackSetConstraint#region_list}
	RegionList *[]*string `field:"required" json:"regionList" yaml:"regionList"`
	// Permission to create, update, and delete stack instances. Choose from ALLOWED and NOT_ALLOWED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#stack_instance_control ServicecatalogStackSetConstraint#stack_instance_control}
	StackInstanceControl *string `field:"required" json:"stackInstanceControl" yaml:"stackInstanceControl"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/servicecatalog_stack_set_constraint#accept_language ServicecatalogStackSetConstraint#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

