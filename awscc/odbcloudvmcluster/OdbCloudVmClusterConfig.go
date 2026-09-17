// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbCloudVmClusterConfig struct {
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
	// The unique identifier of the Exadata infrastructure that this VM cluster belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#cloud_exadata_infrastructure_id OdbCloudVmCluster#cloud_exadata_infrastructure_id}
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The name of the Grid Infrastructure (GI) cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#cluster_name OdbCloudVmCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The number of CPU cores enabled on the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#cpu_core_count OdbCloudVmCluster#cpu_core_count}
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The set of diagnostic collection options enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#data_collection_options OdbCloudVmCluster#data_collection_options}
	DataCollectionOptions *OdbCloudVmClusterDataCollectionOptions `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#data_storage_size_in_t_bs OdbCloudVmCluster#data_storage_size_in_t_bs}
	DataStorageSizeInTBs *float64 `field:"optional" json:"dataStorageSizeInTBs" yaml:"dataStorageSizeInTBs"`
	// The DB nodes that are implicitly created and managed as part of this VM Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#db_nodes OdbCloudVmCluster#db_nodes}
	DbNodes interface{} `field:"optional" json:"dbNodes" yaml:"dbNodes"`
	// The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#db_node_storage_size_in_g_bs OdbCloudVmCluster#db_node_storage_size_in_g_bs}
	DbNodeStorageSizeInGBs *float64 `field:"optional" json:"dbNodeStorageSizeInGBs" yaml:"dbNodeStorageSizeInGBs"`
	// The list of database servers for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#db_servers OdbCloudVmCluster#db_servers}
	DbServers *[]*string `field:"optional" json:"dbServers" yaml:"dbServers"`
	// The user-friendly name for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#display_name OdbCloudVmCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#gi_version OdbCloudVmCluster#gi_version}
	GiVersion *string `field:"optional" json:"giVersion" yaml:"giVersion"`
	// The host name for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#hostname OdbCloudVmCluster#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The AWS Identity and Access Management (IAM) service roles associated with the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#iam_roles OdbCloudVmCluster#iam_roles}
	IamRoles interface{} `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Indicates whether database backups to local Exadata storage is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#is_local_backup_enabled OdbCloudVmCluster#is_local_backup_enabled}
	IsLocalBackupEnabled interface{} `field:"optional" json:"isLocalBackupEnabled" yaml:"isLocalBackupEnabled"`
	// Indicates whether the VM cluster is configured with a sparse disk group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#is_sparse_diskgroup_enabled OdbCloudVmCluster#is_sparse_diskgroup_enabled}
	IsSparseDiskgroupEnabled interface{} `field:"optional" json:"isSparseDiskgroupEnabled" yaml:"isSparseDiskgroupEnabled"`
	// The Oracle license model applied to the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#license_model OdbCloudVmCluster#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#memory_size_in_g_bs OdbCloudVmCluster#memory_size_in_g_bs}
	MemorySizeInGBs *float64 `field:"optional" json:"memorySizeInGBs" yaml:"memorySizeInGBs"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#odb_network_id OdbCloudVmCluster#odb_network_id}
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// Property description not available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#scan_listener_port_tcp OdbCloudVmCluster#scan_listener_port_tcp}
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#ssh_public_keys OdbCloudVmCluster#ssh_public_keys}
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The operating system version of the image chosen for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#system_version OdbCloudVmCluster#system_version}
	SystemVersion *string `field:"optional" json:"systemVersion" yaml:"systemVersion"`
	// Tags to assign to the Vm Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#tags OdbCloudVmCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/odb_cloud_vm_cluster#time_zone OdbCloudVmCluster#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

