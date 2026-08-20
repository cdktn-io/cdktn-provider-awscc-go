// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingDocumentDbEventSourceConfig struct {
	// The name of the collection to consume within the database.
	//
	// If you do not specify a collection, Lambda consumes all collections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#collection_name LambdaEventSourceMapping#collection_name}
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// The name of the database to consume within the DocumentDB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#database_name LambdaEventSourceMapping#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Determines what DocumentDB sends to your event stream during document update operations.
	//
	// If set to UpdateLookup, DocumentDB sends a delta describing the changes, along with a copy of the entire document. Otherwise, DocumentDB sends only a partial document that contains the changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lambda_event_source_mapping#full_document LambdaEventSourceMapping#full_document}
	FullDocument *string `field:"optional" json:"fullDocument" yaml:"fullDocument"`
}

