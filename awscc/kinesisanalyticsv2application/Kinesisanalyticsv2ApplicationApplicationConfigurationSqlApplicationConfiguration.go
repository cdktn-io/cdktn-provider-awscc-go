// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration struct {
	// The array of Input objects describing the input streams used by the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#inputs Kinesisanalyticsv2Application#inputs}
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
}

