// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchserviceapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserviceApplicationConfig struct {
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
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#name OpensearchserviceApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of application configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#app_configs OpensearchserviceApplication#app_configs}
	AppConfigs interface{} `field:"optional" json:"appConfigs" yaml:"appConfigs"`
	// List of data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#data_sources OpensearchserviceApplication#data_sources}
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// The endpoint for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#endpoint OpensearchserviceApplication#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Options for configuring IAM Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#iam_identity_center_options OpensearchserviceApplication#iam_identity_center_options}
	IamIdentityCenterOptions *OpensearchserviceApplicationIamIdentityCenterOptions `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// The ARN of the KMS key used to encrypt the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#kms_key_arn OpensearchserviceApplication#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// An arbitrary set of tags (key-value pairs) for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/opensearchservice_application#tags OpensearchserviceApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

