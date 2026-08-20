// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrationsPagerDutyConfiguration struct {
	// The name of the pagerDuty configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmincidents_response_plan#name SsmincidentsResponsePlan#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The pagerDuty incident configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmincidents_response_plan#pager_duty_incident_configuration SsmincidentsResponsePlan#pager_duty_incident_configuration}
	PagerDutyIncidentConfiguration *SsmincidentsResponsePlanIntegrationsPagerDutyConfigurationPagerDutyIncidentConfiguration `field:"optional" json:"pagerDutyIncidentConfiguration" yaml:"pagerDutyIncidentConfiguration"`
	// The AWS secrets manager secretId storing the pagerDuty token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmincidents_response_plan#secret_id SsmincidentsResponsePlan#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

