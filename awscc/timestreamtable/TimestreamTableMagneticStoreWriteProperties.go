// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable


type TimestreamTableMagneticStoreWriteProperties struct {
	// Boolean flag indicating whether magnetic store writes are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#enable_magnetic_store_writes TimestreamTable#enable_magnetic_store_writes}
	EnableMagneticStoreWrites interface{} `field:"optional" json:"enableMagneticStoreWrites" yaml:"enableMagneticStoreWrites"`
	// Location to store information about records that were asynchronously rejected during magnetic store writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/timestream_table#magnetic_store_rejected_data_location TimestreamTable#magnetic_store_rejected_data_location}
	MagneticStoreRejectedDataLocation *TimestreamTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation `field:"optional" json:"magneticStoreRejectedDataLocation" yaml:"magneticStoreRejectedDataLocation"`
}

