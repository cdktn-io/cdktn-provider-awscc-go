// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeralgorithm


type SagemakerAlgorithmTrainingSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#supported_training_instance_types SagemakerAlgorithm#supported_training_instance_types}.
	SupportedTrainingInstanceTypes *[]*string `field:"required" json:"supportedTrainingInstanceTypes" yaml:"supportedTrainingInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#training_channels SagemakerAlgorithm#training_channels}.
	TrainingChannels interface{} `field:"required" json:"trainingChannels" yaml:"trainingChannels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#training_image SagemakerAlgorithm#training_image}.
	TrainingImage *string `field:"required" json:"trainingImage" yaml:"trainingImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#metric_definitions SagemakerAlgorithm#metric_definitions}.
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#supported_hyper_parameters SagemakerAlgorithm#supported_hyper_parameters}.
	SupportedHyperParameters interface{} `field:"optional" json:"supportedHyperParameters" yaml:"supportedHyperParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#supported_tuning_job_objective_metrics SagemakerAlgorithm#supported_tuning_job_objective_metrics}.
	SupportedTuningJobObjectiveMetrics interface{} `field:"optional" json:"supportedTuningJobObjectiveMetrics" yaml:"supportedTuningJobObjectiveMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#supports_distributed_training SagemakerAlgorithm#supports_distributed_training}.
	SupportsDistributedTraining interface{} `field:"optional" json:"supportsDistributedTraining" yaml:"supportsDistributedTraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_algorithm#training_image_digest SagemakerAlgorithm#training_image_digest}.
	TrainingImageDigest *string `field:"optional" json:"trainingImageDigest" yaml:"trainingImageDigest"`
}

