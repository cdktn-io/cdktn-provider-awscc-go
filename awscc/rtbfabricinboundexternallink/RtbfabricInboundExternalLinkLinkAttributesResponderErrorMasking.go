// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricinboundexternallink


type RtbfabricInboundExternalLinkLinkAttributesResponderErrorMasking struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_inbound_external_link#action RtbfabricInboundExternalLink#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_inbound_external_link#http_code RtbfabricInboundExternalLink#http_code}.
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_inbound_external_link#logging_types RtbfabricInboundExternalLink#logging_types}.
	LoggingTypes *[]*string `field:"optional" json:"loggingTypes" yaml:"loggingTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_inbound_external_link#response_logging_percentage RtbfabricInboundExternalLink#response_logging_percentage}.
	ResponseLoggingPercentage *float64 `field:"optional" json:"responseLoggingPercentage" yaml:"responseLoggingPercentage"`
}

