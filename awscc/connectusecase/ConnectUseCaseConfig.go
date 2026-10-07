// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectusecase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUseCaseConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/connect_use_case#instance_id ConnectUseCase#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The identifier for the integration association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/connect_use_case#integration_association_id ConnectUseCase#integration_association_id}
	IntegrationAssociationId *string `field:"required" json:"integrationAssociationId" yaml:"integrationAssociationId"`
	// The type of use case to associate to the integration association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/connect_use_case#use_case_type ConnectUseCase#use_case_type}
	UseCaseType *string `field:"required" json:"useCaseType" yaml:"useCaseType"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/connect_use_case#tags ConnectUseCase#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

