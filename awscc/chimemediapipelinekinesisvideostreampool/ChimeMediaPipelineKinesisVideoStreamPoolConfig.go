// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediapipelinekinesisvideostreampool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeMediaPipelineKinesisVideoStreamPoolConfig struct {
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
	// The name of the Kinesis Video Stream Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_pipeline_kinesis_video_stream_pool#pool_name ChimeMediaPipelineKinesisVideoStreamPool#pool_name}
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// The configuration settings for the Kinesis video stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_pipeline_kinesis_video_stream_pool#stream_configuration ChimeMediaPipelineKinesisVideoStreamPool#stream_configuration}
	StreamConfiguration *ChimeMediaPipelineKinesisVideoStreamPoolStreamConfiguration `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// The tags associated with the Kinesis Video Stream Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_pipeline_kinesis_video_stream_pool#tags ChimeMediaPipelineKinesisVideoStreamPool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

