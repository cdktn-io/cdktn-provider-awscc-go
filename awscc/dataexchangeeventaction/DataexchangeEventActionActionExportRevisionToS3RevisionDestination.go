// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3RevisionDestination struct {
	// The Amazon S3 bucket that is the destination for the event action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dataexchange_event_action#bucket DataexchangeEventAction#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// A string representing the pattern for generated names of the individual assets in the revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dataexchange_event_action#key_pattern DataexchangeEventAction#key_pattern}
	KeyPattern *string `field:"optional" json:"keyPattern" yaml:"keyPattern"`
}

