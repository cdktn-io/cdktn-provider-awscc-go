// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftscheduledaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftScheduledActionConfig struct {
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
	// The name of the scheduled action. The name must be unique within an account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#scheduled_action_name RedshiftScheduledAction#scheduled_action_name}
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
	// If true, the schedule is enabled. If false, the scheduled action does not trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#enable RedshiftScheduledAction#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#end_time RedshiftScheduledAction#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The IAM role to assume to run the target action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#iam_role RedshiftScheduledAction#iam_role}
	IamRole *string `field:"optional" json:"iamRole" yaml:"iamRole"`
	// The schedule in `at( )` or `cron( )` format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#schedule RedshiftScheduledAction#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The description of the scheduled action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#scheduled_action_description RedshiftScheduledAction#scheduled_action_description}
	ScheduledActionDescription *string `field:"optional" json:"scheduledActionDescription" yaml:"scheduledActionDescription"`
	// The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#start_time RedshiftScheduledAction#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshift_scheduled_action#target_action RedshiftScheduledAction#target_action}
	TargetAction *RedshiftScheduledActionTargetAction `field:"optional" json:"targetAction" yaml:"targetAction"`
}

