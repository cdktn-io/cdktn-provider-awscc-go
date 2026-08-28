// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleWeightConfig struct {
	// Session stickiness provides the ability to define multiple requests from a single viewer as a single session.
	//
	// This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_continuous_deployment_policy#session_stickiness_config CloudfrontContinuousDeploymentPolicy#session_stickiness_config}
	SessionStickinessConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleWeightConfigSessionStickinessConfig `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
	// The percentage of traffic to send to a staging distribution, expressed as a decimal number between 0 and 0.15. For example, a value of 0.10 means 10% of traffic is sent to the staging distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_continuous_deployment_policy#weight CloudfrontContinuousDeploymentPolicy#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

