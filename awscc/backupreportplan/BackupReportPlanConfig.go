// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupreportplan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BackupReportPlanConfig struct {
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
	// A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_report_plan#report_delivery_channel BackupReportPlan#report_delivery_channel}
	ReportDeliveryChannel *BackupReportPlanReportDeliveryChannel `field:"required" json:"reportDeliveryChannel" yaml:"reportDeliveryChannel"`
	// Identifies the report template for the report. Reports are built using a report template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_report_plan#report_setting BackupReportPlan#report_setting}
	ReportSetting *BackupReportPlanReportSetting `field:"required" json:"reportSetting" yaml:"reportSetting"`
	// An optional description of the report plan with a maximum of 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_report_plan#report_plan_description BackupReportPlan#report_plan_description}
	ReportPlanDescription *string `field:"optional" json:"reportPlanDescription" yaml:"reportPlanDescription"`
	// The unique name of the report plan.
	//
	// The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_report_plan#report_plan_name BackupReportPlan#report_plan_name}
	ReportPlanName *string `field:"optional" json:"reportPlanName" yaml:"reportPlanName"`
	// Metadata that you can assign to help organize the report plans that you create.
	//
	// Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_report_plan#report_plan_tags BackupReportPlan#report_plan_tags}
	ReportPlanTags interface{} `field:"optional" json:"reportPlanTags" yaml:"reportPlanTags"`
}

