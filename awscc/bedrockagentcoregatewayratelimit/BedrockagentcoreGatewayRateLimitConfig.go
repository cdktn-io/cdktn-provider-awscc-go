// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewayratelimit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreGatewayRateLimitConfig struct {
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
	// Ordered list of dimension names defining the scope of this limit.
	//
	// Unique per gateway — no two limits can share the same dimensionKeys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rate_limit#dimension_keys BedrockagentcoreGatewayRateLimit#dimension_keys}
	DimensionKeys *[]*string `field:"required" json:"dimensionKeys" yaml:"dimensionKeys"`
	// Rule entries mapping dimension values to rate configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rate_limit#entries BedrockagentcoreGatewayRateLimit#entries}
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
	// Optional human-readable description for this limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rate_limit#description BedrockagentcoreGatewayRateLimit#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rate_limit#gateway_identifier BedrockagentcoreGatewayRateLimit#gateway_identifier}.
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// Limit identifier. Optional on Create (system-generates if not provided by customer). Always present in responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_gateway_rate_limit#rate_limit_id BedrockagentcoreGatewayRateLimit#rate_limit_id}
	RateLimitId *string `field:"optional" json:"rateLimitId" yaml:"rateLimitId"`
}

