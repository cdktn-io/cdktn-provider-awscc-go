// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#composite_sli_config ApplicationsignalsServiceLevelObjective#composite_sli_config}.
	CompositeSliConfig *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricCompositeSliConfig `field:"optional" json:"compositeSliConfig" yaml:"compositeSliConfig"`
	// Configuration for identifying a dependency and its operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#dependency_config ApplicationsignalsServiceLevelObjective#dependency_config}
	DependencyConfig *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricDependencyConfig `field:"optional" json:"dependencyConfig" yaml:"dependencyConfig"`
	// This is a string-to-string map that contains information about the type of object that this SLO is related to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#key_attributes ApplicationsignalsServiceLevelObjective#key_attributes}
	KeyAttributes *map[string]*string `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// The name of the metric for non-Application Signals services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#metric_name ApplicationsignalsServiceLevelObjective#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Configuration for identifying the source of metrics for non-Application Signals services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#metric_source ApplicationsignalsServiceLevelObjective#metric_source}
	MetricSource *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMetricSource `field:"optional" json:"metricSource" yaml:"metricSource"`
	// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#metric_type ApplicationsignalsServiceLevelObjective#metric_type}
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// This structure defines the metric that is used as the "good request" or "bad request" value for a request-based SLO.
	//
	// This value observed for the metric defined in `TotalRequestCountMetric` is divided by the number found for `MonitoredRequestCountMetric` to determine the percentage of successful requests that this SLO tracks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#monitored_request_count_metric ApplicationsignalsServiceLevelObjective#monitored_request_count_metric}
	MonitoredRequestCountMetric *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMonitoredRequestCountMetric `field:"optional" json:"monitoredRequestCountMetric" yaml:"monitoredRequestCountMetric"`
	// If the SLO monitors a specific operation of the service, this field displays that operation name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#operation_name ApplicationsignalsServiceLevelObjective#operation_name}
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// This structure defines the metric that is used as the "total requests" number for a request-based SLO.
	//
	// The number observed for this metric is divided by the number of "good requests" or "bad requests" that is observed for the metric defined in `MonitoredRequestCountMetric`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#total_request_count_metric ApplicationsignalsServiceLevelObjective#total_request_count_metric}
	TotalRequestCountMetric interface{} `field:"optional" json:"totalRequestCountMetric" yaml:"totalRequestCountMetric"`
}

