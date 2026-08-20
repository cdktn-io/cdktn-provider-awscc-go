// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMonitoredRequestCountMetricBadCountMetric struct {
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#account_id ApplicationsignalsServiceLevelObjective#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The math expression to be performed on the returned data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#expression ApplicationsignalsServiceLevelObjective#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A short name used to tie this object to the results in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#id ApplicationsignalsServiceLevelObjective#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO.
	//
	// Within one MetricDataQuery, you must specify either Expression or MetricStat but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#metric_stat ApplicationsignalsServiceLevelObjective#metric_stat}
	MetricStat *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMonitoredRequestCountMetricBadCountMetricMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// This option indicates whether to return the timestamps and raw data values of this metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#return_data ApplicationsignalsServiceLevelObjective#return_data}
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

