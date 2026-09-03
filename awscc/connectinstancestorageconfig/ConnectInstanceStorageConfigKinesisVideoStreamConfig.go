// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigKinesisVideoStreamConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_instance_storage_config#encryption_config ConnectInstanceStorageConfig#encryption_config}.
	EncryptionConfig *ConnectInstanceStorageConfigKinesisVideoStreamConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Prefixes are used to infer logical hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_instance_storage_config#prefix ConnectInstanceStorageConfig#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Number of hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_instance_storage_config#retention_period_hours ConnectInstanceStorageConfig#retention_period_hours}
	RetentionPeriodHours *float64 `field:"optional" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
}

