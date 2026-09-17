// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesCreateRule struct {
	// The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1 year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#cron_expression DlmLifecyclePolicy#cron_expression}
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// The interval between snapshots. The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#interval DlmLifecyclePolicy#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The interval unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#interval_unit DlmLifecyclePolicy#interval_unit}
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// **[Custom snapshot policies only]** Specifies the destination for snapshots created by the policy.
	//
	// The allowed destinations depend on the location of the targeted resources.
	//
	// - If the policy targets resources in a Region, then you must create snapshots in the same Region as the source resource.
	// - If the policy targets resources in a Local Zone, you can create snapshots in the same Local Zone or in its parent Region.
	// - If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost or in its parent Region.
	//
	// Default: `CLOUD`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#location DlmLifecyclePolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// **[Custom snapshot policies that target instances only]** Specifies pre and/or post scripts for a snapshot lifecycle policy that targets instances.
	//
	// This is useful for creating application-consistent snapshots, or for performing specific administrative tasks before or after Amazon Data Lifecycle Manager initiates snapshot creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#scripts DlmLifecyclePolicy#scripts}
	Scripts interface{} `field:"optional" json:"scripts" yaml:"scripts"`
	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon Data Lifecycle Manager selects a time within the next 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#times DlmLifecyclePolicy#times}
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

