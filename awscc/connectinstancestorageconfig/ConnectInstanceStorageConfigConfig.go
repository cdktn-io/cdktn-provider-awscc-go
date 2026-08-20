// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectInstanceStorageConfigConfig struct {
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
	// Connect Instance ID with which the storage config will be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#instance_arn ConnectInstanceStorageConfig#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Specifies the type of storage resource available for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#resource_type ConnectInstanceStorageConfig#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Specifies the storage type to be associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#storage_type ConnectInstanceStorageConfig#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#kinesis_firehose_config ConnectInstanceStorageConfig#kinesis_firehose_config}.
	KinesisFirehoseConfig *ConnectInstanceStorageConfigKinesisFirehoseConfig `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#kinesis_stream_config ConnectInstanceStorageConfig#kinesis_stream_config}.
	KinesisStreamConfig *ConnectInstanceStorageConfigKinesisStreamConfig `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#kinesis_video_stream_config ConnectInstanceStorageConfig#kinesis_video_stream_config}.
	KinesisVideoStreamConfig *ConnectInstanceStorageConfigKinesisVideoStreamConfig `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_instance_storage_config#s3_config ConnectInstanceStorageConfig#s3_config}.
	S3Config *ConnectInstanceStorageConfigS3Config `field:"optional" json:"s3Config" yaml:"s3Config"`
}

