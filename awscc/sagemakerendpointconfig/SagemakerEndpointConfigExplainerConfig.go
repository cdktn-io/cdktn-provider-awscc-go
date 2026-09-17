// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfig struct {
	// A member of ExplainerConfig that contains configuration parameters for the SageMaker Clarify explainer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#clarify_explainer_config SagemakerEndpointConfigA#clarify_explainer_config}
	ClarifyExplainerConfig *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfig `field:"optional" json:"clarifyExplainerConfig" yaml:"clarifyExplainerConfig"`
}

