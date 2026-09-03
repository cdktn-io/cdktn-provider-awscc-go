// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewjob

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabrewJobConfig struct {
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
	// Job name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#name DatabrewJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#role_arn DatabrewJob#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Job type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#type DatabrewJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#database_outputs DatabrewJob#database_outputs}.
	DatabaseOutputs interface{} `field:"optional" json:"databaseOutputs" yaml:"databaseOutputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#data_catalog_outputs DatabrewJob#data_catalog_outputs}.
	DataCatalogOutputs interface{} `field:"optional" json:"dataCatalogOutputs" yaml:"dataCatalogOutputs"`
	// Dataset name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#dataset_name DatabrewJob#dataset_name}
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// Encryption Key Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#encryption_key_arn DatabrewJob#encryption_key_arn}
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Encryption mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#encryption_mode DatabrewJob#encryption_mode}
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// Job Sample.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#job_sample DatabrewJob#job_sample}
	JobSample *DatabrewJobJobSample `field:"optional" json:"jobSample" yaml:"jobSample"`
	// Log subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#log_subscription DatabrewJob#log_subscription}
	LogSubscription *string `field:"optional" json:"logSubscription" yaml:"logSubscription"`
	// Max capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#max_capacity DatabrewJob#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Max retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#max_retries DatabrewJob#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Output location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#output_location DatabrewJob#output_location}
	OutputLocation *DatabrewJobOutputLocation `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#outputs DatabrewJob#outputs}.
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// Profile Job configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#profile_configuration DatabrewJob#profile_configuration}
	ProfileConfiguration *DatabrewJobProfileConfiguration `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// Project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#project_name DatabrewJob#project_name}
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#recipe DatabrewJob#recipe}.
	Recipe *DatabrewJobRecipe `field:"optional" json:"recipe" yaml:"recipe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#tags DatabrewJob#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#timeout DatabrewJob#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Data quality rules configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/databrew_job#validation_configurations DatabrewJob#validation_configurations}
	ValidationConfigurations interface{} `field:"optional" json:"validationConfigurations" yaml:"validationConfigurations"`
}

