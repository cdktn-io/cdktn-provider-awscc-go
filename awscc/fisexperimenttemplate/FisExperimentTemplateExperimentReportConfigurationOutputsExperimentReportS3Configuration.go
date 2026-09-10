// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateExperimentReportConfigurationOutputsExperimentReportS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fis_experiment_template#bucket_name FisExperimentTemplate#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fis_experiment_template#prefix FisExperimentTemplate#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

