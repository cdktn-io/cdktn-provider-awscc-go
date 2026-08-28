// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainnameaccessassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayDomainNameAccessAssociationConfig struct {
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
	// The source of the domain name access association resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_domain_name_access_association#access_association_source ApigatewayDomainNameAccessAssociation#access_association_source}
	AccessAssociationSource *string `field:"required" json:"accessAssociationSource" yaml:"accessAssociationSource"`
	// The source type of the domain name access association resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_domain_name_access_association#access_association_source_type ApigatewayDomainNameAccessAssociation#access_association_source_type}
	AccessAssociationSourceType *string `field:"required" json:"accessAssociationSourceType" yaml:"accessAssociationSourceType"`
	// The amazon resource name (ARN) of the domain name resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_domain_name_access_association#domain_name_arn ApigatewayDomainNameAccessAssociation#domain_name_arn}
	DomainNameArn *string `field:"required" json:"domainNameArn" yaml:"domainNameArn"`
	// An array of arbitrary tags (key-value pairs) to associate with the domainname access association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigateway_domain_name_access_association#tags ApigatewayDomainNameAccessAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

