// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster


type OdbCloudVmClusterDataCollectionOptions struct {
	// Indicates whether diagnostic collection is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#is_diagnostics_events_enabled OdbCloudVmCluster#is_diagnostics_events_enabled}
	IsDiagnosticsEventsEnabled interface{} `field:"optional" json:"isDiagnosticsEventsEnabled" yaml:"isDiagnosticsEventsEnabled"`
	// Indicates whether health monitoring is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#is_health_monitoring_enabled OdbCloudVmCluster#is_health_monitoring_enabled}
	IsHealthMonitoringEnabled interface{} `field:"optional" json:"isHealthMonitoringEnabled" yaml:"isHealthMonitoringEnabled"`
	// Indicates whether incident logs are enabled for the cloud VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/odb_cloud_vm_cluster#is_incident_logs_enabled OdbCloudVmCluster#is_incident_logs_enabled}
	IsIncidentLogsEnabled interface{} `field:"optional" json:"isIncidentLogsEnabled" yaml:"isIncidentLogsEnabled"`
}

