// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigDataCaptureConfig struct {
	// A list of the JSON and CSV content type that the endpoint captures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#capture_content_type_header SagemakerEndpointConfigA#capture_content_type_header}
	CaptureContentTypeHeader *SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeader `field:"optional" json:"captureContentTypeHeader" yaml:"captureContentTypeHeader"`
	// Specifies whether the endpoint captures input data to your model, output data from your model, or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#capture_options SagemakerEndpointConfigA#capture_options}
	CaptureOptions interface{} `field:"optional" json:"captureOptions" yaml:"captureOptions"`
	// The S3 bucket where model monitor stores captured data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#destination_s3_uri SagemakerEndpointConfigA#destination_s3_uri}
	DestinationS3Uri *string `field:"optional" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Set to True to enable data capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#enable_capture SagemakerEndpointConfigA#enable_capture}
	EnableCapture interface{} `field:"optional" json:"enableCapture" yaml:"enableCapture"`
	// The percentage of data to capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#initial_sampling_percentage SagemakerEndpointConfigA#initial_sampling_percentage}
	InitialSamplingPercentage *float64 `field:"optional" json:"initialSamplingPercentage" yaml:"initialSamplingPercentage"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the captured data at rest using Amazon S3 server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#kms_key_id SagemakerEndpointConfigA#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

