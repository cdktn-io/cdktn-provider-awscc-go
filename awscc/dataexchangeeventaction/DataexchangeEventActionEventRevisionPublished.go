// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataexchangeeventaction


type DataexchangeEventActionEventRevisionPublished struct {
	// The data set ID of the published revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dataexchange_event_action#data_set_id DataexchangeEventAction#data_set_id}
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
}

