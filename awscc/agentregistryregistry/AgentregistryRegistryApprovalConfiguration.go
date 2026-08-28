// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistry


type AgentregistryRegistryApprovalConfiguration struct {
	// The rules that determine which registry records are automatically approved on submission.
	//
	// When omitted or empty, submitted records require manual review.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry#auto_approval_rules AgentregistryRegistry#auto_approval_rules}
	AutoApprovalRules *[]*string `field:"optional" json:"autoApprovalRules" yaml:"autoApprovalRules"`
}

