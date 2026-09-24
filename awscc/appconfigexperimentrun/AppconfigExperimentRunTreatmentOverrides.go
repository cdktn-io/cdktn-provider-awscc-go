// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigexperimentrun


type AppconfigExperimentRunTreatmentOverrides struct {
	// Map of entity ID to treatment key (t1, t2, ..., or c for control).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appconfig_experiment_run#inline AppconfigExperimentRun#inline}
	Inline *map[string]*string `field:"optional" json:"inline" yaml:"inline"`
}

