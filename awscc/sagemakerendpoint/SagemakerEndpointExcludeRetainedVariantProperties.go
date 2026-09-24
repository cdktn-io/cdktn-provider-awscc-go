// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointExcludeRetainedVariantProperties struct {
	// The type of variant property (e.g., 'DesiredInstanceCount', 'DesiredWeight', 'DataCaptureConfig').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#variant_property_type SagemakerEndpoint#variant_property_type}
	VariantPropertyType *string `field:"optional" json:"variantPropertyType" yaml:"variantPropertyType"`
}

