// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableContributorInsightsSpecification struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the CloudWatch Contributor Insights mode for a table.
	//
	// Valid values are ``ACCESSED_AND_THROTTLED_KEYS`` (tracks all access and throttled events) or ``THROTTLED_KEYS`` (tracks only throttled events). This setting determines what type of contributor insights data is collected for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#mode DynamodbTable#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

