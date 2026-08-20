// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvidentlyExperimentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#metric_goals EvidentlyExperiment#metric_goals}.
	MetricGoals interface{} `field:"required" json:"metricGoals" yaml:"metricGoals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#name EvidentlyExperiment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#online_ab_config EvidentlyExperiment#online_ab_config}.
	OnlineAbConfig *EvidentlyExperimentOnlineAbConfig `field:"required" json:"onlineAbConfig" yaml:"onlineAbConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#project EvidentlyExperiment#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#treatments EvidentlyExperiment#treatments}.
	Treatments interface{} `field:"required" json:"treatments" yaml:"treatments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#description EvidentlyExperiment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#randomization_salt EvidentlyExperiment#randomization_salt}.
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#remove_segment EvidentlyExperiment#remove_segment}.
	RemoveSegment interface{} `field:"optional" json:"removeSegment" yaml:"removeSegment"`
	// Start Experiment. Default is False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#running_status EvidentlyExperiment#running_status}
	RunningStatus *EvidentlyExperimentRunningStatus `field:"optional" json:"runningStatus" yaml:"runningStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#sampling_rate EvidentlyExperiment#sampling_rate}.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#segment EvidentlyExperiment#segment}.
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/evidently_experiment#tags EvidentlyExperiment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

