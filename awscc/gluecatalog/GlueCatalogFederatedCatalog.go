// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog


type GlueCatalogFederatedCatalog struct {
	// The name of the connection to an external data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_catalog#connection_name GlueCatalog#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A unique identifier for the federated catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_catalog#identifier GlueCatalog#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

