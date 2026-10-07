// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectintegrationassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectIntegrationAssociationConfig struct {
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
	// Amazon Connect instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#instance_id ConnectIntegrationAssociation#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// ARN of Integration being associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#integration_arn ConnectIntegrationAssociation#integration_arn}
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// Specifies the integration type to be associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#integration_type ConnectIntegrationAssociation#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The name of the external application. This is only supported for the EVENT integration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#source_application_name ConnectIntegrationAssociation#source_application_name}
	SourceApplicationName *string `field:"optional" json:"sourceApplicationName" yaml:"sourceApplicationName"`
	// The URL for the external application. This is only supported for the EVENT integration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#source_application_url ConnectIntegrationAssociation#source_application_url}
	SourceApplicationUrl *string `field:"optional" json:"sourceApplicationUrl" yaml:"sourceApplicationUrl"`
	// The type of the data source. This is only supported for the EVENT integration type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#source_type ConnectIntegrationAssociation#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/connect_integration_association#tags ConnectIntegrationAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

