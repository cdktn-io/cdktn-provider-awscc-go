// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveSli struct {
	// The arithmetic operation used when comparing the specified metric to the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#comparison_operator ApplicationsignalsServiceLevelObjective#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value that the SLI metric is compared to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_threshold ApplicationsignalsServiceLevelObjective#metric_threshold}
	MetricThreshold *float64 `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// A structure that contains information about the metric that the SLO monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#sli_metric ApplicationsignalsServiceLevelObjective#sli_metric}
	SliMetric *ApplicationsignalsServiceLevelObjectiveSliSliMetric `field:"optional" json:"sliMetric" yaml:"sliMetric"`
}

