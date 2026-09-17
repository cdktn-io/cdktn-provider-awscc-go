// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package curreportdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CurReportDefinitionConfig struct {
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
	// The compression format that AWS uses for the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#compression CurReportDefinition#compression}
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// The format that AWS saves the report in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#format CurReportDefinition#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months.
	//
	// These charges can include refunds, credits, or support fees.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#refresh_closed_reports CurReportDefinition#refresh_closed_reports}
	RefreshClosedReports interface{} `field:"required" json:"refreshClosedReports" yaml:"refreshClosedReports"`
	// The name of the report that you want to create.
	//
	// The name must be unique, is case sensitive, and can't include spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#report_name CurReportDefinition#report_name}
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#report_versioning CurReportDefinition#report_versioning}
	ReportVersioning *string `field:"required" json:"reportVersioning" yaml:"reportVersioning"`
	// The S3 bucket where AWS delivers the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#s3_bucket CurReportDefinition#s3_bucket}
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#s3_prefix CurReportDefinition#s3_prefix}
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// The region of the S3 bucket that AWS delivers the report into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#s3_region CurReportDefinition#s3_region}
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// The granularity of the line items in the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#time_unit CurReportDefinition#time_unit}
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// A list of manifests that you want Amazon Web Services to create for this report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#additional_artifacts CurReportDefinition#additional_artifacts}
	AdditionalArtifacts *[]*string `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#additional_schema_elements CurReportDefinition#additional_schema_elements}
	AdditionalSchemaElements *[]*string `field:"optional" json:"additionalSchemaElements" yaml:"additionalSchemaElements"`
	// The Amazon resource name of the billing view.
	//
	// You can get this value by using the billing view service public APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#billing_view_arn CurReportDefinition#billing_view_arn}
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cur_report_definition#tags CurReportDefinition#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

