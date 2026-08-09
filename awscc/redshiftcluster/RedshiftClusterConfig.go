// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftClusterConfig struct {
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
	// The type of the cluster.
	//
	// When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_type RedshiftCluster#cluster_type}
	ClusterType *string `field:"required" json:"clusterType" yaml:"clusterType"`
	// The name of the first database to be created when the cluster is created.
	//
	// To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#db_name RedshiftCluster#db_name}
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// The user name associated with the master user account for the cluster that is being created.
	//
	// The user name can't be PUBLIC and first character must be a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#master_username RedshiftCluster#master_username}
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.large | ra3.4xlarge | ra3.16xlarge | rg.large | rg.xlarge | rg.4xlarge | rg.12xlarge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#node_type RedshiftCluster#node_type}
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster.
	//
	// Default value is True
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#allow_version_upgrade RedshiftCluster#allow_version_upgrade}
	AllowVersionUpgrade interface{} `field:"optional" json:"allowVersionUpgrade" yaml:"allowVersionUpgrade"`
	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored.
	//
	// Possible values include the following.
	//
	// enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
	// disabled - Don't use AQUA.
	// auto - Amazon Redshift determines whether to use AQUA.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#aqua_configuration_status RedshiftCluster#aqua_configuration_status}
	AquaConfigurationStatus *string `field:"optional" json:"aquaConfigurationStatus" yaml:"aquaConfigurationStatus"`
	// The number of days that automated snapshots are retained.
	//
	// If the value is 0, automated snapshots are disabled. Default value is 1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#automated_snapshot_retention_period RedshiftCluster#automated_snapshot_retention_period}
	AutomatedSnapshotRetentionPeriod *float64 `field:"optional" json:"automatedSnapshotRetentionPeriod" yaml:"automatedSnapshotRetentionPeriod"`
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster.
	//
	// Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#availability_zone RedshiftCluster#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#availability_zone_relocation RedshiftCluster#availability_zone_relocation}
	AvailabilityZoneRelocation interface{} `field:"optional" json:"availabilityZoneRelocation" yaml:"availabilityZoneRelocation"`
	// The availability zone relocation status of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#availability_zone_relocation_status RedshiftCluster#availability_zone_relocation_status}
	AvailabilityZoneRelocationStatus *string `field:"optional" json:"availabilityZoneRelocationStatus" yaml:"availabilityZoneRelocationStatus"`
	// A boolean value indicating whether the resize operation is using the classic resize process.
	//
	// If you don't provide this parameter or set the value to false , the resize type is elastic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#classic RedshiftCluster#classic}
	Classic interface{} `field:"optional" json:"classic" yaml:"classic"`
	// A unique identifier for the cluster.
	//
	// You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_identifier RedshiftCluster#cluster_identifier}
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the parameter group to be associated with this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_parameter_group_name RedshiftCluster#cluster_parameter_group_name}
	ClusterParameterGroupName *string `field:"optional" json:"clusterParameterGroupName" yaml:"clusterParameterGroupName"`
	// A list of security groups to be associated with this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_security_groups RedshiftCluster#cluster_security_groups}
	ClusterSecurityGroups *[]*string `field:"optional" json:"clusterSecurityGroups" yaml:"clusterSecurityGroups"`
	// The name of a cluster subnet group to be associated with this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_subnet_group_name RedshiftCluster#cluster_subnet_group_name}
	ClusterSubnetGroupName *string `field:"optional" json:"clusterSubnetGroupName" yaml:"clusterSubnetGroupName"`
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#cluster_version RedshiftCluster#cluster_version}
	ClusterVersion *string `field:"optional" json:"clusterVersion" yaml:"clusterVersion"`
	// A boolean indicating whether to enable the deferred maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#defer_maintenance RedshiftCluster#defer_maintenance}
	DeferMaintenance interface{} `field:"optional" json:"deferMaintenance" yaml:"deferMaintenance"`
	// An integer indicating the duration of the maintenance window in days.
	//
	// If you specify a duration, you can't specify an end time. The duration must be 60 days or less.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#defer_maintenance_duration RedshiftCluster#defer_maintenance_duration}
	DeferMaintenanceDuration *float64 `field:"optional" json:"deferMaintenanceDuration" yaml:"deferMaintenanceDuration"`
	// A timestamp indicating end time for the deferred maintenance window.
	//
	// If you specify an end time, you can't specify a duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#defer_maintenance_end_time RedshiftCluster#defer_maintenance_end_time}
	DeferMaintenanceEndTime *string `field:"optional" json:"deferMaintenanceEndTime" yaml:"deferMaintenanceEndTime"`
	// A timestamp indicating the start time for the deferred maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#defer_maintenance_start_time RedshiftCluster#defer_maintenance_start_time}
	DeferMaintenanceStartTime *string `field:"optional" json:"deferMaintenanceStartTime" yaml:"deferMaintenanceStartTime"`
	// The destination AWS Region that you want to copy snapshots to.
	//
	// Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#destination_region RedshiftCluster#destination_region}
	DestinationRegion *string `field:"optional" json:"destinationRegion" yaml:"destinationRegion"`
	// The Elastic IP (EIP) address for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#elastic_ip RedshiftCluster#elastic_ip}
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// If true, the data in the cluster is encrypted at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#encrypted RedshiftCluster#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#endpoint RedshiftCluster#endpoint}.
	Endpoint *RedshiftClusterEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// An option that specifies whether to create the cluster with enhanced VPC routing enabled.
	//
	// To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.
	//
	// If this option is true , enhanced VPC routing is enabled.
	//
	// Default: false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#enhanced_vpc_routing RedshiftCluster#enhanced_vpc_routing}
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#hsm_client_certificate_identifier RedshiftCluster#hsm_client_certificate_identifier}
	HsmClientCertificateIdentifier *string `field:"optional" json:"hsmClientCertificateIdentifier" yaml:"hsmClientCertificateIdentifier"`
	// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#hsm_configuration_identifier RedshiftCluster#hsm_configuration_identifier}
	HsmConfigurationIdentifier *string `field:"optional" json:"hsmConfigurationIdentifier" yaml:"hsmConfigurationIdentifier"`
	// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services.
	//
	// You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#iam_roles RedshiftCluster#iam_roles}
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#kms_key_id RedshiftCluster#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#logging_properties RedshiftCluster#logging_properties}.
	LoggingProperties *RedshiftClusterLoggingProperties `field:"optional" json:"loggingProperties" yaml:"loggingProperties"`
	// The name for the maintenance track that you want to assign for the cluster.
	//
	// This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#maintenance_track_name RedshiftCluster#maintenance_track_name}
	MaintenanceTrackName *string `field:"optional" json:"maintenanceTrackName" yaml:"maintenanceTrackName"`
	// A boolean indicating if the redshift cluster's admin user credentials is managed by Redshift or not.
	//
	// You can't use MasterUserPassword if ManageMasterPassword is true. If ManageMasterPassword is false or not set, Amazon Redshift uses MasterUserPassword for the admin user account's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#manage_master_password RedshiftCluster#manage_master_password}
	ManageMasterPassword interface{} `field:"optional" json:"manageMasterPassword" yaml:"manageMasterPassword"`
	// The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region.
	//
	// If the value is -1, the manual snapshot is retained indefinitely.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#manual_snapshot_retention_period RedshiftCluster#manual_snapshot_retention_period}
	ManualSnapshotRetentionPeriod *float64 `field:"optional" json:"manualSnapshotRetentionPeriod" yaml:"manualSnapshotRetentionPeriod"`
	// The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#master_password_secret_kms_key_id RedshiftCluster#master_password_secret_kms_key_id}
	MasterPasswordSecretKmsKeyId *string `field:"optional" json:"masterPasswordSecretKmsKeyId" yaml:"masterPasswordSecretKmsKeyId"`
	// The password associated with the master user account for the cluster that is being created.
	//
	// You can't use MasterUserPassword if ManageMasterPassword is true. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#master_user_password RedshiftCluster#master_user_password}
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// A boolean indicating if the redshift cluster is multi-az or not.
	//
	// If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#multi_az RedshiftCluster#multi_az}
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The namespace resource policy document that will be attached to a Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#namespace_resource_policy RedshiftCluster#namespace_resource_policy}
	NamespaceResourcePolicy *string `field:"optional" json:"namespaceResourcePolicy" yaml:"namespaceResourcePolicy"`
	// The number of compute nodes in the cluster.
	//
	// This parameter is required when the ClusterType parameter is specified as multi-node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#number_of_nodes RedshiftCluster#number_of_nodes}
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#owner_account RedshiftCluster#owner_account}.
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// The port number on which the cluster accepts incoming connections.
	//
	// The cluster is accessible only via the JDBC and ODBC connection strings
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#port RedshiftCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#preferred_maintenance_window RedshiftCluster#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// If true, the cluster can be accessed from a public network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#publicly_accessible RedshiftCluster#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#resource_action RedshiftCluster#resource_action}
	ResourceAction *string `field:"optional" json:"resourceAction" yaml:"resourceAction"`
	// The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#revision_target RedshiftCluster#revision_target}
	RevisionTarget *string `field:"optional" json:"revisionTarget" yaml:"revisionTarget"`
	// A boolean indicating if we want to rotate Encryption Keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#rotate_encryption_key RedshiftCluster#rotate_encryption_key}
	RotateEncryptionKey interface{} `field:"optional" json:"rotateEncryptionKey" yaml:"rotateEncryptionKey"`
	// The name of the cluster the source snapshot was created from.
	//
	// This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#snapshot_cluster_identifier RedshiftCluster#snapshot_cluster_identifier}
	SnapshotClusterIdentifier *string `field:"optional" json:"snapshotClusterIdentifier" yaml:"snapshotClusterIdentifier"`
	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#snapshot_copy_grant_name RedshiftCluster#snapshot_copy_grant_name}
	SnapshotCopyGrantName *string `field:"optional" json:"snapshotCopyGrantName" yaml:"snapshotCopyGrantName"`
	// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#snapshot_copy_manual RedshiftCluster#snapshot_copy_manual}
	SnapshotCopyManual interface{} `field:"optional" json:"snapshotCopyManual" yaml:"snapshotCopyManual"`
	// The number of days to retain automated snapshots in the destination region after they are copied from the source region.
	//
	// Default is 7.
	//
	//  Constraints: Must be at least 1 and no more than 35.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#snapshot_copy_retention_period RedshiftCluster#snapshot_copy_retention_period}
	SnapshotCopyRetentionPeriod *float64 `field:"optional" json:"snapshotCopyRetentionPeriod" yaml:"snapshotCopyRetentionPeriod"`
	// The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#snapshot_identifier RedshiftCluster#snapshot_identifier}
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The list of tags for the cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#tags RedshiftCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/redshift_cluster#vpc_security_group_ids RedshiftCluster#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

