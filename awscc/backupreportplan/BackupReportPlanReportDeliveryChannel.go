// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupreportplan


type BackupReportPlanReportDeliveryChannel struct {
	// The unique name of the S3 bucket that receives your reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_report_plan#s3_bucket_name BackupReportPlan#s3_bucket_name}
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// A list of the format of your reports: CSV, JSON, or both.
	//
	// If not specified, the default format is CSV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_report_plan#formats BackupReportPlan#formats}
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
	// The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3.
	//
	// The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_report_plan#s3_key_prefix BackupReportPlan#s3_key_prefix}
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

