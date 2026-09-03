// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsmetricfilter


type LogsMetricFilterMetricTransformationsDimensions struct {
	// The name for the CW metric dimension that the metric filter creates.
	//
	// Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/logs_metric_filter#key LogsMetricFilter#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The log event field that will contain the value for this dimension.
	//
	// This dimension will only be published for a metric if the value is found in the log event. For example, ``$.eventType`` for JSON log events, or ``$server`` for space-delimited log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/logs_metric_filter#value LogsMetricFilter#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

