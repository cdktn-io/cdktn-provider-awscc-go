// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueMlTransformConfig struct {
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
	// A list of AWS Glue table definitions used by the transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#input_record_tables GlueMlTransform#input_record_tables}
	InputRecordTables *GlueMlTransformInputRecordTables `field:"required" json:"inputRecordTables" yaml:"inputRecordTables"`
	// The name or ARN of the IAM role with the required permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#role GlueMlTransform#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The algorithm-specific parameters that are associated with the machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#transform_parameters GlueMlTransform#transform_parameters}
	TransformParameters *GlueMlTransformTransformParameters `field:"required" json:"transformParameters" yaml:"transformParameters"`
	// A user-defined, long-form description text for the machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#description GlueMlTransform#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of AWS Glue this machine learning transform is compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#glue_version GlueMlTransform#glue_version}
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// The number of AWS Glue DPUs allocated to task runs for this transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#max_capacity GlueMlTransform#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry after an MLTaskRun fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#max_retries GlueMlTransform#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// A user-defined name for the machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#name GlueMlTransform#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of workers of a defined workerType that are allocated when a task runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#number_of_workers GlueMlTransform#number_of_workers}
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The tags to use with this machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#tags GlueMlTransform#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The timeout in minutes of the machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#timeout GlueMlTransform#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The encryption-at-rest settings of the transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#transform_encryption GlueMlTransform#transform_encryption}
	TransformEncryption *GlueMlTransformTransformEncryption `field:"optional" json:"transformEncryption" yaml:"transformEncryption"`
	// The type of predefined worker that is allocated when a task runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_ml_transform#worker_type GlueMlTransform#worker_type}
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

