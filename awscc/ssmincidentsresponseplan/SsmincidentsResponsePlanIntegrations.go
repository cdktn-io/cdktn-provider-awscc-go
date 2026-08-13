// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrations struct {
	// The pagerDuty configuration to use when starting the incident.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ssmincidents_response_plan#pager_duty_configuration SsmincidentsResponsePlan#pager_duty_configuration}
	PagerDutyConfiguration *SsmincidentsResponsePlanIntegrationsPagerDutyConfiguration `field:"optional" json:"pagerDutyConfiguration" yaml:"pagerDutyConfiguration"`
}

