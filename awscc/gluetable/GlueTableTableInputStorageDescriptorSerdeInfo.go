// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputStorageDescriptorSerdeInfo struct {
	// Name of the SerDe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#name GlueTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// These key-value pairs define initialization parameters for the SerDe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#parameters GlueTable#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Usually the class that implements the SerDe. An example is org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#serialization_library GlueTable#serialization_library}
	SerializationLibrary *string `field:"optional" json:"serializationLibrary" yaml:"serializationLibrary"`
}

