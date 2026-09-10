// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottwinmakerworkspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IottwinmakerWorkspaceConfig struct {
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
	// The ARN of the execution role associated with the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iottwinmaker_workspace#role IottwinmakerWorkspace#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iottwinmaker_workspace#s3_location IottwinmakerWorkspace#s3_location}
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
	// The ID of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iottwinmaker_workspace#workspace_id IottwinmakerWorkspace#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iottwinmaker_workspace#description IottwinmakerWorkspace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of key-value pairs to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iottwinmaker_workspace#tags IottwinmakerWorkspace#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

