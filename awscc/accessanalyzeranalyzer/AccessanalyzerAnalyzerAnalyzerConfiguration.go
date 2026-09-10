// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfiguration struct {
	// Specifies the configuration of an internal access analyzer for an AWS organization or account.
	//
	// This configuration determines how the analyzer evaluates internal access within your AWS environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/accessanalyzer_analyzer#internal_access_configuration AccessanalyzerAnalyzer#internal_access_configuration}
	InternalAccessConfiguration *AccessanalyzerAnalyzerAnalyzerConfigurationInternalAccessConfiguration `field:"optional" json:"internalAccessConfiguration" yaml:"internalAccessConfiguration"`
	// The Configuration for Unused Access Analyzer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/accessanalyzer_analyzer#unused_access_configuration AccessanalyzerAnalyzer#unused_access_configuration}
	UnusedAccessConfiguration *AccessanalyzerAnalyzerAnalyzerConfigurationUnusedAccessConfiguration `field:"optional" json:"unusedAccessConfiguration" yaml:"unusedAccessConfiguration"`
}

