// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationUnusedAccessConfigurationAnalysisRule struct {
	// A list of rules for the analyzer containing criteria to exclude from analysis.
	//
	// Entities that meet the rule criteria will not generate findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/accessanalyzer_analyzer#exclusions AccessanalyzerAnalyzer#exclusions}
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
}

