// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationTags struct {
	// Tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#key AppfabricIngestionDestination#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_ingestion_destination#value AppfabricIngestionDestination#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

