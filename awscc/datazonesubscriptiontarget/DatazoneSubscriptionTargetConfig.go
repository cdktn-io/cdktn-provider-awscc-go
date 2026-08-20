// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonesubscriptiontarget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneSubscriptionTargetConfig struct {
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
	// The asset types that can be included in the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#applicable_asset_types DatazoneSubscriptionTarget#applicable_asset_types}
	ApplicableAssetTypes *[]*string `field:"required" json:"applicableAssetTypes" yaml:"applicableAssetTypes"`
	// The authorized principals of the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#authorized_principals DatazoneSubscriptionTarget#authorized_principals}
	AuthorizedPrincipals *[]*string `field:"required" json:"authorizedPrincipals" yaml:"authorizedPrincipals"`
	// The ID of the Amazon DataZone domain in which subscription target would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#domain_identifier DatazoneSubscriptionTarget#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the environment in which subscription target would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#environment_identifier DatazoneSubscriptionTarget#environment_identifier}
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The name of the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#name DatazoneSubscriptionTarget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration of the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#subscription_target_config DatazoneSubscriptionTarget#subscription_target_config}
	SubscriptionTargetConfig interface{} `field:"required" json:"subscriptionTargetConfig" yaml:"subscriptionTargetConfig"`
	// The type of the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#type DatazoneSubscriptionTarget#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The manage access role that is used to create the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#manage_access_role DatazoneSubscriptionTarget#manage_access_role}
	ManageAccessRole *string `field:"optional" json:"manageAccessRole" yaml:"manageAccessRole"`
	// The provider of the subscription target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_subscription_target#provider_name DatazoneSubscriptionTarget#provider_name}
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
}

