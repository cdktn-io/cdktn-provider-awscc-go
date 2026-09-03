// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontmonitoringsubscription


type CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig struct {
	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_monitoring_subscription#realtime_metrics_subscription_status CloudfrontMonitoringSubscription#realtime_metrics_subscription_status}
	RealtimeMetricsSubscriptionStatus *string `field:"optional" json:"realtimeMetricsSubscriptionStatus" yaml:"realtimeMetricsSubscriptionStatus"`
}

