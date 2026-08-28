// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleHeaderConfig struct {
	// The request header name that you want CloudFront to send to your staging distribution.
	//
	// The header must contain the prefix ``aws-cf-cd-``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_continuous_deployment_policy#header CloudfrontContinuousDeploymentPolicy#header}
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The request header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_continuous_deployment_policy#value CloudfrontContinuousDeploymentPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

