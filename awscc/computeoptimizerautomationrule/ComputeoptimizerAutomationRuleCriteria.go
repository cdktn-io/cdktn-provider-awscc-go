// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeoptimizerautomationrule


type ComputeoptimizerAutomationRuleCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#ebs_volume_size_in_gib ComputeoptimizerAutomationRule#ebs_volume_size_in_gib}.
	EbsVolumeSizeInGib interface{} `field:"optional" json:"ebsVolumeSizeInGib" yaml:"ebsVolumeSizeInGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#ebs_volume_type ComputeoptimizerAutomationRule#ebs_volume_type}.
	EbsVolumeType interface{} `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#estimated_monthly_savings ComputeoptimizerAutomationRule#estimated_monthly_savings}.
	EstimatedMonthlySavings interface{} `field:"optional" json:"estimatedMonthlySavings" yaml:"estimatedMonthlySavings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#look_back_period_in_days ComputeoptimizerAutomationRule#look_back_period_in_days}.
	LookBackPeriodInDays interface{} `field:"optional" json:"lookBackPeriodInDays" yaml:"lookBackPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#region ComputeoptimizerAutomationRule#region}.
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#resource_arn ComputeoptimizerAutomationRule#resource_arn}.
	ResourceArn interface{} `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#resource_tag ComputeoptimizerAutomationRule#resource_tag}.
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/computeoptimizer_automation_rule#restart_needed ComputeoptimizerAutomationRule#restart_needed}.
	RestartNeeded interface{} `field:"optional" json:"restartNeeded" yaml:"restartNeeded"`
}

