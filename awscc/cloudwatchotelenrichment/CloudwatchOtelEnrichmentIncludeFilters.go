// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchotelenrichment


type CloudwatchOtelEnrichmentIncludeFilters struct {
	// Absent or empty means every metric in this namespace, in whichever direction this selector appears.
	//
	// Present means only these metric names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_otel_enrichment#metric_names CloudwatchOtelEnrichment#metric_names}
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// The CloudWatch namespace this selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_otel_enrichment#namespace CloudwatchOtelEnrichment#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

