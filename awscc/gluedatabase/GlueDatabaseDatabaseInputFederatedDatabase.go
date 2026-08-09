// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedatabase


type GlueDatabaseDatabaseInputFederatedDatabase struct {
	// The name of the connection to the external metastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_database#connection_name GlueDatabase#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A unique identifier for the federated database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/glue_database#identifier GlueDatabase#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

