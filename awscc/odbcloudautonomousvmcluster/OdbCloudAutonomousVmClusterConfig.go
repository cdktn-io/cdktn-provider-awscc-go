// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbCloudAutonomousVmClusterConfig struct {
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
	// The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#autonomous_data_storage_size_in_t_bs OdbCloudAutonomousVmCluster#autonomous_data_storage_size_in_t_bs}
	AutonomousDataStorageSizeInTBs *float64 `field:"optional" json:"autonomousDataStorageSizeInTBs" yaml:"autonomousDataStorageSizeInTBs"`
	// The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#cloud_exadata_infrastructure_id OdbCloudAutonomousVmCluster#cloud_exadata_infrastructure_id}
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The number of CPU cores enabled per node in the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#cpu_core_count_per_node OdbCloudAutonomousVmCluster#cpu_core_count_per_node}
	CpuCoreCountPerNode *float64 `field:"optional" json:"cpuCoreCountPerNode" yaml:"cpuCoreCountPerNode"`
	// The list of database servers associated with the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#db_servers OdbCloudAutonomousVmCluster#db_servers}
	DbServers *[]*string `field:"optional" json:"dbServers" yaml:"dbServers"`
	// The user-provided description of the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#description OdbCloudAutonomousVmCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#display_name OdbCloudAutonomousVmCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The AWS Identity and Access Management (IAM) service roles associated with the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#iam_roles OdbCloudAutonomousVmCluster#iam_roles}
	IamRoles interface{} `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#is_mtls_enabled_vm_cluster OdbCloudAutonomousVmCluster#is_mtls_enabled_vm_cluster}
	IsMtlsEnabledVmCluster interface{} `field:"optional" json:"isMtlsEnabledVmCluster" yaml:"isMtlsEnabledVmCluster"`
	// The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#license_model OdbCloudAutonomousVmCluster#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#maintenance_window OdbCloudAutonomousVmCluster#maintenance_window}
	MaintenanceWindow *OdbCloudAutonomousVmClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The amount of memory allocated per Oracle Compute Unit, in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#memory_per_oracle_compute_unit_in_g_bs OdbCloudAutonomousVmCluster#memory_per_oracle_compute_unit_in_g_bs}
	MemoryPerOracleComputeUnitInGBs *float64 `field:"optional" json:"memoryPerOracleComputeUnitInGBs" yaml:"memoryPerOracleComputeUnitInGBs"`
	// The unique identifier of the ODB network associated with this Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#odb_network_id OdbCloudAutonomousVmCluster#odb_network_id}
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The SCAN listener port for non-TLS (TCP) protocol. The default is 1521.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#scan_listener_port_non_tls OdbCloudAutonomousVmCluster#scan_listener_port_non_tls}
	ScanListenerPortNonTls *float64 `field:"optional" json:"scanListenerPortNonTls" yaml:"scanListenerPortNonTls"`
	// The SCAN listener port for TLS (TCP) protocol. The default is 2484.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#scan_listener_port_tls OdbCloudAutonomousVmCluster#scan_listener_port_tls}
	ScanListenerPortTls *float64 `field:"optional" json:"scanListenerPortTls" yaml:"scanListenerPortTls"`
	// The tags associated with the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#tags OdbCloudAutonomousVmCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#time_zone OdbCloudAutonomousVmCluster#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/odb_cloud_autonomous_vm_cluster#total_container_databases OdbCloudAutonomousVmCluster#total_container_databases}
	TotalContainerDatabases *float64 `field:"optional" json:"totalContainerDatabases" yaml:"totalContainerDatabases"`
}

