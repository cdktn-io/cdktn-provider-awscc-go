// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53healthcheck


type Route53HealthCheckHealthCheckConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#type Route53HealthCheck#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#alarm_identifier Route53HealthCheck#alarm_identifier}
	AlarmIdentifier *Route53HealthCheckHealthCheckConfigAlarmIdentifier `field:"optional" json:"alarmIdentifier" yaml:"alarmIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#child_health_checks Route53HealthCheck#child_health_checks}.
	ChildHealthChecks *[]*string `field:"optional" json:"childHealthChecks" yaml:"childHealthChecks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#enable_sni Route53HealthCheck#enable_sni}.
	EnableSni interface{} `field:"optional" json:"enableSni" yaml:"enableSni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#failure_threshold Route53HealthCheck#failure_threshold}.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#fully_qualified_domain_name Route53HealthCheck#fully_qualified_domain_name}.
	FullyQualifiedDomainName *string `field:"optional" json:"fullyQualifiedDomainName" yaml:"fullyQualifiedDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#health_threshold Route53HealthCheck#health_threshold}.
	HealthThreshold *float64 `field:"optional" json:"healthThreshold" yaml:"healthThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#insufficient_data_health_status Route53HealthCheck#insufficient_data_health_status}.
	InsufficientDataHealthStatus *string `field:"optional" json:"insufficientDataHealthStatus" yaml:"insufficientDataHealthStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#inverted Route53HealthCheck#inverted}.
	Inverted interface{} `field:"optional" json:"inverted" yaml:"inverted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#ip_address Route53HealthCheck#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#measure_latency Route53HealthCheck#measure_latency}.
	MeasureLatency interface{} `field:"optional" json:"measureLatency" yaml:"measureLatency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#port Route53HealthCheck#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#regions Route53HealthCheck#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#request_interval Route53HealthCheck#request_interval}.
	RequestInterval *float64 `field:"optional" json:"requestInterval" yaml:"requestInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#resource_path Route53HealthCheck#resource_path}.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#routing_control_arn Route53HealthCheck#routing_control_arn}.
	RoutingControlArn *string `field:"optional" json:"routingControlArn" yaml:"routingControlArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/route53_health_check#search_string Route53HealthCheck#search_string}.
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
}

