// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogConfig struct {
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
	// The name of the catalog to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#name GlueCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#allow_full_table_external_data_access GlueCatalog#allow_full_table_external_data_access}
	AllowFullTableExternalDataAccess *string `field:"optional" json:"allowFullTableExternalDataAccess" yaml:"allowFullTableExternalDataAccess"`
	// A structure that specifies data lake access properties and other custom properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#catalog_properties GlueCatalog#catalog_properties}
	CatalogProperties *GlueCatalogCatalogProperties `field:"optional" json:"catalogProperties" yaml:"catalogProperties"`
	// An array of PrincipalPermissions objects for default database permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#create_database_default_permissions GlueCatalog#create_database_default_permissions}
	CreateDatabaseDefaultPermissions interface{} `field:"optional" json:"createDatabaseDefaultPermissions" yaml:"createDatabaseDefaultPermissions"`
	// An array of PrincipalPermissions objects for default table permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#create_table_default_permissions GlueCatalog#create_table_default_permissions}
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#description GlueCatalog#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A FederatedCatalog structure that references an entity outside the Glue Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#federated_catalog GlueCatalog#federated_catalog}
	FederatedCatalog *GlueCatalogFederatedCatalog `field:"optional" json:"federatedCatalog" yaml:"federatedCatalog"`
	// Specifies whether to overwrite child resource permissions with the default permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#overwrite_child_resource_permissions_with_default GlueCatalog#overwrite_child_resource_permissions_with_default}
	OverwriteChildResourcePermissionsWithDefault *string `field:"optional" json:"overwriteChildResourcePermissionsWithDefault" yaml:"overwriteChildResourcePermissionsWithDefault"`
	// A map of key-value pairs that define parameters and properties of the catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#parameters GlueCatalog#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#tags GlueCatalog#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A structure that describes a target catalog for resource linking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_catalog#target_redshift_catalog GlueCatalog#target_redshift_catalog}
	TargetRedshiftCatalog *GlueCatalogTargetRedshiftCatalog `field:"optional" json:"targetRedshiftCatalog" yaml:"targetRedshiftCatalog"`
}

