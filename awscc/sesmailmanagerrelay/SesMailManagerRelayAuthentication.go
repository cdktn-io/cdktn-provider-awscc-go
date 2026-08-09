// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagerrelay


type SesMailManagerRelayAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_mail_manager_relay#no_authentication SesMailManagerRelay#no_authentication}.
	NoAuthentication *string `field:"optional" json:"noAuthentication" yaml:"noAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_mail_manager_relay#secret_arn SesMailManagerRelay#secret_arn}.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

