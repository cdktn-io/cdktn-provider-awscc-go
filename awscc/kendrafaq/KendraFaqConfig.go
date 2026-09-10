// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrafaq

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraFaqConfig struct {
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
	// Index ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#index_id KendraFaq#index_id}
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// FAQ name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#name KendraFaq#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// FAQ role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#role_arn KendraFaq#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// FAQ S3 path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#s3_path KendraFaq#s3_path}
	S3Path *KendraFaqS3Path `field:"required" json:"s3Path" yaml:"s3Path"`
	// FAQ description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#description KendraFaq#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// FAQ file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#file_format KendraFaq#file_format}
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// The code for a language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#language_code KendraFaq#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Tags for labeling the FAQ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kendra_faq#tags KendraFaq#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

