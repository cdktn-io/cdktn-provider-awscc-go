// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSource struct {
	// Optional additional attributes for the metric source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#metric_source_attributes ApplicationsignalsServiceLevelObjective#metric_source_attributes}
	MetricSourceAttributes *map[string]*string `field:"optional" json:"metricSourceAttributes" yaml:"metricSourceAttributes"`
	// Required attributes that identify the metric source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#metric_source_key_attributes ApplicationsignalsServiceLevelObjective#metric_source_key_attributes}
	MetricSourceKeyAttributes *map[string]*string `field:"optional" json:"metricSourceKeyAttributes" yaml:"metricSourceKeyAttributes"`
}

