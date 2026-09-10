// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMonitoredRequestCountMetricBadCountMetricMetricStatMetric struct {
	// An array of one or more dimensions to use to define the metric that you want to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationsignals_service_level_objective#dimensions ApplicationsignalsServiceLevelObjective#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationsignals_service_level_objective#metric_name ApplicationsignalsServiceLevelObjective#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationsignals_service_level_objective#namespace ApplicationsignalsServiceLevelObjective#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

