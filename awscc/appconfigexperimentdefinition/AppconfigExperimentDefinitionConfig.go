// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigExperimentDefinitionConfig struct {
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
	// The application ID, name, or ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#application_identifier AppconfigExperimentDefinition#application_identifier}
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// Rule expression defining the experiment audience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#audience_rule AppconfigExperimentDefinition#audience_rule}
	AudienceRule *string `field:"required" json:"audienceRule" yaml:"audienceRule"`
	// The configuration profile ID, name, or ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#configuration_profile_identifier AppconfigExperimentDefinition#configuration_profile_identifier}
	ConfigurationProfileIdentifier *string `field:"required" json:"configurationProfileIdentifier" yaml:"configurationProfileIdentifier"`
	// The control (baseline) variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#control AppconfigExperimentDefinition#control}
	Control *AppconfigExperimentDefinitionControl `field:"required" json:"control" yaml:"control"`
	// The environment ID, name, or ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#environment_identifier AppconfigExperimentDefinition#environment_identifier}
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The key of the existing flag in the configuration profile to experiment on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#flag_key AppconfigExperimentDefinition#flag_key}
	FlagKey *string `field:"required" json:"flagKey" yaml:"flagKey"`
	// A name for the experiment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#name AppconfigExperimentDefinition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Treatment variants (1-5).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#treatments AppconfigExperimentDefinition#treatments}
	Treatments interface{} `field:"required" json:"treatments" yaml:"treatments"`
	// Human-readable description of the audience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#audience_description AppconfigExperimentDefinition#audience_description}
	AudienceDescription *string `field:"optional" json:"audienceDescription" yaml:"audienceDescription"`
	// The hypothesis of the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#hypothesis AppconfigExperimentDefinition#hypothesis}
	Hypothesis *string `field:"optional" json:"hypothesis" yaml:"hypothesis"`
	// Criteria for launching the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#launch_criteria AppconfigExperimentDefinition#launch_criteria}
	LaunchCriteria *string `field:"optional" json:"launchCriteria" yaml:"launchCriteria"`
	// Tags to associate with the experiment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_definition#tags AppconfigExperimentDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

