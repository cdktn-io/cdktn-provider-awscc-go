// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofileseventtrigger


type CustomerprofilesEventTriggerIntentFilterDefinitionGroupsConditionsIntent struct {
	// The domain of the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#intent_domain CustomerprofilesEventTrigger#intent_domain}
	IntentDomain *string `field:"optional" json:"intentDomain" yaml:"intentDomain"`
	// The type of the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/customerprofiles_event_trigger#type CustomerprofilesEventTrigger#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

