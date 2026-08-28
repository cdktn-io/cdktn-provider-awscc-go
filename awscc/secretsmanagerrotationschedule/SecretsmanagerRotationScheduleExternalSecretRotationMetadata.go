// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagerrotationschedule


type SecretsmanagerRotationScheduleExternalSecretRotationMetadata struct {
	// The key name of the metadata item. You can specify a value that's 1 to 256 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#key SecretsmanagerRotationSchedule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the metadata item. You can specify a value that's 1 to 2048 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#value SecretsmanagerRotationSchedule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

