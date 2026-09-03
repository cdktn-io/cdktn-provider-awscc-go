// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleWeightConfigSessionStickinessConfig struct {
	// The amount of time after which you want sessions to cease if no requests are received.
	//
	// Allowed values are 300?3600 seconds (5?60 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_continuous_deployment_policy#idle_ttl CloudfrontContinuousDeploymentPolicy#idle_ttl}
	IdleTtl *float64 `field:"optional" json:"idleTtl" yaml:"idleTtl"`
	// The maximum amount of time to consider requests from the viewer as being part of the same session.
	//
	// Allowed values are 300?3600 seconds (5?60 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_continuous_deployment_policy#maximum_ttl CloudfrontContinuousDeploymentPolicy#maximum_ttl}
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
}

