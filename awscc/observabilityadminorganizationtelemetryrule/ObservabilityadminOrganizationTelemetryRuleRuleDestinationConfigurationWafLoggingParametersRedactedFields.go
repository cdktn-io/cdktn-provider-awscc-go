// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule


type ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationWafLoggingParametersRedactedFields struct {
	// The method with which to match this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#method ObservabilityadminOrganizationTelemetryRule#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The query string to find the resource to match this field to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#query_string ObservabilityadminOrganizationTelemetryRule#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Header for the field to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#single_header ObservabilityadminOrganizationTelemetryRule#single_header}
	SingleHeader *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationWafLoggingParametersRedactedFieldsSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// This is the URI path to match this rule to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/observabilityadmin_organization_telemetry_rule#uri_path ObservabilityadminOrganizationTelemetryRule#uri_path}
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

