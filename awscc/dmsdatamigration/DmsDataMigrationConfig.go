// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdatamigration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsDataMigrationConfig struct {
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
	// The property describes the type of migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#data_migration_type DmsDataMigration#data_migration_type}
	DataMigrationType *string `field:"required" json:"dataMigrationType" yaml:"dataMigrationType"`
	// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#migration_project_identifier DmsDataMigration#migration_project_identifier}
	MigrationProjectIdentifier *string `field:"required" json:"migrationProjectIdentifier" yaml:"migrationProjectIdentifier"`
	// The property describes Amazon Resource Name (ARN) of the service access role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#service_access_role_arn DmsDataMigration#service_access_role_arn}
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// The property describes an ARN of the data migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#data_migration_identifier DmsDataMigration#data_migration_identifier}
	DataMigrationIdentifier *string `field:"optional" json:"dataMigrationIdentifier" yaml:"dataMigrationIdentifier"`
	// The property describes a name to identify the data migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#data_migration_name DmsDataMigration#data_migration_name}
	DataMigrationName *string `field:"optional" json:"dataMigrationName" yaml:"dataMigrationName"`
	// The property describes the settings for the data migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#data_migration_settings DmsDataMigration#data_migration_settings}
	DataMigrationSettings *DmsDataMigrationDataMigrationSettings `field:"optional" json:"dataMigrationSettings" yaml:"dataMigrationSettings"`
	// The property describes the settings for the data migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#source_data_settings DmsDataMigration#source_data_settings}
	SourceDataSettings interface{} `field:"optional" json:"sourceDataSettings" yaml:"sourceDataSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#tags DmsDataMigration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

