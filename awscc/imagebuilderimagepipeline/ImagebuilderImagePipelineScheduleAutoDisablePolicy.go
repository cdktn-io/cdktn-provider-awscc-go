// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineScheduleAutoDisablePolicy struct {
	// The number of consecutive failures after which the pipeline should be automatically disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/imagebuilder_image_pipeline#failure_count ImagebuilderImagePipeline#failure_count}
	FailureCount *float64 `field:"optional" json:"failureCount" yaml:"failureCount"`
}

