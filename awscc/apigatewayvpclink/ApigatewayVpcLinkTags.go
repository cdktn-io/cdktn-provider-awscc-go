// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayvpclink


type ApigatewayVpcLinkTags struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_vpc_link#key ApigatewayVpcLink#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the specified tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_vpc_link#value ApigatewayVpcLink#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

