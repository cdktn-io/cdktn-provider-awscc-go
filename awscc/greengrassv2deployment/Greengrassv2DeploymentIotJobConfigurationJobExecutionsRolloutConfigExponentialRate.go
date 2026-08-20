// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#base_rate_per_minute Greengrassv2Deployment#base_rate_per_minute}.
	BaseRatePerMinute *float64 `field:"optional" json:"baseRatePerMinute" yaml:"baseRatePerMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#increment_factor Greengrassv2Deployment#increment_factor}.
	IncrementFactor *float64 `field:"optional" json:"incrementFactor" yaml:"incrementFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#rate_increase_criteria Greengrassv2Deployment#rate_increase_criteria}.
	RateIncreaseCriteria *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRateRateIncreaseCriteria `field:"optional" json:"rateIncreaseCriteria" yaml:"rateIncreaseCriteria"`
}

