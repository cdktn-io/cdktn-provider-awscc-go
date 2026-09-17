// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigProductionVariantsManagedInstanceScaling struct {
	// The maximum number of instances that the endpoint can provision when it scales up to accommodate an increase in traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#max_instance_count SagemakerEndpointConfigA#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// The minimum number of instances that the endpoint must retain when it scales down to accommodate a decrease in traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#min_instance_count SagemakerEndpointConfigA#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// Configures the scale-in behavior for managed instance scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#scale_in_policy SagemakerEndpointConfigA#scale_in_policy}
	ScaleInPolicy *SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicy `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// Indicates whether managed instance scaling is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#status SagemakerEndpointConfigA#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

