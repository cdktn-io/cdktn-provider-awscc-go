// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53hostedzone


type Route53HostedZoneQueryLoggingConfig struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/route53_hosted_zone#cloudwatch_logs_log_group_arn Route53HostedZone#cloudwatch_logs_log_group_arn}
	CloudwatchLogsLogGroupArn *string `field:"optional" json:"cloudwatchLogsLogGroupArn" yaml:"cloudwatchLogsLogGroupArn"`
}

