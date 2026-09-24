// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricTotalRequestCountMetricMetricStat struct {
	// This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#metric ApplicationsignalsServiceLevelObjective#metric}
	Metric *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricTotalRequestCountMetricMetricStatMetric `field:"optional" json:"metric" yaml:"metric"`
	// The granularity, in seconds, to be used for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#period ApplicationsignalsServiceLevelObjective#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#stat ApplicationsignalsServiceLevelObjective#stat}
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch.
	//
	// If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#unit ApplicationsignalsServiceLevelObjective#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

