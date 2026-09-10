// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerexperimenttrialcomponent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerExperimentTrialComponentConfig struct {
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
	// The name of the trial component. The name must be unique in your AWS account and is not case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#trial_component_name SagemakerExperimentTrialComponent#trial_component_name}
	TrialComponentName *string `field:"required" json:"trialComponentName" yaml:"trialComponentName"`
	// The name of the component as displayed.
	//
	// The name doesn't need to be unique. If DisplayName isn't specified, TrialComponentName is displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#display_name SagemakerExperimentTrialComponent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// When the component ended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#end_time SagemakerExperimentTrialComponent#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#metadata_properties SagemakerExperimentTrialComponent#metadata_properties}
	MetadataProperties *SagemakerExperimentTrialComponentMetadataProperties `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// When the component started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#start_time SagemakerExperimentTrialComponent#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The status of the trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#status SagemakerExperimentTrialComponent#status}
	Status *SagemakerExperimentTrialComponentStatus `field:"optional" json:"status" yaml:"status"`
	// A list of tags to associate with the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_experiment_trial_component#tags SagemakerExperimentTrialComponent#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

