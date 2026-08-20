// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingInputsS3Input struct {
	// The local path in your container where you want Amazon SageMaker to write input data to.
	//
	// `LocalPath` is an absolute path to the input data and must begin with `/opt/ml/processing/`. LocalPath is a required parameter when `AppManaged` is `False` (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#local_path SagemakerProcessingJob#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Whether to GZIP-decompress the data in Amazon S3 as it is streamed into the processing container.
	//
	// `Gzip` can only be used when `Pipe` mode is specified as the `S3InputMode`. In `Pipe` mode, Amazon SageMaker streams input data from the source directly to your container without using the EBS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_compression_type SagemakerProcessingJob#s3_compression_type}
	S3CompressionType *string `field:"optional" json:"s3CompressionType" yaml:"s3CompressionType"`
	// Whether to distribute the data from Amazon S3 to all processing instances with `FullyReplicated`, or whether the data from Amazon S3 is shared by Amazon S3 key, downloading one shard of data to each processing instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_data_distribution_type SagemakerProcessingJob#s3_data_distribution_type}
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether you use an S3Prefix or a ManifestFile for the data type.
	//
	// If you choose S3Prefix, S3Uri identifies a key name prefix. Amazon SageMaker uses all objects with the specified key name prefix for the processing job. If you choose ManifestFile, S3Uri identifies an object that is a manifest file containing a list of object keys that you want Amazon SageMaker to use for the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_data_type SagemakerProcessingJob#s3_data_type}
	S3DataType *string `field:"optional" json:"s3DataType" yaml:"s3DataType"`
	// Whether to use File or Pipe input mode.
	//
	// In File mode, Amazon SageMaker copies the data from the input source onto the local ML storage volume before starting your processing container. This is the most commonly used input mode. In Pipe mode, Amazon SageMaker streams input data from the source directly to your processing container into named pipes without using the ML storage volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_input_mode SagemakerProcessingJob#s3_input_mode}
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
	// The URI of the Amazon S3 prefix Amazon SageMaker downloads data required to run a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_uri SagemakerProcessingJob#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

