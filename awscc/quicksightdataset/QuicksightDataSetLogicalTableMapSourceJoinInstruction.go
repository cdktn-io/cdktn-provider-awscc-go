// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapSourceJoinInstruction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#left_join_key_properties QuicksightDataSet#left_join_key_properties}.
	LeftJoinKeyProperties *QuicksightDataSetLogicalTableMapSourceJoinInstructionLeftJoinKeyProperties `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// <p>Left operand.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#left_operand QuicksightDataSet#left_operand}
	LeftOperand *string `field:"optional" json:"leftOperand" yaml:"leftOperand"`
	// <p>On Clause.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#on_clause QuicksightDataSet#on_clause}
	OnClause *string `field:"optional" json:"onClause" yaml:"onClause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#right_join_key_properties QuicksightDataSet#right_join_key_properties}.
	RightJoinKeyProperties *QuicksightDataSetLogicalTableMapSourceJoinInstructionRightJoinKeyProperties `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
	// <p>Right operand.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#right_operand QuicksightDataSet#right_operand}
	RightOperand *string `field:"optional" json:"rightOperand" yaml:"rightOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#type QuicksightDataSet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

