// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MemorydbClusterConfig struct {
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
	// The name of the Access Control List to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#acl_name MemorydbCluster#acl_name}
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
	// The name of the cluster. This value must be unique as it also serves as the cluster identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#cluster_name MemorydbCluster#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The compute and memory capacity of the nodes in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#node_type MemorydbCluster#node_type}
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// A flag that enables automatic minor version upgrade when set to true.
	//
	// You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#auto_minor_version_upgrade MemorydbCluster#auto_minor_version_upgrade}
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The cluster endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#cluster_endpoint MemorydbCluster#cluster_endpoint}
	ClusterEndpoint *MemorydbClusterClusterEndpoint `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// Enables data tiering.
	//
	// Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#data_tiering MemorydbCluster#data_tiering}
	DataTiering *string `field:"optional" json:"dataTiering" yaml:"dataTiering"`
	// An optional description of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#description MemorydbCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The engine type used by the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#engine MemorydbCluster#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The Redis engine version used by the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#engine_version MemorydbCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The user-supplied name of a final cluster snapshot.
	//
	// This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#final_snapshot_name MemorydbCluster#final_snapshot_name}
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// For clusters wth dual stack NetworkType, IpDiscovery controls the Ip protocol (ipv4 or ipv6) returned by the engine commands such as `cluster info` and `cluster nodes` which are used by clients to connect to the nodes in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#ip_discovery MemorydbCluster#ip_discovery}
	IpDiscovery *string `field:"optional" json:"ipDiscovery" yaml:"ipDiscovery"`
	// The ID of the KMS key used to encrypt the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#kms_key_id MemorydbCluster#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	//
	// It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#maintenance_window MemorydbCluster#maintenance_window}
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The name of the Global Datastore, it is generated by MemoryDB adding a prefix to MultiRegionClusterNameSuffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#multi_region_cluster_name MemorydbCluster#multi_region_cluster_name}
	MultiRegionClusterName *string `field:"optional" json:"multiRegionClusterName" yaml:"multiRegionClusterName"`
	// Must be either ipv4 | ipv6 | dual_stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#network_type MemorydbCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The number of replicas to apply to each shard. The limit is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#num_replicas_per_shard MemorydbCluster#num_replicas_per_shard}
	NumReplicasPerShard *float64 `field:"optional" json:"numReplicasPerShard" yaml:"numReplicasPerShard"`
	// The number of shards the cluster will contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#num_shards MemorydbCluster#num_shards}
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// The name of the parameter group associated with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#parameter_group_name MemorydbCluster#parameter_group_name}
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// The port number on which each member of the cluster accepts connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#port MemorydbCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// One or more Amazon VPC security groups associated with this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#security_group_ids MemorydbCluster#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3.
	//
	// The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#snapshot_arns MemorydbCluster#snapshot_arns}
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new cluster.
	//
	// The snapshot status changes to restoring while the new cluster is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#snapshot_name MemorydbCluster#snapshot_name}
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The number of days for which MemoryDB retains automatic snapshots before deleting them.
	//
	// For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#snapshot_retention_limit MemorydbCluster#snapshot_retention_limit}
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#snapshot_window MemorydbCluster#snapshot_window}
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#sns_topic_arn MemorydbCluster#sns_topic_arn}
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#sns_topic_status MemorydbCluster#sns_topic_status}
	SnsTopicStatus *string `field:"optional" json:"snsTopicStatus" yaml:"snsTopicStatus"`
	// The name of the subnet group to be used for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#subnet_group_name MemorydbCluster#subnet_group_name}
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// An array of key-value pairs to apply to this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#tags MemorydbCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#tls_enabled MemorydbCluster#tls_enabled}
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
}

