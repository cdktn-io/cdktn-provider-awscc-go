// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stepfunctionsstatemachine

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StepfunctionsStateMachineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#role_arn StepfunctionsStateMachine#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#definition StepfunctionsStateMachine#definition}.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#definition_s3_location StepfunctionsStateMachine#definition_s3_location}.
	DefinitionS3Location *StepfunctionsStateMachineDefinitionS3Location `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#definition_string StepfunctionsStateMachine#definition_string}.
	DefinitionString *string `field:"optional" json:"definitionString" yaml:"definitionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#definition_substitutions StepfunctionsStateMachine#definition_substitutions}.
	DefinitionSubstitutions *map[string]*string `field:"optional" json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#logging_configuration StepfunctionsStateMachine#logging_configuration}.
	LoggingConfiguration *StepfunctionsStateMachineLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#state_machine_name StepfunctionsStateMachine#state_machine_name}.
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#state_machine_type StepfunctionsStateMachine#state_machine_type}.
	StateMachineType *string `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#tags StepfunctionsStateMachine#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine#tracing_configuration StepfunctionsStateMachine#tracing_configuration}.
	TracingConfiguration *StepfunctionsStateMachineTracingConfiguration `field:"optional" json:"tracingConfiguration" yaml:"tracingConfiguration"`
}

