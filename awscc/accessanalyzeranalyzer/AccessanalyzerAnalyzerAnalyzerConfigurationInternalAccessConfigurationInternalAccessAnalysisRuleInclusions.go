// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accessanalyzeranalyzer


type AccessanalyzerAnalyzerAnalyzerConfigurationInternalAccessConfigurationInternalAccessAnalysisRuleInclusions struct {
	// A list of AWS account IDs to apply to the internal access analysis rule criteria.
	//
	// Account IDs can only be applied to the analysis rule criteria for organization-level analyzers and cannot include the organization owner account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/accessanalyzer_analyzer#account_ids AccessanalyzerAnalyzer#account_ids}
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// A list of resource ARNs to apply to the internal access analysis rule criteria.
	//
	// The analyzer will only generate findings for resources that match these ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/accessanalyzer_analyzer#resource_arns AccessanalyzerAnalyzer#resource_arns}
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// A list of resource types to apply to the internal access analysis rule criteria.
	//
	// The analyzer will only generate findings for resources of these types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/accessanalyzer_analyzer#resource_types AccessanalyzerAnalyzer#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

