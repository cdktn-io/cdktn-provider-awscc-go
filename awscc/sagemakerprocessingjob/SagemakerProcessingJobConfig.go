// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerProcessingJobConfig struct {
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
	// Configures the processing job to run a specified Docker container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#app_specification SagemakerProcessingJob#app_specification}
	AppSpecification *SagemakerProcessingJobAppSpecification `field:"required" json:"appSpecification" yaml:"appSpecification"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.
	//
	// In distributed training, you specify more than one instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#processing_resources SagemakerProcessingJob#processing_resources}
	ProcessingResources *SagemakerProcessingJobProcessingResources `field:"required" json:"processingResources" yaml:"processingResources"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#role_arn SagemakerProcessingJob#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#environment SagemakerProcessingJob#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Associates a SageMaker job as a trial component with an experiment and trial.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#experiment_config SagemakerProcessingJob#experiment_config}
	ExperimentConfig *SagemakerProcessingJobExperimentConfig `field:"optional" json:"experimentConfig" yaml:"experimentConfig"`
	// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#network_config SagemakerProcessingJob#network_config}
	NetworkConfig *SagemakerProcessingJobNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// An array of inputs configuring the data to download into the processing container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#processing_inputs SagemakerProcessingJob#processing_inputs}
	ProcessingInputs interface{} `field:"optional" json:"processingInputs" yaml:"processingInputs"`
	// The name of the processing job. The name must be unique within an AWS Region in the AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#processing_job_name SagemakerProcessingJob#processing_job_name}
	ProcessingJobName *string `field:"optional" json:"processingJobName" yaml:"processingJobName"`
	// Configuration for uploading output from the processing container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#processing_output_config SagemakerProcessingJob#processing_output_config}
	ProcessingOutputConfig *SagemakerProcessingJobProcessingOutputConfig `field:"optional" json:"processingOutputConfig" yaml:"processingOutputConfig"`
	// Configures conditions under which the processing job should be stopped, such as how long the processing job has been running.
	//
	// After the condition is met, the processing job is stopped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#stopping_condition SagemakerProcessingJob#stopping_condition}
	StoppingCondition *SagemakerProcessingJobStoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// (Optional) An array of key-value pairs.
	//
	// For more information, see Using Cost Allocation Tags(https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the AWS Billing and Cost Management User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#tags SagemakerProcessingJob#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

