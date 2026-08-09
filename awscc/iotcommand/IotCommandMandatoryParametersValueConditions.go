// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand


type IotCommandMandatoryParametersValueConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_command#comparison_operator IotCommand#comparison_operator}.
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_command#operand IotCommand#operand}.
	Operand *IotCommandMandatoryParametersValueConditionsOperand `field:"optional" json:"operand" yaml:"operand"`
}

