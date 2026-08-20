// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stepfunctionsstatemachine


type StepfunctionsStateMachineDefinitionS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/stepfunctions_state_machine#bucket StepfunctionsStateMachine#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/stepfunctions_state_machine#key StepfunctionsStateMachine#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/stepfunctions_state_machine#version StepfunctionsStateMachine#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

