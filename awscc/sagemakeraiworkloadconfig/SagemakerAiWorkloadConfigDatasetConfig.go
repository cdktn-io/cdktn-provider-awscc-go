// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigDatasetConfig struct {
	// An array of input data channel configurations for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#input_data_config SagemakerAiWorkloadConfig#input_data_config}
	InputDataConfig interface{} `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
}

