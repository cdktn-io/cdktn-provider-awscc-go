// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53healthcheck


type Route53HealthCheckHealthCheckConfigAlarmIdentifier struct {
	// The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_health_check#name Route53HealthCheck#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_health_check#region Route53HealthCheck#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

