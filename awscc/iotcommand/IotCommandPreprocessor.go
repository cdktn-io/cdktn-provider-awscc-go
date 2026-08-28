// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand


type IotCommandPreprocessor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_command#aws_json_substitution IotCommand#aws_json_substitution}.
	AwsJsonSubstitution *IotCommandPreprocessorAwsJsonSubstitution `field:"optional" json:"awsJsonSubstitution" yaml:"awsJsonSubstitution"`
}

