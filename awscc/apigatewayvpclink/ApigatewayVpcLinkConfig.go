// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayvpclink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayVpcLinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_vpc_link#name ApigatewayVpcLink#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_vpc_link#target_arns ApigatewayVpcLink#target_arns}.
	TargetArns *[]*string `field:"required" json:"targetArns" yaml:"targetArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_vpc_link#description ApigatewayVpcLink#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_vpc_link#tags ApigatewayVpcLink#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

