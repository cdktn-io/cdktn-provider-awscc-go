// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectWorkspaceConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#instance_arn ConnectWorkspace#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#name ConnectWorkspace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource ARNs associated with the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#associations ConnectWorkspace#associations}
	Associations *[]*string `field:"optional" json:"associations" yaml:"associations"`
	// The description of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#description ConnectWorkspace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The media items for the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#media ConnectWorkspace#media}
	Media interface{} `field:"optional" json:"media" yaml:"media"`
	// The pages associated with the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#pages ConnectWorkspace#pages}
	Pages interface{} `field:"optional" json:"pages" yaml:"pages"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#tags ConnectWorkspace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The theme configuration for the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#theme ConnectWorkspace#theme}
	Theme *ConnectWorkspaceTheme `field:"optional" json:"theme" yaml:"theme"`
	// The title of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#title ConnectWorkspace#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The visibility of the workspace. Will always be set to ASSIGNED oninitial creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_workspace#visibility ConnectWorkspace#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

