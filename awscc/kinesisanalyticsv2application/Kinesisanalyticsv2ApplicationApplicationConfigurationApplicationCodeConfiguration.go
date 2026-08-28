// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration struct {
	// The location and type of the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#code_content Kinesisanalyticsv2Application#code_content}
	CodeContent *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent `field:"optional" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisanalyticsv2_application#code_content_type Kinesisanalyticsv2Application#code_content_type}
	CodeContentType *string `field:"optional" json:"codeContentType" yaml:"codeContentType"`
}

