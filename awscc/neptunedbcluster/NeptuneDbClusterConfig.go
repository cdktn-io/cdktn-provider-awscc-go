// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunedbcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptuneDbClusterConfig struct {
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
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#associated_roles NeptuneDbCluster#associated_roles}
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#availability_zones NeptuneDbCluster#availability_zones}
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Specifies the number of days for which automatic DB snapshots are retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#backup_retention_period NeptuneDbCluster#backup_retention_period}
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default behaviour is not to copy them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#copy_tags_to_snapshot NeptuneDbCluster#copy_tags_to_snapshot}
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The DB cluster identifier.
	//
	// Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#db_cluster_identifier NeptuneDbCluster#db_cluster_identifier}
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Provides the name of the DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#db_cluster_parameter_group_name NeptuneDbCluster#db_cluster_parameter_group_name}
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	//
	// Used only in case of a major EngineVersion upgrade request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#db_instance_parameter_group_name NeptuneDbCluster#db_instance_parameter_group_name}
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// The port number on which the DB instances in the DB cluster accept connections.
	//
	// If not specified, the default port used is `8182`.
	//
	// Note: `Port` property will soon be deprecated from this resource. Please update existing templates to rename it with new property `DBPort` having same functionalities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#db_port NeptuneDbCluster#db_port}
	DbPort *float64 `field:"optional" json:"dbPort" yaml:"dbPort"`
	// Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#db_subnet_group_name NeptuneDbCluster#db_subnet_group_name}
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Indicates whether or not the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#deletion_protection NeptuneDbCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Specifies a list of log types that are enabled for export to CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#enable_cloudwatch_logs_exports NeptuneDbCluster#enable_cloudwatch_logs_exports}
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Indicates the database engine version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#engine_version NeptuneDbCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The ID of the Neptune global database to which this new DB cluster should be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#global_cluster_identifier NeptuneDbCluster#global_cluster_identifier}
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#iam_auth_enabled NeptuneDbCluster#iam_auth_enabled}
	IamAuthEnabled interface{} `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	// If you enable the StorageEncrypted property but don't specify this property, the default KMS key is used. If you specify this property, you must set the StorageEncrypted property to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#kms_key_id NeptuneDbCluster#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The network type of the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#network_type NeptuneDbCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#preferred_backup_window NeptuneDbCluster#preferred_backup_window}
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#preferred_maintenance_window NeptuneDbCluster#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#restore_to_time NeptuneDbCluster#restore_to_time}
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#restore_type NeptuneDbCluster#restore_type}
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#serverless_scaling_configuration NeptuneDbCluster#serverless_scaling_configuration}
	ServerlessScalingConfiguration *NeptuneDbClusterServerlessScalingConfiguration `field:"optional" json:"serverlessScalingConfiguration" yaml:"serverlessScalingConfiguration"`
	// Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.
	//
	// After you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.
	//
	// However, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#snapshot_identifier NeptuneDbCluster#snapshot_identifier}
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#source_db_cluster_identifier NeptuneDbCluster#source_db_cluster_identifier}
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the KmsKeyId property, then you must enable encryption and set this property to true.
	//
	// If you enable the StorageEncrypted property but don't specify KmsKeyId property, then the default KMS key is used. If you specify KmsKeyId property, then that KMS Key is used to encrypt the database instances in the DB cluster.
	//
	// If you specify the SourceDBClusterIdentifier property and don't specify this property or disable it. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the KmsKeyId property from the source cluster is used.
	//
	// If you specify the DBSnapshotIdentifier and don't specify this property or disable it. The value is inherited from the snapshot, and the specified KmsKeyId property from the snapshot is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#storage_encrypted NeptuneDbCluster#storage_encrypted}
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The tags assigned to this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#tags NeptuneDbCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#use_latest_restorable_time NeptuneDbCluster#use_latest_restorable_time}
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// Provides a list of VPC security groups that the DB cluster belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/neptune_db_cluster#vpc_security_group_ids NeptuneDbCluster#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

