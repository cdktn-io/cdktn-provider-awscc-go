// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codebuildreportgroup


type CodebuildReportGroupExportConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codebuild_report_group#export_config_type CodebuildReportGroup#export_config_type}.
	ExportConfigType *string `field:"required" json:"exportConfigType" yaml:"exportConfigType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codebuild_report_group#s3_destination CodebuildReportGroup#s3_destination}.
	S3Destination *CodebuildReportGroupExportConfigS3Destination `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

