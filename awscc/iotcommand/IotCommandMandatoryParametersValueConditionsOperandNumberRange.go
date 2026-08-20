// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand


type IotCommandMandatoryParametersValueConditionsOperandNumberRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_command#max IotCommand#max}.
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_command#min IotCommand#min}.
	Min *string `field:"optional" json:"min" yaml:"min"`
}

