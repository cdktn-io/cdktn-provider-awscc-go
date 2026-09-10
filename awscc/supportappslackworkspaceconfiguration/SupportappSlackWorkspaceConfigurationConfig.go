// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportappslackworkspaceconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SupportappSlackWorkspaceConfigurationConfig struct {
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
	// The team ID in Slack, which uniquely identifies a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/supportapp_slack_workspace_configuration#team_id SupportappSlackWorkspaceConfiguration#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/supportapp_slack_workspace_configuration#version_id SupportappSlackWorkspaceConfiguration#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

