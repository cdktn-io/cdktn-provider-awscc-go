// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment


type EvidentlyExperimentOnlineAbConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/evidently_experiment#control_treatment_name EvidentlyExperiment#control_treatment_name}.
	ControlTreatmentName *string `field:"optional" json:"controlTreatmentName" yaml:"controlTreatmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/evidently_experiment#treatment_weights EvidentlyExperiment#treatment_weights}.
	TreatmentWeights interface{} `field:"optional" json:"treatmentWeights" yaml:"treatmentWeights"`
}

