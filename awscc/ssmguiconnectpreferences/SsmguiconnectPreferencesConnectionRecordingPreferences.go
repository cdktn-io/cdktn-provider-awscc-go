// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmguiconnectpreferences


type SsmguiconnectPreferencesConnectionRecordingPreferences struct {
	// The ARN of a AWS KMS key that is used to encrypt data while it is being processed by the service.
	//
	// This key must exist in the same AWS Region as the node you start an RDP connection to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmguiconnect_preferences#kms_key_arn SsmguiconnectPreferences#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Determines where recordings of RDP connections are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssmguiconnect_preferences#recording_destinations SsmguiconnectPreferences#recording_destinations}
	RecordingDestinations *SsmguiconnectPreferencesConnectionRecordingPreferencesRecordingDestinations `field:"optional" json:"recordingDestinations" yaml:"recordingDestinations"`
}

