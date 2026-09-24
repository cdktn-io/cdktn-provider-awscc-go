// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmguiconnectpreferences


type SsmguiconnectPreferencesConnectionRecordingPreferencesRecordingDestinationsS3Buckets struct {
	// The name of the S3 bucket where RDP connection recordings are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmguiconnect_preferences#bucket_name SsmguiconnectPreferences#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The AWS account number that owns the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ssmguiconnect_preferences#bucket_owner SsmguiconnectPreferences#bucket_owner}
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
}

