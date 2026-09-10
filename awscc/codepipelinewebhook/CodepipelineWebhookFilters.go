// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinewebhook


type CodepipelineWebhookFilters struct {
	// A JsonPath expression that is applied to the body/payload of the webhook.
	//
	// The value selected by the JsonPath expression must match the value specified in the MatchEquals field. Otherwise, the request is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_webhook#json_path CodepipelineWebhook#json_path}
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The value selected by the JsonPath expression must match what is supplied in the MatchEquals field.
	//
	// Otherwise, the request is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/codepipeline_webhook#match_equals CodepipelineWebhook#match_equals}
	MatchEquals *string `field:"optional" json:"matchEquals" yaml:"matchEquals"`
}

