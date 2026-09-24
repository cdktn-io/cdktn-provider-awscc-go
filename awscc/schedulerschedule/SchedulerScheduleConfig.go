// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SchedulerScheduleConfig struct {
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
	// Flexible time window allows configuration of a window within which a schedule can be invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#flexible_time_window SchedulerSchedule#flexible_time_window}
	FlexibleTimeWindow *SchedulerScheduleFlexibleTimeWindow `field:"required" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// The scheduling expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#schedule_expression SchedulerSchedule#schedule_expression}
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The schedule target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#target SchedulerSchedule#target}
	Target *SchedulerScheduleTarget `field:"required" json:"target" yaml:"target"`
	// The description of the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#description SchedulerSchedule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date, in UTC, before which the schedule can invoke its target.
	//
	// Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#end_date SchedulerSchedule#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The name of the schedule group to associate with this schedule.
	//
	// If you omit this, the default schedule group is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#group_name SchedulerSchedule#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The ARN for a KMS Key that will be used to encrypt customer data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#kms_key_arn SchedulerSchedule#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#name SchedulerSchedule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The timezone in which the scheduling expression is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#schedule_expression_timezone SchedulerSchedule#schedule_expression_timezone}
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// The date, in UTC, after which the schedule can begin invoking its target.
	//
	// Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#start_date SchedulerSchedule#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Specifies whether the schedule is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/scheduler_schedule#state SchedulerSchedule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

