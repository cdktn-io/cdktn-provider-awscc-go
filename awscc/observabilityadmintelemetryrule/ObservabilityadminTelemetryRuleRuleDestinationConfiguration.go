// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule


type ObservabilityadminTelemetryRuleRuleDestinationConfiguration struct {
	// Telemetry parameters for Cloudtrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#cloudtrail_parameters ObservabilityadminTelemetryRule#cloudtrail_parameters}
	CloudtrailParameters *ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParameters `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// Pattern for telemetry data destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#destination_pattern ObservabilityadminTelemetryRule#destination_pattern}
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Type of telemetry destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#destination_type ObservabilityadminTelemetryRule#destination_type}
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Telemetry parameters for ELB/NLB Load Balancer Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#elb_load_balancer_logging_parameters ObservabilityadminTelemetryRule#elb_load_balancer_logging_parameters}
	ElbLoadBalancerLoggingParameters *ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParameters `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// Parameters for log delivery configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#log_delivery_parameters ObservabilityadminTelemetryRule#log_delivery_parameters}
	LogDeliveryParameters *ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParameters `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// Number of days to retain the telemetry data in the specified destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#retention_in_days ObservabilityadminTelemetryRule#retention_in_days}
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Telemetry parameters for VPC Flow logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#vpc_flow_log_parameters ObservabilityadminTelemetryRule#vpc_flow_log_parameters}
	VpcFlowLogParameters *ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// Telemetry parameters for WAF v2 Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/observabilityadmin_telemetry_rule#waf_logging_parameters ObservabilityadminTelemetryRule#waf_logging_parameters}
	WafLoggingParameters *ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParameters `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

