// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionActionExportRevisionToS3 struct {
	// Encryption configuration of the export job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dataexchange_event_action#encryption DataexchangeEventAction#encryption}
	Encryption *DataexchangeEventActionActionExportRevisionToS3Encryption `field:"optional" json:"encryption" yaml:"encryption"`
	// A revision destination is the Amazon S3 bucket folder destination to where the export will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dataexchange_event_action#revision_destination DataexchangeEventAction#revision_destination}
	RevisionDestination *DataexchangeEventActionActionExportRevisionToS3RevisionDestination `field:"optional" json:"revisionDestination" yaml:"revisionDestination"`
}

