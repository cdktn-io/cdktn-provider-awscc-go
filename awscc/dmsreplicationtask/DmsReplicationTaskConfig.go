// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationtask

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsReplicationTaskConfig struct {
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
	// The migration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#migration_type DmsReplicationTask#migration_type}
	MigrationType *string `field:"required" json:"migrationType" yaml:"migrationType"`
	// The Amazon Resource Name (ARN) of a replication instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#replication_instance_arn DmsReplicationTask#replication_instance_arn}
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#source_endpoint_arn DmsReplicationTask#source_endpoint_arn}
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// The table mappings for the task, in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#table_mappings DmsReplicationTask#table_mappings}
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#target_endpoint_arn DmsReplicationTask#target_endpoint_arn}
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Indicates when you want a change data capture (CDC) operation to start.
	//
	// Use either CdcStartPosition or CdcStartTime to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#cdc_start_position DmsReplicationTask#cdc_start_position}
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Indicates the start time for a change data capture (CDC) operation.
	//
	// Use either CdcStartTime or CdcStartPosition to specify when you want a CDC operation to start. Specifying both values results in an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#cdc_start_time DmsReplicationTask#cdc_start_time}
	CdcStartTime *float64 `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// Indicates when you want a change data capture (CDC) operation to stop.
	//
	// The value can be either server time or commit time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#cdc_stop_position DmsReplicationTask#cdc_stop_position}
	CdcStopPosition *string `field:"optional" json:"cdcStopPosition" yaml:"cdcStopPosition"`
	// An identifier for the replication task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#replication_task_identifier DmsReplicationTask#replication_task_identifier}
	ReplicationTaskIdentifier *string `field:"optional" json:"replicationTaskIdentifier" yaml:"replicationTaskIdentifier"`
	// Overall settings for the task, in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#replication_task_settings DmsReplicationTask#replication_task_settings}
	ReplicationTaskSettings *string `field:"optional" json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// A friendly name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#resource_identifier DmsReplicationTask#resource_identifier}
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#tags DmsReplicationTask#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Supplemental information that the task requires to migrate the data for certain source and target endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_replication_task#task_data DmsReplicationTask#task_data}
	TaskData *string `field:"optional" json:"taskData" yaml:"taskData"`
}

