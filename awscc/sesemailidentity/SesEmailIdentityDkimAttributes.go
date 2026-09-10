// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesemailidentity


type SesEmailIdentityDkimAttributes struct {
	// Sets the DKIM signing configuration for the identity.
	//
	// When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_email_identity#signing_enabled SesEmailIdentity#signing_enabled}
	SigningEnabled interface{} `field:"optional" json:"signingEnabled" yaml:"signingEnabled"`
}

