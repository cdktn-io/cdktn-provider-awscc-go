// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceNetworkAccessControl struct {
	// The list of prefix list IDs.
	//
	// A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration then no IP addresses will be allowed to access the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/grafana_workspace#prefix_list_ids GrafanaWorkspace#prefix_list_ids}
	PrefixListIds *[]*string `field:"optional" json:"prefixListIds" yaml:"prefixListIds"`
	// The list of Amazon VPC endpoint IDs for the workspace.
	//
	// If a NetworkAccessConfiguration is specified then only VPC endpoints specified here will be allowed to access the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/grafana_workspace#vpce_ids GrafanaWorkspace#vpce_ids}
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

