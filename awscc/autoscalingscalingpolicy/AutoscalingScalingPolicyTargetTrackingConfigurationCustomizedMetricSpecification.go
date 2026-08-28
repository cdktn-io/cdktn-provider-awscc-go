// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscalingpolicy


type AutoscalingScalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#dimensions AutoscalingScalingPolicy#dimensions}.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#metric_name AutoscalingScalingPolicy#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#metrics AutoscalingScalingPolicy#metrics}.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#namespace AutoscalingScalingPolicy#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#period AutoscalingScalingPolicy#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#statistic AutoscalingScalingPolicy#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/autoscaling_scaling_policy#unit AutoscalingScalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

