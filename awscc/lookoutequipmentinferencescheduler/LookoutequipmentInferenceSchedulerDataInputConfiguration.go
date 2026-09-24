// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerDataInputConfiguration struct {
	// Specifies configuration information for the input data for the inference, including input data S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutequipment_inference_scheduler#s3_input_configuration LookoutequipmentInferenceScheduler#s3_input_configuration}
	S3InputConfiguration *LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfiguration `field:"required" json:"s3InputConfiguration" yaml:"s3InputConfiguration"`
	// Specifies configuration information for the input data for the inference, including timestamp format and delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutequipment_inference_scheduler#inference_input_name_configuration LookoutequipmentInferenceScheduler#inference_input_name_configuration}
	InferenceInputNameConfiguration *LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfiguration `field:"optional" json:"inferenceInputNameConfiguration" yaml:"inferenceInputNameConfiguration"`
	// Indicates the difference between your time zone and Greenwich Mean Time (GMT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lookoutequipment_inference_scheduler#input_time_zone_offset LookoutequipmentInferenceScheduler#input_time_zone_offset}
	InputTimeZoneOffset *string `field:"optional" json:"inputTimeZoneOffset" yaml:"inputTimeZoneOffset"`
}

