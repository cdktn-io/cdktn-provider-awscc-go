// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesworkspaceipgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesWorkspaceIpGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspaces_workspace_ip_group#group_name WorkspacesWorkspaceIpGroup#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The description of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspaces_workspace_ip_group#group_desc WorkspacesWorkspaceIpGroup#group_desc}
	GroupDesc *string `field:"optional" json:"groupDesc" yaml:"groupDesc"`
	// The tags for the IP access control group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspaces_workspace_ip_group#tags WorkspacesWorkspaceIpGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The rules for the IP access control group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/workspaces_workspace_ip_group#user_rules WorkspacesWorkspaceIpGroup#user_rules}
	UserRules interface{} `field:"optional" json:"userRules" yaml:"userRules"`
}

