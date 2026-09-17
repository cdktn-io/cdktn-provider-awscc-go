// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluedatabase


type GlueDatabaseDatabaseInput struct {
	// Creates a set of default permissions on the table for principals.
	//
	// Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#create_table_default_permissions GlueDatabase#create_table_default_permissions}
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#description GlueDatabase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A FederatedDatabase structure that references an entity outside the AWS Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#federated_database GlueDatabase#federated_database}
	FederatedDatabase *GlueDatabaseDatabaseInputFederatedDatabase `field:"optional" json:"federatedDatabase" yaml:"federatedDatabase"`
	// The location of the database (for example, an HDFS path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#location_uri GlueDatabase#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The name of the database. For hive compatibility, this is folded to lowercase when it is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#name GlueDatabase#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// These key-value pairs define parameters and properties of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#parameters GlueDatabase#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// A DatabaseIdentifier structure that describes a target database for resource linking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_database#target_database GlueDatabase#target_database}
	TargetDatabase *GlueDatabaseDatabaseInputTargetDatabase `field:"optional" json:"targetDatabase" yaml:"targetDatabase"`
}

