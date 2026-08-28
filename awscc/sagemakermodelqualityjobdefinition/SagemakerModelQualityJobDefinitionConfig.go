// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerModelQualityJobDefinitionConfig struct {
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
	// Identifies the resources to deploy for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#job_resources SagemakerModelQualityJobDefinition#job_resources}
	JobResources *SagemakerModelQualityJobDefinitionJobResources `field:"required" json:"jobResources" yaml:"jobResources"`
	// Container image configuration object for the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#model_quality_app_specification SagemakerModelQualityJobDefinition#model_quality_app_specification}
	ModelQualityAppSpecification *SagemakerModelQualityJobDefinitionModelQualityAppSpecification `field:"required" json:"modelQualityAppSpecification" yaml:"modelQualityAppSpecification"`
	// The inputs for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#model_quality_job_input SagemakerModelQualityJobDefinition#model_quality_job_input}
	ModelQualityJobInput *SagemakerModelQualityJobDefinitionModelQualityJobInput `field:"required" json:"modelQualityJobInput" yaml:"modelQualityJobInput"`
	// The output configuration for monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#model_quality_job_output_config SagemakerModelQualityJobDefinition#model_quality_job_output_config}
	ModelQualityJobOutputConfig *SagemakerModelQualityJobDefinitionModelQualityJobOutputConfig `field:"required" json:"modelQualityJobOutputConfig" yaml:"modelQualityJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#role_arn SagemakerModelQualityJobDefinition#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the endpoint used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#endpoint_name SagemakerModelQualityJobDefinition#endpoint_name}
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the job definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#job_definition_name SagemakerModelQualityJobDefinition#job_definition_name}
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#model_quality_baseline_config SagemakerModelQualityJobDefinition#model_quality_baseline_config}
	ModelQualityBaselineConfig *SagemakerModelQualityJobDefinitionModelQualityBaselineConfig `field:"optional" json:"modelQualityBaselineConfig" yaml:"modelQualityBaselineConfig"`
	// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#network_config SagemakerModelQualityJobDefinition#network_config}
	NetworkConfig *SagemakerModelQualityJobDefinitionNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#stopping_condition SagemakerModelQualityJobDefinition#stopping_condition}
	StoppingCondition *SagemakerModelQualityJobDefinitionStoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_quality_job_definition#tags SagemakerModelQualityJobDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

