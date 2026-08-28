// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingbillingview

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BillingBillingViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#name BillingBillingView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of strings that define the billing view's source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#source_views BillingBillingView#source_views}
	SourceViews *[]*string `field:"required" json:"sourceViews" yaml:"sourceViews"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#data_filter_expression BillingBillingView#data_filter_expression}.
	DataFilterExpression *BillingBillingViewDataFilterExpression `field:"optional" json:"dataFilterExpression" yaml:"dataFilterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#description BillingBillingView#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs associated to the billing view being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/billing_billing_view#tags BillingBillingView#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

