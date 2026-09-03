// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupreportplan


type BackupReportPlanReportSetting struct {
	// Identifies the report template for the report.
	//
	// Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_report_plan#report_template BackupReportPlan#report_template}
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// The list of AWS accounts that a report covers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_report_plan#accounts BackupReportPlan#accounts}
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// The Amazon Resource Names (ARNs) of the frameworks a report covers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_report_plan#framework_arns BackupReportPlan#framework_arns}
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// The list of AWS organization units that a report covers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_report_plan#organization_units BackupReportPlan#organization_units}
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// The list of AWS regions that a report covers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_report_plan#regions BackupReportPlan#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

