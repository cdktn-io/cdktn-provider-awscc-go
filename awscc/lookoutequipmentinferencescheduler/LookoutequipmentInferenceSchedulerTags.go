// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerTags struct {
	// The key for the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lookoutequipment_inference_scheduler#key LookoutequipmentInferenceScheduler#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lookoutequipment_inference_scheduler#value LookoutequipmentInferenceScheduler#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

