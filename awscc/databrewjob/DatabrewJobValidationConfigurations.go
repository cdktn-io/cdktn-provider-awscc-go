// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob


type DatabrewJobValidationConfigurations struct {
	// Arn of the Ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#ruleset_arn DatabrewJob#ruleset_arn}
	RulesetArn *string `field:"optional" json:"rulesetArn" yaml:"rulesetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/databrew_job#validation_mode DatabrewJob#validation_mode}.
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

