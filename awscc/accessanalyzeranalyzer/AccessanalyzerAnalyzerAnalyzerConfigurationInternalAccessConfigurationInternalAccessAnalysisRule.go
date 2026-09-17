// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationInternalAccessConfigurationInternalAccessAnalysisRule struct {
	// A list of rules for the internal access analyzer containing criteria to include in analysis.
	//
	// Only resources that meet the rule criteria will generate findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/accessanalyzer_analyzer#inclusions AccessanalyzerAnalyzer#inclusions}
	Inclusions interface{} `field:"optional" json:"inclusions" yaml:"inclusions"`
}

