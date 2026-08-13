// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment


type EvidentlyExperimentRunningStatus struct {
	// Provide the analysis Completion time for an experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#analysis_complete_time EvidentlyExperiment#analysis_complete_time}
	AnalysisCompleteTime *string `field:"optional" json:"analysisCompleteTime" yaml:"analysisCompleteTime"`
	// Provide CANCELLED or COMPLETED desired state when stopping an experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#desired_state EvidentlyExperiment#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Reason is a required input for stopping the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#reason EvidentlyExperiment#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// Provide START or STOP action to apply on an experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#status EvidentlyExperiment#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

