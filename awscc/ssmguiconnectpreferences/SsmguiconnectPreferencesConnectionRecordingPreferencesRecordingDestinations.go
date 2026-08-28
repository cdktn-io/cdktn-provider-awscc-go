// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmguiconnectpreferences


type SsmguiconnectPreferencesConnectionRecordingPreferencesRecordingDestinations struct {
	// The S3 bucket where RDP connection recordings are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssmguiconnect_preferences#s3_buckets SsmguiconnectPreferences#s3_buckets}
	S3Buckets interface{} `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
}

