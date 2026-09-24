// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigAiWorkloadConfigs struct {
	// The workload specification that defines benchmark parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_ai_workload_config#workload_spec SagemakerAiWorkloadConfig#workload_spec}
	WorkloadSpec *SagemakerAiWorkloadConfigAiWorkloadConfigsWorkloadSpec `field:"optional" json:"workloadSpec" yaml:"workloadSpec"`
}

