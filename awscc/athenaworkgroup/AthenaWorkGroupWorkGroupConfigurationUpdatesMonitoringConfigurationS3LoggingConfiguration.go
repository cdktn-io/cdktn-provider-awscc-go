// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationS3LoggingConfiguration struct {
	// Enables S3 log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#enabled AthenaWorkGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The KMS key ARN to encrypt the logs published to the given Amazon S3 destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#kms_key AthenaWorkGroup#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Amazon S3 destination URI for log publishing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/athena_work_group#log_location AthenaWorkGroup#log_location}
	LogLocation *string `field:"optional" json:"logLocation" yaml:"logLocation"`
}

