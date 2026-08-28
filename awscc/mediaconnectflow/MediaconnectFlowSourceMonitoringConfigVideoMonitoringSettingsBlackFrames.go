// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceMonitoringConfigVideoMonitoringSettingsBlackFrames struct {
	// Indicates whether the BlackFrames metric is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#state MediaconnectFlow#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Specifies the number of consecutive seconds of black frames that triggers an event or alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#threshold_seconds MediaconnectFlow#threshold_seconds}
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

