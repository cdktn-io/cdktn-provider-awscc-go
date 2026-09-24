// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand


type IotCommandMandatoryParametersValueConditionsOperand struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_command#number IotCommand#number}.
	Number *string `field:"optional" json:"number" yaml:"number"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_command#number_range IotCommand#number_range}.
	NumberRange *IotCommandMandatoryParametersValueConditionsOperandNumberRange `field:"optional" json:"numberRange" yaml:"numberRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_command#numbers IotCommand#numbers}.
	Numbers *[]*string `field:"optional" json:"numbers" yaml:"numbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_command#string IotCommand#string}.
	String *string `field:"optional" json:"string" yaml:"string"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_command#strings IotCommand#strings}.
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
}

