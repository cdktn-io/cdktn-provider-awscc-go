// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotscheduledaudit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotScheduledAuditConfig struct {
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
	// How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#frequency IotScheduledAudit#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Which checks are performed during the scheduled audit. Checks must be enabled for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#target_check_names IotScheduledAudit#target_check_names}
	TargetCheckNames *[]*string `field:"required" json:"targetCheckNames" yaml:"targetCheckNames"`
	// The day of the month on which the scheduled audit takes place.
	//
	// Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#day_of_month IotScheduledAudit#day_of_month}
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week on which the scheduled audit takes place.
	//
	// Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#day_of_week IotScheduledAudit#day_of_week}
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// The name you want to give to the scheduled audit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#scheduled_audit_name IotScheduledAudit#scheduled_audit_name}
	ScheduledAuditName *string `field:"optional" json:"scheduledAuditName" yaml:"scheduledAuditName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_scheduled_audit#tags IotScheduledAudit#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

