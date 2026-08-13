// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LookoutequipmentInferenceSchedulerConfig struct {
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
	// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#data_input_configuration LookoutequipmentInferenceScheduler#data_input_configuration}
	DataInputConfiguration *LookoutequipmentInferenceSchedulerDataInputConfiguration `field:"required" json:"dataInputConfiguration" yaml:"dataInputConfiguration"`
	// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#data_output_configuration LookoutequipmentInferenceScheduler#data_output_configuration}
	DataOutputConfiguration *LookoutequipmentInferenceSchedulerDataOutputConfiguration `field:"required" json:"dataOutputConfiguration" yaml:"dataOutputConfiguration"`
	// How often data is uploaded to the source S3 bucket for the input data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#data_upload_frequency LookoutequipmentInferenceScheduler#data_upload_frequency}
	DataUploadFrequency *string `field:"required" json:"dataUploadFrequency" yaml:"dataUploadFrequency"`
	// The name of the previously trained ML model being used to create the inference scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#model_name LookoutequipmentInferenceScheduler#model_name}
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#role_arn LookoutequipmentInferenceScheduler#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A period of time (in minutes) by which inference on the data is delayed after the data starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#data_delay_offset_in_minutes LookoutequipmentInferenceScheduler#data_delay_offset_in_minutes}
	DataDelayOffsetInMinutes *float64 `field:"optional" json:"dataDelayOffsetInMinutes" yaml:"dataDelayOffsetInMinutes"`
	// The name of the inference scheduler being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#inference_scheduler_name LookoutequipmentInferenceScheduler#inference_scheduler_name}
	InferenceSchedulerName *string `field:"optional" json:"inferenceSchedulerName" yaml:"inferenceSchedulerName"`
	// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#server_side_kms_key_id LookoutequipmentInferenceScheduler#server_side_kms_key_id}
	ServerSideKmsKeyId *string `field:"optional" json:"serverSideKmsKeyId" yaml:"serverSideKmsKeyId"`
	// Any tags associated with the inference scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lookoutequipment_inference_scheduler#tags LookoutequipmentInferenceScheduler#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

