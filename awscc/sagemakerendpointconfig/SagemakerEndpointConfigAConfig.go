// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigAConfig struct {
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
	// A list of ProductionVariant objects, one for each model that you want to host at this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#production_variants SagemakerEndpointConfigA#production_variants}
	ProductionVariants interface{} `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// Specifies configuration for how an endpoint performs asynchronous inference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#async_inference_config SagemakerEndpointConfigA#async_inference_config}
	AsyncInferenceConfig *SagemakerEndpointConfigAsyncInferenceConfig `field:"optional" json:"asyncInferenceConfig" yaml:"asyncInferenceConfig"`
	// Specifies how to capture endpoint data for model monitor.
	//
	// The data capture configuration applies to all production variants hosted at the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#data_capture_config SagemakerEndpointConfigA#data_capture_config}
	DataCaptureConfig *SagemakerEndpointConfigDataCaptureConfig `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// Sets whether all model containers deployed to the endpoint are isolated.
	//
	// If they are, no inbound or outbound network calls can be made to or from the model containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#enable_network_isolation SagemakerEndpointConfigA#enable_network_isolation}
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// The name of the endpoint configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#endpoint_config_name SagemakerEndpointConfigA#endpoint_config_name}
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform actions on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#execution_role_arn SagemakerEndpointConfigA#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A parameter to activate explainers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#explainer_config SagemakerEndpointConfigA#explainer_config}
	ExplainerConfig *SagemakerEndpointConfigExplainerConfig `field:"optional" json:"explainerConfig" yaml:"explainerConfig"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#kms_key_id SagemakerEndpointConfigA#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the metrics that the endpoint publishes to Amazon CloudWatch, the frequency of publication, and whether to enable enhanced or detailed observability metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#metrics_config SagemakerEndpointConfigA#metrics_config}
	MetricsConfig *SagemakerEndpointConfigMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// Array of ProductionVariant objects.
	//
	// There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#shadow_production_variants SagemakerEndpointConfigA#shadow_production_variants}
	ShadowProductionVariants interface{} `field:"optional" json:"shadowProductionVariants" yaml:"shadowProductionVariants"`
	// A list of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#tags SagemakerEndpointConfigA#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#vpc_config SagemakerEndpointConfigA#vpc_config}
	VpcConfig *SagemakerEndpointConfigVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

