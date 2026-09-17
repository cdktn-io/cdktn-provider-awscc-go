// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationManagedQueryResultsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#enabled AthenaWorkGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the encryption configuration for Athena Managed Storage.
	//
	// If not setting this field, Managed Storage will encrypt the query results with Athena's encryption key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#encryption_configuration AthenaWorkGroup#encryption_configuration}
	EncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationManagedQueryResultsConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

