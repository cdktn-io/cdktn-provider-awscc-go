// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionAction struct {
	// Details of the operation to be performed by the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dataexchange_event_action#export_revision_to_s3 DataexchangeEventAction#export_revision_to_s3}
	ExportRevisionToS3 *DataexchangeEventActionActionExportRevisionToS3 `field:"optional" json:"exportRevisionToS3" yaml:"exportRevisionToS3"`
}

