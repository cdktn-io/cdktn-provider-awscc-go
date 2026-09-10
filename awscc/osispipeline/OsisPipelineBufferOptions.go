// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipeline


type OsisPipelineBufferOptions struct {
	// Whether persistent buffering should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/osis_pipeline#persistent_buffer_enabled OsisPipeline#persistent_buffer_enabled}
	PersistentBufferEnabled interface{} `field:"optional" json:"persistentBufferEnabled" yaml:"persistentBufferEnabled"`
}

