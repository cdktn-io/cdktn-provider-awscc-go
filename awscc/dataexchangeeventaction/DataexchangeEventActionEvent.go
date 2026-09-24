// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionEvent struct {
	// Information about the published revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dataexchange_event_action#revision_published DataexchangeEventAction#revision_published}
	RevisionPublished *DataexchangeEventActionEventRevisionPublished `field:"optional" json:"revisionPublished" yaml:"revisionPublished"`
}

