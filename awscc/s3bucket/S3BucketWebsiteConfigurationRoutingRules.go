// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsiteConfigurationRoutingRules struct {
	// Container for redirect information.
	//
	// You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#redirect_rule S3Bucket#redirect_rule}
	RedirectRule *S3BucketWebsiteConfigurationRoutingRulesRedirectRule `field:"optional" json:"redirectRule" yaml:"redirectRule"`
	// A container for describing a condition that must be met for the specified redirect to apply.
	//
	// For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#routing_rule_condition S3Bucket#routing_rule_condition}
	RoutingRuleCondition *S3BucketWebsiteConfigurationRoutingRulesRoutingRuleCondition `field:"optional" json:"routingRuleCondition" yaml:"routingRuleCondition"`
}

