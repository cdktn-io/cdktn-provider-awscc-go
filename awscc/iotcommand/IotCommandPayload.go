// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand


type IotCommandPayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_command#content IotCommand#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_command#content_type IotCommand#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

