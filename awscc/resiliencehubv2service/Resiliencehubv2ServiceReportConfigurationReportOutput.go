// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceReportConfigurationReportOutput struct {
	// S3 configuration for report output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/resiliencehubv2_service#s3 Resiliencehubv2Service#s3}
	S3 *Resiliencehubv2ServiceReportConfigurationReportOutputS3 `field:"optional" json:"s3" yaml:"s3"`
}

