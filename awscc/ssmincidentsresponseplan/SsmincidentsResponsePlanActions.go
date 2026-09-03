// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanActions struct {
	// The configuration to use when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmincidents_response_plan#ssm_automation SsmincidentsResponsePlan#ssm_automation}
	SsmAutomation *SsmincidentsResponsePlanActionsSsmAutomation `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

