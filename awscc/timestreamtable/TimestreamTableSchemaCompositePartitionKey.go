// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable


type TimestreamTableSchemaCompositePartitionKey struct {
	// The level of enforcement for the specification of a dimension key in ingested records.
	//
	// Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_table#enforcement_in_record TimestreamTable#enforcement_in_record}
	EnforcementInRecord *string `field:"optional" json:"enforcementInRecord" yaml:"enforcementInRecord"`
	// The name of the attribute used for a dimension key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_table#name TimestreamTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/timestream_table#type TimestreamTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

