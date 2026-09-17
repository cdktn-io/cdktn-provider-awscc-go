// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesinstances_workspace_instance#enable_resource_name_dns_aaaa_record WorkspacesinstancesWorkspaceInstance#enable_resource_name_dns_aaaa_record}.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesinstances_workspace_instance#enable_resource_name_dns_a_record WorkspacesinstancesWorkspaceInstance#enable_resource_name_dns_a_record}.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/workspacesinstances_workspace_instance#hostname_type WorkspacesinstancesWorkspaceInstance#hostname_type}.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

