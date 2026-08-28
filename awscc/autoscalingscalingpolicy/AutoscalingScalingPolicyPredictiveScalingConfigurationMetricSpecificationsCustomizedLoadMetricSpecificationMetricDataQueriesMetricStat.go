// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscalingpolicy


type AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedLoadMetricSpecificationMetricDataQueriesMetricStat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#metric AutoscalingScalingPolicy#metric}.
	Metric *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedLoadMetricSpecificationMetricDataQueriesMetricStatMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#stat AutoscalingScalingPolicy#stat}.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#unit AutoscalingScalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

