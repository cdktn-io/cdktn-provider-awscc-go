// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#filters FisExperimentTemplate#filters}.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#parameters FisExperimentTemplate#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Names (ARNs) of the target resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#resource_arns FisExperimentTemplate#resource_arns}
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#resource_tags FisExperimentTemplate#resource_tags}.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The AWS resource type. The resource type must be supported for the specified action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#resource_type FisExperimentTemplate#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#selection_mode FisExperimentTemplate#selection_mode}
	SelectionMode *string `field:"optional" json:"selectionMode" yaml:"selectionMode"`
}

