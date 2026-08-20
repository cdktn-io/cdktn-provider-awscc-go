// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateActions struct {
	// The ID of the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_experiment_template#action_id FisExperimentTemplate#action_id}
	ActionId *string `field:"optional" json:"actionId" yaml:"actionId"`
	// A description for the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_experiment_template#description FisExperimentTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameters for the action, if applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_experiment_template#parameters FisExperimentTemplate#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The names of the actions that must be completed before the current action starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_experiment_template#start_after FisExperimentTemplate#start_after}
	StartAfter *[]*string `field:"optional" json:"startAfter" yaml:"startAfter"`
	// One or more targets for the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_experiment_template#targets FisExperimentTemplate#targets}
	Targets *map[string]*string `field:"optional" json:"targets" yaml:"targets"`
}

