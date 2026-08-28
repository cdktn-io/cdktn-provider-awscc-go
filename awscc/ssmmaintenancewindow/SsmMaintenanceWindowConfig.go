// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmMaintenanceWindowConfig struct {
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
	// Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets.
	//
	// If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#allow_unassociated_targets SsmMaintenanceWindow#allow_unassociated_targets}
	AllowUnassociatedTargets interface{} `field:"required" json:"allowUnassociatedTargets" yaml:"allowUnassociatedTargets"`
	// The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#cutoff SsmMaintenanceWindow#cutoff}
	Cutoff *float64 `field:"required" json:"cutoff" yaml:"cutoff"`
	// The duration of the maintenance window in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#duration SsmMaintenanceWindow#duration}
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The name of the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#name SsmMaintenanceWindow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule of the maintenance window in the form of a cron or rate expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#schedule SsmMaintenanceWindow#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// A description of the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#description SsmMaintenanceWindow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#end_date SsmMaintenanceWindow#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The number of days to wait to run a maintenance window after the scheduled cron expression date and time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#schedule_offset SsmMaintenanceWindow#schedule_offset}
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#schedule_timezone SsmMaintenanceWindow#schedule_timezone}
	ScheduleTimezone *string `field:"optional" json:"scheduleTimezone" yaml:"scheduleTimezone"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.
	//
	// StartDate allows you to delay activation of the maintenance window until the specified future date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#start_date SsmMaintenanceWindow#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_maintenance_window#tags SsmMaintenanceWindow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

