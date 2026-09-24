// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmscheduledreport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BcmScheduledReportConfig struct {
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
	// The ARN of the dashboard associated with the scheduled report. Managed dashboards cannot be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#dashboard_arn BcmScheduledReport#dashboard_arn}
	DashboardArn *string `field:"required" json:"dashboardArn" yaml:"dashboardArn"`
	// The name of the scheduled report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#name BcmScheduledReport#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule configuration that defines when and how often the report is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#schedule_config BcmScheduledReport#schedule_config}
	ScheduleConfig *BcmScheduledReportScheduleConfig `field:"required" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The ARN of the IAM role that the scheduled report uses to execute.
	//
	// AWS Billing and Cost Management Dashboards assumes this IAM role while executing the scheduled report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#scheduled_report_execution_role_arn BcmScheduledReport#scheduled_report_execution_role_arn}
	ScheduledReportExecutionRoleArn *string `field:"required" json:"scheduledReportExecutionRoleArn" yaml:"scheduledReportExecutionRoleArn"`
	// A description of the scheduled report's purpose or contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#description BcmScheduledReport#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags applied to the scheduled report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#tags BcmScheduledReport#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The date range override applied to widgets in the scheduled report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#widget_date_range_override BcmScheduledReport#widget_date_range_override}
	WidgetDateRangeOverride *BcmScheduledReportWidgetDateRangeOverride `field:"optional" json:"widgetDateRangeOverride" yaml:"widgetDateRangeOverride"`
	// The list of widget identifiers included in the scheduled report.
	//
	// If not specified, all widgets in the dashboard are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bcm_scheduled_report#widget_ids BcmScheduledReport#widget_ids}
	WidgetIds *[]*string `field:"optional" json:"widgetIds" yaml:"widgetIds"`
}

