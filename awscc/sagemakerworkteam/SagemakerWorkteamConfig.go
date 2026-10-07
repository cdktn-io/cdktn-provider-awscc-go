// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerWorkteamConfig struct {
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
	// A description of the work team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#description SagemakerWorkteam#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A list of MemberDefinition objects that contains objects that identify the workers that make up the work team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#member_definitions SagemakerWorkteam#member_definitions}
	MemberDefinitions interface{} `field:"required" json:"memberDefinitions" yaml:"memberDefinitions"`
	// Configures SNS notifications of available or expiring work items for work teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#notification_configuration SagemakerWorkteam#notification_configuration}
	NotificationConfiguration *SagemakerWorkteamNotificationConfiguration `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// An array of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#tags SagemakerWorkteam#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the Workforce.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#workforce_name SagemakerWorkteam#workforce_name}
	WorkforceName *string `field:"optional" json:"workforceName" yaml:"workforceName"`
	// The name of the work team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/sagemaker_workteam#workteam_name SagemakerWorkteam#workteam_name}
	WorkteamName *string `field:"optional" json:"workteamName" yaml:"workteamName"`
}

