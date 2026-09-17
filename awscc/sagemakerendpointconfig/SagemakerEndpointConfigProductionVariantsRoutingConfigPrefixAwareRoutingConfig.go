// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigProductionVariantsRoutingConfigPrefixAwareRoutingConfig struct {
	// The maximum number of in-flight requests on the target instance before the endpoint routes to another instance.
	//
	// Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1 to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#concurrency_threshold SagemakerEndpointConfigA#concurrency_threshold}
	ConcurrencyThreshold *float64 `field:"optional" json:"concurrencyThreshold" yaml:"concurrencyThreshold"`
	// The maximum length of the prefix used for routing decisions.
	//
	// Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1024 to 65536.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#prefix_length SagemakerEndpointConfigA#prefix_length}
	PrefixLength *float64 `field:"optional" json:"prefixLength" yaml:"prefixLength"`
}

