// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviorsCriteria struct {
	// The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#comparison_operator IotSecurityProfile#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs.
	//
	// If not specified, the default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#consecutive_datapoints_to_alarm IotSecurityProfile#consecutive_datapoints_to_alarm}
	ConsecutiveDatapointsToAlarm *float64 `field:"optional" json:"consecutiveDatapointsToAlarm" yaml:"consecutiveDatapointsToAlarm"`
	// If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared.
	//
	// If not specified, the default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#consecutive_datapoints_to_clear IotSecurityProfile#consecutive_datapoints_to_clear}
	ConsecutiveDatapointsToClear *float64 `field:"optional" json:"consecutiveDatapointsToClear" yaml:"consecutiveDatapointsToClear"`
	// Use this to specify the time duration over which the behavior is evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#duration_seconds IotSecurityProfile#duration_seconds}
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The configuration of an ML Detect Security Profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#ml_detection_config IotSecurityProfile#ml_detection_config}
	MlDetectionConfig *IotSecurityProfileBehaviorsCriteriaMlDetectionConfig `field:"optional" json:"mlDetectionConfig" yaml:"mlDetectionConfig"`
	// A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#statistical_threshold IotSecurityProfile#statistical_threshold}
	StatisticalThreshold *IotSecurityProfileBehaviorsCriteriaStatisticalThreshold `field:"optional" json:"statisticalThreshold" yaml:"statisticalThreshold"`
	// The value to be compared with the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_security_profile#value IotSecurityProfile#value}
	Value *IotSecurityProfileBehaviorsCriteriaValue `field:"optional" json:"value" yaml:"value"`
}

