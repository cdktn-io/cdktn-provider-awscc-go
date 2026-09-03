// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionsSsmAutomationDynamicParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmincidents_response_plan#key SsmincidentsResponsePlan#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value of the dynamic parameter to set when starting the SSM automation document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmincidents_response_plan#value SsmincidentsResponsePlan#value}
	Value *SsmincidentsResponsePlanActionsSsmAutomationDynamicParametersValue `field:"optional" json:"value" yaml:"value"`
}

