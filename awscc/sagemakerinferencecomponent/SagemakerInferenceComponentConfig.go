// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerInferenceComponentConfig struct {
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
	// The name of the endpoint the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#endpoint_name SagemakerInferenceComponent#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The deployment config for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#deployment_config SagemakerInferenceComponent#deployment_config}
	DeploymentConfig *SagemakerInferenceComponentDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The Amazon Resource Name (ARN) of the endpoint the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#endpoint_arn SagemakerInferenceComponent#endpoint_arn}
	EndpointArn *string `field:"optional" json:"endpointArn" yaml:"endpointArn"`
	// The name of the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#inference_component_name SagemakerInferenceComponent#inference_component_name}
	InferenceComponentName *string `field:"optional" json:"inferenceComponentName" yaml:"inferenceComponentName"`
	// The runtime config for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#runtime_config SagemakerInferenceComponent#runtime_config}
	RuntimeConfig *SagemakerInferenceComponentRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
	// The specification for the inference component, for an endpoint with a single instance type.
	//
	// Specify exactly one of Specification or Specifications. InstanceType is not accepted here; use Specifications for per instance type configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#specification SagemakerInferenceComponent#specification}
	Specification *SagemakerInferenceComponentSpecification `field:"optional" json:"specification" yaml:"specification"`
	// A list of specification objects for the inference component, one per instance type.
	//
	// The service requires at least two entries; use the singular Specification for a single instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#specifications SagemakerInferenceComponent#specifications}
	Specifications interface{} `field:"optional" json:"specifications" yaml:"specifications"`
	// An array of tags to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#tags SagemakerInferenceComponent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the endpoint variant the inference component is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#variant_name SagemakerInferenceComponent#variant_name}
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
}

