// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigAiWorkloadConfigsWorkloadSpec struct {
	// An inline YAML or JSON string that defines benchmark parameters.
	//
	// The service validates the document against its own benchmark schema: it must declare a benchmark object whose type member matches the pattern ^(aiperf)$.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_ai_workload_config#inline SagemakerAiWorkloadConfig#inline}
	Inline *string `field:"optional" json:"inline" yaml:"inline"`
}

