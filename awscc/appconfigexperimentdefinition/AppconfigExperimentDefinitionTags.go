// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentdefinition


type AppconfigExperimentDefinitionTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_definition#key AppconfigExperimentDefinition#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_definition#value AppconfigExperimentDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

