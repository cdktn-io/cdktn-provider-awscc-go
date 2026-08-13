// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentsecurityrequirementpack


type SecurityagentSecurityRequirementPackSecurityRequirements struct {
	// Description of the security requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_security_requirement_pack#description SecurityagentSecurityRequirementPack#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Security domain this requirement belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_security_requirement_pack#domain SecurityagentSecurityRequirementPack#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// How to evaluate compliance with this requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_security_requirement_pack#evaluation SecurityagentSecurityRequirementPack#evaluation}
	Evaluation *string `field:"optional" json:"evaluation" yaml:"evaluation"`
	// Name of the security requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_security_requirement_pack#name SecurityagentSecurityRequirementPack#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// How to remediate non-compliance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securityagent_security_requirement_pack#remediation SecurityagentSecurityRequirementPack#remediation}
	Remediation *string `field:"optional" json:"remediation" yaml:"remediation"`
}

