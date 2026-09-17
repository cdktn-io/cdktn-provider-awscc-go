// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codebuildreportgroup


type CodebuildReportGroupExportConfigS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#bucket CodebuildReportGroup#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#bucket_owner CodebuildReportGroup#bucket_owner}.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#encryption_disabled CodebuildReportGroup#encryption_disabled}.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#encryption_key CodebuildReportGroup#encryption_key}.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#packaging CodebuildReportGroup#packaging}.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_report_group#path CodebuildReportGroup#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

