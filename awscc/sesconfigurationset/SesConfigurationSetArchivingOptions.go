// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationset


type SesConfigurationSetArchivingOptions struct {
	// The ARN of the MailManager archive to associate with the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_configuration_set#archive_arn SesConfigurationSet#archive_arn}
	ArchiveArn *string `field:"optional" json:"archiveArn" yaml:"archiveArn"`
}

