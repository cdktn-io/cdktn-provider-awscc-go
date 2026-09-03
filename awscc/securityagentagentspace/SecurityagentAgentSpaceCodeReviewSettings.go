// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceCodeReviewSettings struct {
	// Whether Controls are utilized for code review analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#controls_scanning SecurityagentAgentSpace#controls_scanning}
	ControlsScanning interface{} `field:"optional" json:"controlsScanning" yaml:"controlsScanning"`
	// Whether general purpose analysis is performed for code review.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#general_purpose_scanning SecurityagentAgentSpace#general_purpose_scanning}
	GeneralPurposeScanning interface{} `field:"optional" json:"generalPurposeScanning" yaml:"generalPurposeScanning"`
}

