// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#actions BedrockagentcoreGatewayRule#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#priority BedrockagentcoreGatewayRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#conditions BedrockagentcoreGatewayRule#conditions}.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#description BedrockagentcoreGatewayRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rule#gateway_identifier BedrockagentcoreGatewayRule#gateway_identifier}.
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
}

