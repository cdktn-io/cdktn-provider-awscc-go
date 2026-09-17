// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmincidentsResponsePlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The incident template configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#incident_template SsmincidentsResponsePlan#incident_template}
	IncidentTemplate *SsmincidentsResponsePlanIncidentTemplate `field:"required" json:"incidentTemplate" yaml:"incidentTemplate"`
	// The name of the response plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#name SsmincidentsResponsePlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#actions SsmincidentsResponsePlan#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The chat channel configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#chat_channel SsmincidentsResponsePlan#chat_channel}
	ChatChannel *SsmincidentsResponsePlanChatChannel `field:"optional" json:"chatChannel" yaml:"chatChannel"`
	// The display name of the response plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#display_name SsmincidentsResponsePlan#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The list of engagements to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#engagements SsmincidentsResponsePlan#engagements}
	Engagements *[]*string `field:"optional" json:"engagements" yaml:"engagements"`
	// The list of integrations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#integrations SsmincidentsResponsePlan#integrations}
	Integrations interface{} `field:"optional" json:"integrations" yaml:"integrations"`
	// The tags to apply to the response plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmincidents_response_plan#tags SsmincidentsResponsePlan#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

