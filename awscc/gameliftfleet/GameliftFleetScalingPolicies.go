// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetScalingPolicies struct {
	// Comparison operator to use when measuring a metric against the threshold value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#comparison_operator GameliftFleet#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#evaluation_periods GameliftFleet#evaluation_periods}
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#location GameliftFleet#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Name of the Amazon GameLift-defined metric that is used to trigger a scaling adjustment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#metric_name GameliftFleet#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#name GameliftFleet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of scaling policy to create.
	//
	// For a target-based policy, set the parameter MetricName to 'PercentAvailableGameSessions' and specify a TargetConfiguration. For a rule-based policy set the following parameters: MetricName, ComparisonOperator, Threshold, EvaluationPeriods, ScalingAdjustmentType, and ScalingAdjustment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#policy_type GameliftFleet#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Amount of adjustment to make, based on the scaling adjustment type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#scaling_adjustment GameliftFleet#scaling_adjustment}
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The type of adjustment to make to a fleet's instance count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#scaling_adjustment_type GameliftFleet#scaling_adjustment_type}
	ScalingAdjustmentType *string `field:"optional" json:"scalingAdjustmentType" yaml:"scalingAdjustmentType"`
	// Current status of the scaling policy.
	//
	// The scaling policy can be in force only when in an ACTIVE status. Scaling policies can be suspended for individual fleets. If the policy is suspended for a fleet, the policy status does not change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#status GameliftFleet#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An object that contains settings for a target-based scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#target_configuration GameliftFleet#target_configuration}
	TargetConfiguration *GameliftFleetScalingPoliciesTargetConfiguration `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// Metric value used to trigger a scaling event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#threshold GameliftFleet#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// The current status of the fleet's scaling policies in a requested fleet location.
	//
	// The status PENDING_UPDATE indicates that an update was requested for the fleet but has not yet been completed for the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/gamelift_fleet#update_status GameliftFleet#update_status}
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

