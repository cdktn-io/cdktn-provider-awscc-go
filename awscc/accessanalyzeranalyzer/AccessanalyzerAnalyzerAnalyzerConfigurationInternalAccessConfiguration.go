// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationInternalAccessConfiguration struct {
	// Contains information about analysis rules for the internal access analyzer.
	//
	// Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/accessanalyzer_analyzer#internal_access_analysis_rule AccessanalyzerAnalyzer#internal_access_analysis_rule}
	InternalAccessAnalysisRule *AccessanalyzerAnalyzerAnalyzerConfigurationInternalAccessConfigurationInternalAccessAnalysisRule `field:"optional" json:"internalAccessAnalysisRule" yaml:"internalAccessAnalysisRule"`
}

