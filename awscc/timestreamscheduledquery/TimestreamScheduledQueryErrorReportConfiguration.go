// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamscheduledquery


type TimestreamScheduledQueryErrorReportConfiguration struct {
	// Details on S3 location for error reports that result from running a query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/timestream_scheduled_query#s3_configuration TimestreamScheduledQuery#s3_configuration}
	S3Configuration *TimestreamScheduledQueryErrorReportConfigurationS3Configuration `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
}

