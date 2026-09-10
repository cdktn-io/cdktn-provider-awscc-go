// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsKendraConfiguration struct {
	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#kendra_index LexBot#kendra_index}
	KendraIndex *string `field:"optional" json:"kendraIndex" yaml:"kendraIndex"`
	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#query_filter_string LexBot#query_filter_string}
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Determines whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#query_filter_string_enabled LexBot#query_filter_string_enabled}
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

