// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayInterceptorConfigurationsInputConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_gateway#pass_request_headers BedrockagentcoreGateway#pass_request_headers}.
	PassRequestHeaders interface{} `field:"optional" json:"passRequestHeaders" yaml:"passRequestHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrockagentcore_gateway#payload_filter BedrockagentcoreGateway#payload_filter}.
	PayloadFilter *BedrockagentcoreGatewayInterceptorConfigurationsInputConfigurationPayloadFilter `field:"optional" json:"payloadFilter" yaml:"payloadFilter"`
}

