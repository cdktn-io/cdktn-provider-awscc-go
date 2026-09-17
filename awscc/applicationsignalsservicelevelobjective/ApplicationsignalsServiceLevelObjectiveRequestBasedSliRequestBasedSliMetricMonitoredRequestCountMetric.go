// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetricMonitoredRequestCountMetric struct {
	// If you want to count "bad requests" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as "bad requests" in this structure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#bad_count_metric ApplicationsignalsServiceLevelObjective#bad_count_metric}
	BadCountMetric interface{} `field:"optional" json:"badCountMetric" yaml:"badCountMetric"`
	// If you want to count "good requests" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as "good requests" in this structure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#good_count_metric ApplicationsignalsServiceLevelObjective#good_count_metric}
	GoodCountMetric interface{} `field:"optional" json:"goodCountMetric" yaml:"goodCountMetric"`
}

