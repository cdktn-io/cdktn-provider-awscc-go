// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerDataOutputConfiguration struct {
	// Specifies configuration information for the output results from the inference, including output S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#s3_output_configuration LookoutequipmentInferenceScheduler#s3_output_configuration}
	S3OutputConfiguration *LookoutequipmentInferenceSchedulerDataOutputConfigurationS3OutputConfiguration `field:"required" json:"s3OutputConfiguration" yaml:"s3OutputConfiguration"`
	// The ID number for the AWS KMS key used to encrypt the inference output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#kms_key_id LookoutequipmentInferenceScheduler#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

