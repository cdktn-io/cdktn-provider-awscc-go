// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentagentspace


type DevopsagentAgentSpaceOperatorApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_agent_space#iam DevopsagentAgentSpace#iam}.
	Iam *DevopsagentAgentSpaceOperatorAppIam `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_agent_space#idc DevopsagentAgentSpace#idc}.
	Idc *DevopsagentAgentSpaceOperatorAppIdc `field:"optional" json:"idc" yaml:"idc"`
}

