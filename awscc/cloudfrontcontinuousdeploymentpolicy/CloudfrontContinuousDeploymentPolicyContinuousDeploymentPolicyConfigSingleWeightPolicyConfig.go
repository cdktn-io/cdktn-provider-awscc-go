// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfig struct {
	// Session stickiness provides the ability to define multiple requests from a single viewer as a single session.
	//
	// This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_continuous_deployment_policy#session_stickiness_config CloudfrontContinuousDeploymentPolicy#session_stickiness_config}
	SessionStickinessConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigSessionStickinessConfig `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_continuous_deployment_policy#weight CloudfrontContinuousDeploymentPolicy#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

