// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent struct {
	// Information about the Amazon S3 bucket that contains the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#s3_content_location Kinesisanalyticsv2Application#s3_content_location}
	S3ContentLocation *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
	// The text-format code for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#text_content Kinesisanalyticsv2Application#text_content}
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
	// The zip-format code for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#zip_file_content Kinesisanalyticsv2Application#zip_file_content}
	ZipFileContent *string `field:"optional" json:"zipFileContent" yaml:"zipFileContent"`
}

