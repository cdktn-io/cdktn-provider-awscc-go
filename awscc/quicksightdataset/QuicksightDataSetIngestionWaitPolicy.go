// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetIngestionWaitPolicy struct {
	// <p>The maximum time (in hours) to wait for Ingestion to complete.
	//
	// Default timeout is 36 hours.
	//  Applicable only when DataSetImportMode mode is set to SPICE and WaitForSpiceIngestion is set to true.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_data_set#ingestion_wait_time_in_hours QuicksightDataSet#ingestion_wait_time_in_hours}
	IngestionWaitTimeInHours *float64 `field:"optional" json:"ingestionWaitTimeInHours" yaml:"ingestionWaitTimeInHours"`
	// <p>Wait for SPICE ingestion to finish to mark dataset creation/update successful.
	//
	// Default (true).
	//   Applicable only when DataSetImportMode mode is set to SPICE.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_data_set#wait_for_spice_ingestion QuicksightDataSet#wait_for_spice_ingestion}
	WaitForSpiceIngestion interface{} `field:"optional" json:"waitForSpiceIngestion" yaml:"waitForSpiceIngestion"`
}

