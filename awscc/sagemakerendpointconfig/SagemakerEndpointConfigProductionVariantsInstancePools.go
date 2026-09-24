// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigProductionVariantsInstancePools struct {
	// The ML compute instance type for the instance pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#instance_type SagemakerEndpointConfigA#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The name of a SageMaker model to use for this instance pool instead of the model specified for the production variant.
	//
	// Use this to deploy a different model optimized for the instance type in this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#model_name_override SagemakerEndpointConfigA#model_name_override}
	ModelNameOverride *string `field:"optional" json:"modelNameOverride" yaml:"modelNameOverride"`
	// The priority for the instance pool.
	//
	// SageMaker attempts to provision instances in order of priority, starting with the lowest value. If instances for a higher-priority pool are unavailable, SageMaker attempts to provision from the next pool. Valid values: 1 to 5, where 1 is the highest priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#priority SagemakerEndpointConfigA#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

