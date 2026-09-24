// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet


type GameliftContainerFleetScalingPolicies struct {
	// Comparison operator to use when measuring a metric against the threshold value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#comparison_operator GameliftContainerFleet#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#evaluation_periods GameliftContainerFleet#evaluation_periods}
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Name of the Amazon GameLift-defined metric that is used to trigger a scaling adjustment.
	//
	// This is required for RuleBased and TargetBased policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#metric_name GameliftContainerFleet#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#name GameliftContainerFleet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of scaling policy to create.
	//
	// For a target-based policy, set the parameter MetricName to 'PercentAvailableGameSessions' and specify a TargetConfiguration. For a rule-based policy set the following parameters: MetricName, ComparisonOperator, Threshold, EvaluationPeriods, ScalingAdjustmentType, and ScalingAdjustment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#policy_type GameliftContainerFleet#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Amount of adjustment to make, based on the scaling adjustment type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#scaling_adjustment GameliftContainerFleet#scaling_adjustment}
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The type of adjustment to make to a fleet's instance count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#scaling_adjustment_type GameliftContainerFleet#scaling_adjustment_type}
	ScalingAdjustmentType *string `field:"optional" json:"scalingAdjustmentType" yaml:"scalingAdjustmentType"`
	// An object that contains settings for a target-based scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#target_configuration GameliftContainerFleet#target_configuration}
	TargetConfiguration *GameliftContainerFleetScalingPoliciesTargetConfiguration `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// Metric value used to trigger a scaling event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/gamelift_container_fleet#threshold GameliftContainerFleet#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

