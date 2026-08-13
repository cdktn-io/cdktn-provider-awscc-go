// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment


type EvidentlyExperimentOnlineAbConfigTreatmentWeights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#split_weight EvidentlyExperiment#split_weight}.
	SplitWeight *float64 `field:"optional" json:"splitWeight" yaml:"splitWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/evidently_experiment#treatment EvidentlyExperiment#treatment}.
	Treatment *string `field:"optional" json:"treatment" yaml:"treatment"`
}

