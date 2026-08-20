// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesIntegrationConfig struct {
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
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#domain_name CustomerprofilesIntegration#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A list of unique names for active event triggers associated with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#event_trigger_names CustomerprofilesIntegration#event_trigger_names}
	EventTriggerNames *[]*string `field:"optional" json:"eventTriggerNames" yaml:"eventTriggerNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#flow_definition CustomerprofilesIntegration#flow_definition}.
	FlowDefinition *CustomerprofilesIntegrationFlowDefinition `field:"optional" json:"flowDefinition" yaml:"flowDefinition"`
	// The name of the ObjectType defined for the 3rd party data in Profile Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#object_type_name CustomerprofilesIntegration#object_type_name}
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The mapping between 3rd party event types and ObjectType names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#object_type_names CustomerprofilesIntegration#object_type_names}
	ObjectTypeNames interface{} `field:"optional" json:"objectTypeNames" yaml:"objectTypeNames"`
	// Scope of the integration, such as 'PROFILE' or 'DOMAIN'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#scope CustomerprofilesIntegration#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The tags (keys and values) associated with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#tags CustomerprofilesIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The URI of the S3 bucket or any other type of data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/customerprofiles_integration#uri CustomerprofilesIntegration#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

