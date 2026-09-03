// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecataloglaunchroleconstraint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogLaunchRoleConstraintConfig struct {
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
	// The ID of the portfolio to which this launch role constraint applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#portfolio_id ServicecatalogLaunchRoleConstraint#portfolio_id}
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The ID of the product to which this launch role constraint applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#product_id ServicecatalogLaunchRoleConstraint#product_id}
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The language code for the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#accept_language ServicecatalogLaunchRoleConstraint#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the launch role constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#description ServicecatalogLaunchRoleConstraint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The local IAM role name to use in the launch constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#local_role_name ServicecatalogLaunchRoleConstraint#local_role_name}
	LocalRoleName *string `field:"optional" json:"localRoleName" yaml:"localRoleName"`
	// The ARN of the IAM role used for the launch constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/servicecatalog_launch_role_constraint#role_arn ServicecatalogLaunchRoleConstraint#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

