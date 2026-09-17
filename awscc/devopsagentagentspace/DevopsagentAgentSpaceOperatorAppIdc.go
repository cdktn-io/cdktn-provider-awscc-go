// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace


type DevopsagentAgentSpaceOperatorAppIdc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_agent_space#idc_instance_arn DevopsagentAgentSpace#idc_instance_arn}.
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_agent_space#operator_app_role_arn DevopsagentAgentSpace#operator_app_role_arn}.
	OperatorAppRoleArn *string `field:"optional" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
}

