// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizedataset


type PersonalizeDatasetDatasetImportJob struct {
	// The ARN of the dataset that receives the imported data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_dataset#dataset_arn PersonalizeDataset#dataset_arn}
	DatasetArn *string `field:"optional" json:"datasetArn" yaml:"datasetArn"`
	// The ARN of the dataset import job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_dataset#dataset_import_job_arn PersonalizeDataset#dataset_import_job_arn}
	DatasetImportJobArn *string `field:"optional" json:"datasetImportJobArn" yaml:"datasetImportJobArn"`
	// The Amazon S3 bucket that contains the training data to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_dataset#data_source PersonalizeDataset#data_source}
	DataSource *PersonalizeDatasetDatasetImportJobDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The name for the dataset import job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_dataset#job_name PersonalizeDataset#job_name}
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The ARN of the IAM role that has permissions to read from the Amazon S3 data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_dataset#role_arn PersonalizeDataset#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

