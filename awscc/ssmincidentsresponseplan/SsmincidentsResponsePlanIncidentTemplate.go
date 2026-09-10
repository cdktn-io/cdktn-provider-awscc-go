// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIncidentTemplate struct {
	// The impact value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#impact SsmincidentsResponsePlan#impact}
	Impact *float64 `field:"required" json:"impact" yaml:"impact"`
	// The title string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#title SsmincidentsResponsePlan#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The deduplication string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#dedupe_string SsmincidentsResponsePlan#dedupe_string}
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Tags that get applied to incidents created by the StartIncident API action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#incident_tags SsmincidentsResponsePlan#incident_tags}
	IncidentTags interface{} `field:"optional" json:"incidentTags" yaml:"incidentTags"`
	// The list of notification targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#notification_targets SsmincidentsResponsePlan#notification_targets}
	NotificationTargets interface{} `field:"optional" json:"notificationTargets" yaml:"notificationTargets"`
	// The summary string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmincidents_response_plan#summary SsmincidentsResponsePlan#summary}
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

