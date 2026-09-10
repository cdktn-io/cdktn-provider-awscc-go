// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildatabase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailDatabaseConfig struct {
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
	// The name of the database to create when the Lightsail database resource is created.
	//
	// For MySQL, if this parameter isn't specified, no database is created in the database resource. For PostgreSQL, if this parameter isn't specified, a database named postgres is created in the database resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#master_database_name LightsailDatabase#master_database_name}
	MasterDatabaseName *string `field:"required" json:"masterDatabaseName" yaml:"masterDatabaseName"`
	// The name for the master user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#master_username LightsailDatabase#master_username}
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// The blueprint ID for your new database. A blueprint describes the major engine version of a database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#relational_database_blueprint_id LightsailDatabase#relational_database_blueprint_id}
	RelationalDatabaseBlueprintId *string `field:"required" json:"relationalDatabaseBlueprintId" yaml:"relationalDatabaseBlueprintId"`
	// The bundle ID for your new database. A bundle describes the performance specifications for your database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#relational_database_bundle_id LightsailDatabase#relational_database_bundle_id}
	RelationalDatabaseBundleId *string `field:"required" json:"relationalDatabaseBundleId" yaml:"relationalDatabaseBundleId"`
	// The name to use for your new Lightsail database resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#relational_database_name LightsailDatabase#relational_database_name}
	RelationalDatabaseName *string `field:"required" json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// The Availability Zone in which to create your new database. Use the us-east-2a case-sensitive format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#availability_zone LightsailDatabase#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// When true, enables automated backup retention for your database.
	//
	// Updates are applied during the next maintenance window because this can result in an outage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#backup_retention LightsailDatabase#backup_retention}
	BackupRetention interface{} `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// Indicates the certificate that needs to be associated with the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#ca_certificate_identifier LightsailDatabase#ca_certificate_identifier}
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// The password for the master user.
	//
	// The password can include any printable ASCII character except "/", """, or "@". It cannot contain spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#master_user_password LightsailDatabase#master_user_password}
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The daily time range during which automated backups are created for your new database if automated backups are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#preferred_backup_window LightsailDatabase#preferred_backup_window}
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur on your new database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#preferred_maintenance_window LightsailDatabase#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for your new database.
	//
	// A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#publicly_accessible LightsailDatabase#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Update one or more parameters of the relational database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#relational_database_parameters LightsailDatabase#relational_database_parameters}
	RelationalDatabaseParameters interface{} `field:"optional" json:"relationalDatabaseParameters" yaml:"relationalDatabaseParameters"`
	// When true, the master user password is changed to a new strong password generated by Lightsail.
	//
	// Use the get relational database master user password operation to get the new password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#rotate_master_user_password LightsailDatabase#rotate_master_user_password}
	RotateMasterUserPassword interface{} `field:"optional" json:"rotateMasterUserPassword" yaml:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lightsail_database#tags LightsailDatabase#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

