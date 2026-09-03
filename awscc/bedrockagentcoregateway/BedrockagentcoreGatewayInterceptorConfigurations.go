// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayInterceptorConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway#input_configuration BedrockagentcoreGateway#input_configuration}.
	InputConfiguration *BedrockagentcoreGatewayInterceptorConfigurationsInputConfiguration `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway#interception_points BedrockagentcoreGateway#interception_points}.
	InterceptionPoints *[]*string `field:"optional" json:"interceptionPoints" yaml:"interceptionPoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_gateway#interceptor BedrockagentcoreGateway#interceptor}.
	Interceptor *BedrockagentcoreGatewayInterceptorConfigurationsInterceptor `field:"optional" json:"interceptor" yaml:"interceptor"`
}

