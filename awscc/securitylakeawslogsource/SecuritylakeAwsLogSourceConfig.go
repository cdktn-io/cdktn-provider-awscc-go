// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitylakeawslogsource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecuritylakeAwsLogSourceConfig struct {
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
	// The ARN for the data lake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_aws_log_source#data_lake_arn SecuritylakeAwsLogSource#data_lake_arn}
	DataLakeArn *string `field:"required" json:"dataLakeArn" yaml:"dataLakeArn"`
	// The name for a AWS source. This must be a Regionally unique value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_aws_log_source#source_name SecuritylakeAwsLogSource#source_name}
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The version for a AWS source. This must be a Regionally unique value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_aws_log_source#source_version SecuritylakeAwsLogSource#source_version}
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
	// AWS account where you want to collect logs from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/securitylake_aws_log_source#accounts SecuritylakeAwsLogSource#accounts}
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
}

