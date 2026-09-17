// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerDataOutputConfigurationS3OutputConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutequipment_inference_scheduler#bucket LookoutequipmentInferenceScheduler#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutequipment_inference_scheduler#prefix LookoutequipmentInferenceScheduler#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

