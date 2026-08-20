// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters struct {
	// The fields to include in the flow log record.
	//
	// If you omit this parameter, the flow log is created using the default format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_telemetry_rule#log_format ObservabilityadminTelemetryRule#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.
	//
	// Default is 600s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_telemetry_rule#max_aggregation_interval ObservabilityadminTelemetryRule#max_aggregation_interval}
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic captured for the flow log. Default is ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_telemetry_rule#traffic_type ObservabilityadminTelemetryRule#traffic_type}
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

