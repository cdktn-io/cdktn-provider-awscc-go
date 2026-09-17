// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stepfunctionsstatemachinealias


type StepfunctionsStateMachineAliasDeploymentPreference struct {
	// A list of CloudWatch alarm names that will be monitored during the deployment.
	//
	// The deployment will fail and rollback if any alarms go into ALARM state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine_alias#alarms StepfunctionsStateMachineAlias#alarms}
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// The time in minutes between each traffic shifting increment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine_alias#interval StepfunctionsStateMachineAlias#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The percentage of traffic to shift to the new version in each increment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine_alias#percentage StepfunctionsStateMachineAlias#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine_alias#state_machine_version_arn StepfunctionsStateMachineAlias#state_machine_version_arn}.
	StateMachineVersionArn *string `field:"optional" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// The type of deployment to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/stepfunctions_state_machine_alias#type StepfunctionsStateMachineAlias#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

