// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalabletarget


type ApplicationautoscalingScalableTargetScheduledActions struct {
	// The date and time that the action is scheduled to end, in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#end_time ApplicationautoscalingScalableTarget#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The new minimum and maximum capacity.
	//
	// You can set both values or just one. At the scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling scales in to the maximum capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#scalable_target_action ApplicationautoscalingScalableTarget#scalable_target_action}
	ScalableTargetAction *ApplicationautoscalingScalableTargetScheduledActionsScalableTargetAction `field:"optional" json:"scalableTargetAction" yaml:"scalableTargetAction"`
	// The schedule for this action.
	//
	// The following formats are supported:
	//   +  At expressions - "``at(yyyy-mm-ddThh:mm:ss)``"
	//   +  Rate expressions - "``rate(valueunit)``"
	//   +  Cron expressions - "``cron(fields)``"
	//
	//  At expressions are useful for one-time schedules. Cron expressions are useful for scheduled actions that run periodically at a specified date and time, and rate expressions are useful for scheduled actions that run at a regular interval.
	//  At and cron expressions use Universal Coordinated Time (UTC) by default.
	//  The cron format consists of six fields separated by white spaces: [Minutes] [Hours] [Day_of_Month] [Month] [Day_of_Week] [Year].
	//  For rate expressions, *value* is a positive integer and *unit* is ``minute`` | ``minutes`` | ``hour`` | ``hours`` | ``day`` | ``days``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#schedule ApplicationautoscalingScalableTarget#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The name of the scheduled action.
	//
	// This name must be unique among all other scheduled actions on the specified scalable target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#scheduled_action_name ApplicationautoscalingScalableTarget#scheduled_action_name}
	ScheduledActionName *string `field:"optional" json:"scheduledActionName" yaml:"scheduledActionName"`
	// The date and time that the action is scheduled to begin, in UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#start_time ApplicationautoscalingScalableTarget#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The time zone used when referring to the date and time of a scheduled action, when the scheduled action uses an at or cron expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationautoscaling_scalable_target#timezone ApplicationautoscalingScalableTarget#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

