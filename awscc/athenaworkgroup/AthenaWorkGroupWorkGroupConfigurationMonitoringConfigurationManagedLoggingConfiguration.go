// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfiguration struct {
	// Enables managed log persistence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#enabled AthenaWorkGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ARN to encrypt the logs stored in managed log persistence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#kms_key AthenaWorkGroup#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

