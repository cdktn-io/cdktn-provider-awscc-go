// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotstream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotStreamConfig struct {
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
	// The files to stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_stream#files IotStream#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// An IAM role that allows the IoT service principal to access your S3 files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_stream#role_arn IotStream#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The stream ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_stream#stream_id IotStream#stream_id}
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
	// The description of the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_stream#description IotStream#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata which can be used to manage streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_stream#tags IotStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

