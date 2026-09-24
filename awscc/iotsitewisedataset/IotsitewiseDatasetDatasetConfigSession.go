// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset


type IotsitewiseDatasetDatasetConfigSession struct {
	// The end time of the session as an ISO 8601 UTC instant, for example 2024-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#session_end_time IotsitewiseDataset#session_end_time}
	SessionEndTime *string `field:"optional" json:"sessionEndTime" yaml:"sessionEndTime"`
	// The start time of the session as an ISO 8601 UTC instant, for example 2024-01-01T00:00:00Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#session_start_time IotsitewiseDataset#session_start_time}
	SessionStartTime *string `field:"optional" json:"sessionStartTime" yaml:"sessionStartTime"`
}

