// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsReplicationConfigConfig struct {
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
	// Configuration parameters for provisioning a AWS DMS Serverless replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#compute_config DmsReplicationConfig#compute_config}
	ComputeConfig *DmsReplicationConfigComputeConfig `field:"required" json:"computeConfig" yaml:"computeConfig"`
	// A unique identifier of replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#replication_config_identifier DmsReplicationConfig#replication_config_identifier}
	ReplicationConfigIdentifier *string `field:"required" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#replication_type DmsReplicationConfig#replication_type}
	ReplicationType *string `field:"required" json:"replicationType" yaml:"replicationType"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#source_endpoint_arn DmsReplicationConfig#source_endpoint_arn}
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#table_mappings DmsReplicationConfig#table_mappings}
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#target_endpoint_arn DmsReplicationConfig#target_endpoint_arn}
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#replication_settings DmsReplicationConfig#replication_settings}
	ReplicationSettings *string `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#resource_identifier DmsReplicationConfig#resource_identifier}
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// JSON settings for specifying supplemental data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#supplemental_settings DmsReplicationConfig#supplemental_settings}
	SupplementalSettings *string `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_replication_config#tags DmsReplicationConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

