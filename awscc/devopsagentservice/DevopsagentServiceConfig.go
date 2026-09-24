// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentServiceConfig struct {
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
	// The type of service being registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#service_type DevopsagentService#service_type}
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// The name of the private connection to use for OAuth token exchange requests only.
	//
	// Cannot be specified when PrivateConnectionName is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#exchange_url_private_connection_name DevopsagentService#exchange_url_private_connection_name}
	ExchangeUrlPrivateConnectionName *string `field:"optional" json:"exchangeUrlPrivateConnectionName" yaml:"exchangeUrlPrivateConnectionName"`
	// The ARN of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#kms_key_arn DevopsagentService#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the private connection to use for VPC connectivity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#private_connection_name DevopsagentService#private_connection_name}
	PrivateConnectionName *string `field:"optional" json:"privateConnectionName" yaml:"privateConnectionName"`
	// Service-specific configuration details for create operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#service_details DevopsagentService#service_details}
	ServiceDetails *DevopsagentServiceServiceDetails `field:"optional" json:"serviceDetails" yaml:"serviceDetails"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#tags DevopsagentService#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the private connection to use for API calls (target URL) only.
	//
	// Cannot be specified when PrivateConnectionName is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#target_url_private_connection_name DevopsagentService#target_url_private_connection_name}
	TargetUrlPrivateConnectionName *string `field:"optional" json:"targetUrlPrivateConnectionName" yaml:"targetUrlPrivateConnectionName"`
}

