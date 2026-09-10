// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentrun

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigExperimentRunConfig struct {
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
	// The application name or ID used to create the experiment run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#application_identifier AppconfigExperimentRun#application_identifier}
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The experiment definition name or ID used to create the experiment run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#experiment_definition_identifier AppconfigExperimentRun#experiment_definition_identifier}
	ExperimentDefinitionIdentifier *string `field:"required" json:"experimentDefinitionIdentifier" yaml:"experimentDefinitionIdentifier"`
	// Percentage of traffic exposed to the experiment (0-100).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#exposure_percentage AppconfigExperimentRun#exposure_percentage}
	ExposurePercentage *float64 `field:"required" json:"exposurePercentage" yaml:"exposurePercentage"`
	// Description of the experiment run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#description AppconfigExperimentRun#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to associate with the experiment run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#tags AppconfigExperimentRun#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Treatment overrides for specific entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appconfig_experiment_run#treatment_overrides AppconfigExperimentRun#treatment_overrides}
	TreatmentOverrides *AppconfigExperimentRunTreatmentOverrides `field:"optional" json:"treatmentOverrides" yaml:"treatmentOverrides"`
}

