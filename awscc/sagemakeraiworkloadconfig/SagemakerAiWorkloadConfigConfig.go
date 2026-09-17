// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerAiWorkloadConfigConfig struct {
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
	// The name of the AI workload configuration.
	//
	// The name must be unique within your AWS account in the current AWS Region. Only lowercase letters and digits are accepted: DeleteAIWorkloadConfig lowercases the name before looking it up, so a name containing an uppercase letter produces a configuration that can be created and read but never deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#ai_workload_config_name SagemakerAiWorkloadConfig#ai_workload_config_name}
	AiWorkloadConfigName *string `field:"required" json:"aiWorkloadConfigName" yaml:"aiWorkloadConfigName"`
	// The benchmark tool configuration and workload specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#ai_workload_configs SagemakerAiWorkloadConfig#ai_workload_configs}
	AiWorkloadConfigs *SagemakerAiWorkloadConfigAiWorkloadConfigs `field:"optional" json:"aiWorkloadConfigs" yaml:"aiWorkloadConfigs"`
	// The dataset configuration for the workload. Specify input data channels with their data sources for benchmark workloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#dataset_config SagemakerAiWorkloadConfig#dataset_config}
	DatasetConfig *SagemakerAiWorkloadConfigDatasetConfig `field:"optional" json:"datasetConfig" yaml:"datasetConfig"`
	// The metadata that you apply to the AI workload configuration to help you categorize and organize it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#tags SagemakerAiWorkloadConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

