// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster


type OdbCloudVmClusterDbNodes struct {
	// The Oracle Cloud ID (OCID) of the backup IP address that's associated with the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#backup_ip_id OdbCloudVmCluster#backup_ip_id}
	BackupIpId *string `field:"optional" json:"backupIpId" yaml:"backupIpId"`
	// The OCID of the second backup virtual network interface card (VNIC) for the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#backup_vnic_2_id OdbCloudVmCluster#backup_vnic_2_id}
	BackupVnic2Id *string `field:"optional" json:"backupVnic2Id" yaml:"backupVnic2Id"`
	// The number of CPU cores enabled on the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#cpu_core_count OdbCloudVmCluster#cpu_core_count}
	CpuCoreCount *float64 `field:"optional" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The Amazon Resource Name (ARN) of the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#db_node_arn OdbCloudVmCluster#db_node_arn}
	DbNodeArn *string `field:"optional" json:"dbNodeArn" yaml:"dbNodeArn"`
	// The unique identifier of the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#db_node_id OdbCloudVmCluster#db_node_id}
	DbNodeId *string `field:"optional" json:"dbNodeId" yaml:"dbNodeId"`
	// The amount of local node storage, in gigabytes (GB), that's allocated on the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#db_node_storage_size_in_g_bs OdbCloudVmCluster#db_node_storage_size_in_g_bs}
	DbNodeStorageSizeInGBs *float64 `field:"optional" json:"dbNodeStorageSizeInGBs" yaml:"dbNodeStorageSizeInGBs"`
	// The unique identifier of the database server that's associated with the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#db_server_id OdbCloudVmCluster#db_server_id}
	DbServerId *string `field:"optional" json:"dbServerId" yaml:"dbServerId"`
	// The OCID of the DB system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#db_system_id OdbCloudVmCluster#db_system_id}
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// The OCID of the host IP address that's associated with the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#host_ip_id OdbCloudVmCluster#host_ip_id}
	HostIpId *string `field:"optional" json:"hostIpId" yaml:"hostIpId"`
	// The host name for the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#hostname OdbCloudVmCluster#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The amount of memory, in gigabytes (GB), that allocated on the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#memory_size_in_g_bs OdbCloudVmCluster#memory_size_in_g_bs}
	MemorySizeInGBs *float64 `field:"optional" json:"memorySizeInGBs" yaml:"memorySizeInGBs"`
	// The OCID of the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#ocid OdbCloudVmCluster#ocid}
	Ocid *string `field:"optional" json:"ocid" yaml:"ocid"`
	// The current status of the DB node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#status OdbCloudVmCluster#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#tags OdbCloudVmCluster#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The OCID of the second VNIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#vnic_2_id OdbCloudVmCluster#vnic_2_id}
	Vnic2Id *string `field:"optional" json:"vnic2Id" yaml:"vnic2Id"`
	// The OCID of the VNIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#vnic_id OdbCloudVmCluster#vnic_id}
	VnicId *string `field:"optional" json:"vnicId" yaml:"vnicId"`
}

