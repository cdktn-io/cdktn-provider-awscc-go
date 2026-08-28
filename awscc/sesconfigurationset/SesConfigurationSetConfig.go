// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesConfigurationSetConfig struct {
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
	// An object that defines a MailManager archive that is used to preserve emails that you send using the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#archiving_options SesConfigurationSet#archiving_options}
	ArchivingOptions *SesConfigurationSetArchivingOptions `field:"optional" json:"archivingOptions" yaml:"archivingOptions"`
	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#delivery_options SesConfigurationSet#delivery_options}
	DeliveryOptions *SesConfigurationSetDeliveryOptions `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// The name of the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#name SesConfigurationSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#reputation_options SesConfigurationSet#reputation_options}
	ReputationOptions *SesConfigurationSetReputationOptions `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// An object that defines whether or not Amazon SES can send email that you send using the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#sending_options SesConfigurationSet#sending_options}
	SendingOptions *SesConfigurationSetSendingOptions `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// An object that contains information about the suppression list preferences for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#suppression_options SesConfigurationSet#suppression_options}
	SuppressionOptions *SesConfigurationSetSuppressionOptions `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// The tags (keys and values) associated with the contact list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#tags SesConfigurationSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// An object that defines the open and click tracking options for emails that you send using the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#tracking_options SesConfigurationSet#tracking_options}
	TrackingOptions *SesConfigurationSetTrackingOptions `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// An object that contains Virtual Deliverability Manager (VDM) settings for this configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ses_configuration_set#vdm_options SesConfigurationSet#vdm_options}
	VdmOptions *SesConfigurationSetVdmOptions `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}

