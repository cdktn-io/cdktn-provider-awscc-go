// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveSliSliMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#composite_sli_config ApplicationsignalsServiceLevelObjective#composite_sli_config}.
	CompositeSliConfig *ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfig `field:"optional" json:"compositeSliConfig" yaml:"compositeSliConfig"`
	// Configuration for identifying a dependency and its operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#dependency_config ApplicationsignalsServiceLevelObjective#dependency_config}
	DependencyConfig *ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfig `field:"optional" json:"dependencyConfig" yaml:"dependencyConfig"`
	// This is a string-to-string map that contains information about the type of object that this SLO is related to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#key_attributes ApplicationsignalsServiceLevelObjective#key_attributes}
	KeyAttributes *map[string]*string `field:"optional" json:"keyAttributes" yaml:"keyAttributes"`
	// If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, this structure includes the information about that metric or expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_data_queries ApplicationsignalsServiceLevelObjective#metric_data_queries}
	MetricDataQueries interface{} `field:"optional" json:"metricDataQueries" yaml:"metricDataQueries"`
	// The name of the metric for non-Application Signals services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_name ApplicationsignalsServiceLevelObjective#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Configuration for identifying the source of metrics for non-Application Signals services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_source ApplicationsignalsServiceLevelObjective#metric_source}
	MetricSource *ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSource `field:"optional" json:"metricSource" yaml:"metricSource"`
	// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_type ApplicationsignalsServiceLevelObjective#metric_type}
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// If the SLO monitors a specific operation of the service, this field displays that operation name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#operation_name ApplicationsignalsServiceLevelObjective#operation_name}
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The number of seconds to use as the period for SLO evaluation.
	//
	// Your application's performance is compared to the SLI during each period. For each period, the application is determined to have either achieved or not achieved the necessary performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#period_seconds ApplicationsignalsServiceLevelObjective#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#statistic ApplicationsignalsServiceLevelObjective#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

