// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostssite

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OutpostsSiteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#name OutpostsSite#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#description OutpostsSite#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#notes OutpostsSite#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#operating_address OutpostsSite#operating_address}.
	OperatingAddress *OutpostsSiteOperatingAddress `field:"optional" json:"operatingAddress" yaml:"operatingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#rack_physical_properties OutpostsSite#rack_physical_properties}.
	RackPhysicalProperties *OutpostsSiteRackPhysicalProperties `field:"optional" json:"rackPhysicalProperties" yaml:"rackPhysicalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#shipping_address OutpostsSite#shipping_address}.
	ShippingAddress *OutpostsSiteShippingAddress `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/outposts_site#tags OutpostsSite#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

