// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscheduledaction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingScheduledActionConfig struct {
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
	// The name of the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#auto_scaling_group_name AutoscalingScheduledAction#auto_scaling_group_name}
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#desired_capacity AutoscalingScheduledAction#desired_capacity}
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#end_time AutoscalingScheduledAction#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The minimum size of the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#max_size AutoscalingScheduledAction#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum size of the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#min_size AutoscalingScheduledAction#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// The recurring schedule for the action, in Unix cron syntax format.
	//
	// When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#recurrence AutoscalingScheduledAction#recurrence}
	Recurrence *string `field:"optional" json:"recurrence" yaml:"recurrence"`
	// The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#start_time AutoscalingScheduledAction#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The time zone for the cron expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scheduled_action#time_zone AutoscalingScheduledAction#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

