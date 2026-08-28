// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationUnusedAccessConfigurationAnalysisRuleExclusions struct {
	// A list of AWS account IDs to apply to the analysis rule criteria.
	//
	// The accounts cannot include the organization analyzer owner account. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/accessanalyzer_analyzer#account_ids AccessanalyzerAnalyzer#account_ids}
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// An array of key-value pairs to match for your resources.
	//
	// You can use the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// For the tag key, you can specify a value that is 1 to 128 characters in length and cannot be prefixed with aws:.
	//
	// For the tag value, you can specify a value that is 0 to 256 characters in length. If the specified tag value is 0 characters, the rule is applied to all principals with the specified tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/accessanalyzer_analyzer#resource_tags AccessanalyzerAnalyzer#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

