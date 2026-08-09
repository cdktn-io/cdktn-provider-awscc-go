// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsmigrationproject

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsMigrationProjectConfig struct {
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
	// The optional description of the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#description DmsMigrationProject#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The property describes an instance profile arn for the migration project. For read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#instance_profile_arn DmsMigrationProject#instance_profile_arn}
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The property describes an instance profile identifier for the migration project. For create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#instance_profile_identifier DmsMigrationProject#instance_profile_identifier}
	InstanceProfileIdentifier *string `field:"optional" json:"instanceProfileIdentifier" yaml:"instanceProfileIdentifier"`
	// The property describes an instance profile name for the migration project. For read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#instance_profile_name DmsMigrationProject#instance_profile_name}
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The property describes a creating time of the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#migration_project_creation_time DmsMigrationProject#migration_project_creation_time}
	MigrationProjectCreationTime *string `field:"optional" json:"migrationProjectCreationTime" yaml:"migrationProjectCreationTime"`
	// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#migration_project_identifier DmsMigrationProject#migration_project_identifier}
	MigrationProjectIdentifier *string `field:"optional" json:"migrationProjectIdentifier" yaml:"migrationProjectIdentifier"`
	// The property describes a name to identify the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#migration_project_name DmsMigrationProject#migration_project_name}
	MigrationProjectName *string `field:"optional" json:"migrationProjectName" yaml:"migrationProjectName"`
	// The property describes schema conversion application attributes for the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#schema_conversion_application_attributes DmsMigrationProject#schema_conversion_application_attributes}
	SchemaConversionApplicationAttributes *DmsMigrationProjectSchemaConversionApplicationAttributes `field:"optional" json:"schemaConversionApplicationAttributes" yaml:"schemaConversionApplicationAttributes"`
	// The property describes source data provider descriptors for the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#source_data_provider_descriptors DmsMigrationProject#source_data_provider_descriptors}
	SourceDataProviderDescriptors interface{} `field:"optional" json:"sourceDataProviderDescriptors" yaml:"sourceDataProviderDescriptors"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#tags DmsMigrationProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The property describes target data provider descriptors for the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#target_data_provider_descriptors DmsMigrationProject#target_data_provider_descriptors}
	TargetDataProviderDescriptors interface{} `field:"optional" json:"targetDataProviderDescriptors" yaml:"targetDataProviderDescriptors"`
	// The property describes transformation rules for the migration project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_migration_project#transformation_rules DmsMigrationProject#transformation_rules}
	TransformationRules *string `field:"optional" json:"transformationRules" yaml:"transformationRules"`
}

