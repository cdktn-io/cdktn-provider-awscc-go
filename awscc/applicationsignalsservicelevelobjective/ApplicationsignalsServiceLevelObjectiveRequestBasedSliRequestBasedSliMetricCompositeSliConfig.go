// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricCompositeSliConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_service_level_objective#composite_sli_components ApplicationsignalsServiceLevelObjective#composite_sli_components}.
	CompositeSliComponents interface{} `field:"optional" json:"compositeSliComponents" yaml:"compositeSliComponents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_service_level_objective#selection_config ApplicationsignalsServiceLevelObjective#selection_config}.
	SelectionConfig *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricCompositeSliConfigSelectionConfig `field:"optional" json:"selectionConfig" yaml:"selectionConfig"`
}

