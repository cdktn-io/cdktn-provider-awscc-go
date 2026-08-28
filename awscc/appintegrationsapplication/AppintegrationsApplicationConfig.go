// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationsapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppintegrationsApplicationConfig struct {
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
	// Application source config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#application_source_config AppintegrationsApplication#application_source_config}
	ApplicationSourceConfig *AppintegrationsApplicationApplicationSourceConfig `field:"required" json:"applicationSourceConfig" yaml:"applicationSourceConfig"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#name AppintegrationsApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#namespace AppintegrationsApplication#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The application configuration. Cannot be used when IsService is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#application_config AppintegrationsApplication#application_config}
	ApplicationConfig *AppintegrationsApplicationApplicationConfig `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The type of application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#application_type AppintegrationsApplication#application_type}
	ApplicationType *string `field:"optional" json:"applicationType" yaml:"applicationType"`
	// The application description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#description AppintegrationsApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The iframe configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#iframe_config AppintegrationsApplication#iframe_config}
	IframeConfig *AppintegrationsApplicationIframeConfig `field:"optional" json:"iframeConfig" yaml:"iframeConfig"`
	// The initialization timeout in milliseconds. Required when IsService is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#initialization_timeout AppintegrationsApplication#initialization_timeout}
	InitializationTimeout *float64 `field:"optional" json:"initializationTimeout" yaml:"initializationTimeout"`
	// Indicates if the application is a service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#is_service AppintegrationsApplication#is_service}
	IsService interface{} `field:"optional" json:"isService" yaml:"isService"`
	// The configuration of events or requests that the application has access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#permissions AppintegrationsApplication#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The tags (keys and values) associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_application#tags AppintegrationsApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

