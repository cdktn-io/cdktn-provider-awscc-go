// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesworkspaceipgroup


type WorkspacesWorkspaceIpGroupUserRules struct {
	// The IP address range, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/workspaces_workspace_ip_group#ip_rule WorkspacesWorkspaceIpGroup#ip_rule}
	IpRule *string `field:"optional" json:"ipRule" yaml:"ipRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/workspaces_workspace_ip_group#rule_desc WorkspacesWorkspaceIpGroup#rule_desc}
	RuleDesc *string `field:"optional" json:"ruleDesc" yaml:"ruleDesc"`
}

