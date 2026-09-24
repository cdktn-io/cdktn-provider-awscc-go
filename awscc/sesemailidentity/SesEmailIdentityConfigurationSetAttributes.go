// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesemailidentity


type SesEmailIdentityConfigurationSetAttributes struct {
	// The configuration set to use by default when sending from this identity.
	//
	// Note that any configuration set defined in the email sending request takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ses_email_identity#configuration_set_name SesEmailIdentity#configuration_set_name}
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
}

