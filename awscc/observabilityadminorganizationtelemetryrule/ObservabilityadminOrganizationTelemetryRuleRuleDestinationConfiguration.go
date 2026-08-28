// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfiguration struct {
	// Telemetry parameters for Cloudtrail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#cloudtrail_parameters ObservabilityadminOrganizationTelemetryRule#cloudtrail_parameters}
	CloudtrailParameters *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationCloudtrailParameters `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// Pattern for telemetry data destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#destination_pattern ObservabilityadminOrganizationTelemetryRule#destination_pattern}
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Type of telemetry destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#destination_type ObservabilityadminOrganizationTelemetryRule#destination_type}
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Telemetry parameters for ELB/NLB Load Balancer Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#elb_load_balancer_logging_parameters ObservabilityadminOrganizationTelemetryRule#elb_load_balancer_logging_parameters}
	ElbLoadBalancerLoggingParameters *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParameters `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// The Amazon Resource Name (ARN) of the customer-managed AWS KMS key used to encrypt the destination log groups specified in the Telemetry Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#kms_key_arn ObservabilityadminOrganizationTelemetryRule#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Parameters for log delivery configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#log_delivery_parameters ObservabilityadminOrganizationTelemetryRule#log_delivery_parameters}
	LogDeliveryParameters *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationLogDeliveryParameters `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// Number of days to retain the telemetry data in the specified destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#retention_in_days ObservabilityadminOrganizationTelemetryRule#retention_in_days}
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Telemetry parameters for VPC Flow logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#vpc_flow_log_parameters ObservabilityadminOrganizationTelemetryRule#vpc_flow_log_parameters}
	VpcFlowLogParameters *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// Telemetry parameters for WAF v2 Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#waf_logging_parameters ObservabilityadminOrganizationTelemetryRule#waf_logging_parameters}
	WafLoggingParameters *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationWafLoggingParameters `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

