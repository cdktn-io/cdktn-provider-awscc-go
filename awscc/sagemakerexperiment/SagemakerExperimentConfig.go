// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerexperiment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerExperimentConfig struct {
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
	// The name of the experiment. Must be unique in your AWS account and is not case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_experiment#experiment_name SagemakerExperiment#experiment_name}
	ExperimentName *string `field:"required" json:"experimentName" yaml:"experimentName"`
	// The description of the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_experiment#description SagemakerExperiment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the experiment as displayed. The name does not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_experiment#display_name SagemakerExperiment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of tags to associate with the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_experiment#tags SagemakerExperiment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

