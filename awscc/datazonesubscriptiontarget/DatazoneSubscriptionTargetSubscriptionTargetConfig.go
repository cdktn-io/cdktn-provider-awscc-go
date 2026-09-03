// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonesubscriptiontarget


type DatazoneSubscriptionTargetSubscriptionTargetConfig struct {
	// The content of the subscription target configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_subscription_target#content DatazoneSubscriptionTarget#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The form name included in the subscription target configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_subscription_target#form_name DatazoneSubscriptionTarget#form_name}
	FormName *string `field:"required" json:"formName" yaml:"formName"`
}

