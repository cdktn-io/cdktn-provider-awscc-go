// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppfabricIngestionDestinationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) of the app bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#app_bundle_arn AppfabricIngestionDestination#app_bundle_arn}
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// Contains information about the destination of ingested data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#destination_configuration AppfabricIngestionDestination#destination_configuration}
	DestinationConfiguration *AppfabricIngestionDestinationDestinationConfiguration `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// The Amazon Resource Name (ARN) of the ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#ingestion_arn AppfabricIngestionDestination#ingestion_arn}
	IngestionArn *string `field:"required" json:"ingestionArn" yaml:"ingestionArn"`
	// Contains information about how ingested data is processed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#processing_configuration AppfabricIngestionDestination#processing_configuration}
	ProcessingConfiguration *AppfabricIngestionDestinationProcessingConfiguration `field:"required" json:"processingConfiguration" yaml:"processingConfiguration"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/appfabric_ingestion_destination#tags AppfabricIngestionDestination#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

