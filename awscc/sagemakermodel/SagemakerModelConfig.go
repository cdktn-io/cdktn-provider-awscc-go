// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelConfig struct {
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
	// Specifies the containers in the inference pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#containers SagemakerModel#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#enable_network_isolation SagemakerModel#enable_network_isolation}
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// The Amazon Resource Name (ARN) of the IAM role that you specified for the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#execution_role_arn SagemakerModel#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies details about how containers in a multi-container endpoint are run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#inference_execution_config SagemakerModel#inference_execution_config}
	InferenceExecutionConfig *SagemakerModelInferenceExecutionConfig `field:"optional" json:"inferenceExecutionConfig" yaml:"inferenceExecutionConfig"`
	// The name of the new model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#model_name SagemakerModel#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Describes the container, as part of model definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#primary_container SagemakerModel#primary_container}
	PrimaryContainer *SagemakerModelPrimaryContainer `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// An array of key-value pairs.
	//
	// You can use tags to categorize your AWS resources in different ways, for example, by purpose, owner, or environment. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#tags SagemakerModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_model#vpc_config SagemakerModel#vpc_config}
	VpcConfig *SagemakerModelVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

