// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceDevOpsAgentConfiguration struct {
	// The ARN of the DevOps Agent Agent Space to link to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/resiliencehubv2_service#agent_space_arn Resiliencehubv2Service#agent_space_arn}
	AgentSpaceArn *string `field:"optional" json:"agentSpaceArn" yaml:"agentSpaceArn"`
}

