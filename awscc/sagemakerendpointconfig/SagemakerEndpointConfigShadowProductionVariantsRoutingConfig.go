// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigShadowProductionVariantsRoutingConfig struct {
	// The configuration for prefix-aware routing. Specify this property only when you set RoutingStrategy to PREFIX_AWARE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#prefix_aware_routing_config SagemakerEndpointConfigA#prefix_aware_routing_config}
	PrefixAwareRoutingConfig *SagemakerEndpointConfigShadowProductionVariantsRoutingConfigPrefixAwareRoutingConfig `field:"optional" json:"prefixAwareRoutingConfig" yaml:"prefixAwareRoutingConfig"`
	// Sets how the endpoint routes incoming traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#routing_strategy SagemakerEndpointConfigA#routing_strategy}
	RoutingStrategy *string `field:"optional" json:"routingStrategy" yaml:"routingStrategy"`
}

