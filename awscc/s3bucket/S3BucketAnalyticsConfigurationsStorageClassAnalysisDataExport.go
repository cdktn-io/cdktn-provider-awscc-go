// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketAnalyticsConfigurationsStorageClassAnalysisDataExport struct {
	// The place to store the data for an analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketAnalyticsConfigurationsStorageClassAnalysisDataExportDestination `field:"optional" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data. Must be ``V_1``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3_bucket#output_schema_version S3Bucket#output_schema_version}
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

