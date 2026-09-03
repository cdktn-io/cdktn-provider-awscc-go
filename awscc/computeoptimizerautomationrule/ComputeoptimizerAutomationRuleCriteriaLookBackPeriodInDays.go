// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerautomationrule


type ComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDays struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/computeoptimizer_automation_rule#comparison ComputeoptimizerAutomationRule#comparison}.
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/computeoptimizer_automation_rule#values ComputeoptimizerAutomationRule#values}.
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

