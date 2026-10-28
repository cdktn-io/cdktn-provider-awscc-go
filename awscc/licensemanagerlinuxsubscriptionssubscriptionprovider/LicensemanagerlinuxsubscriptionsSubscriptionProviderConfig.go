// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlinuxsubscriptionssubscriptionprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerlinuxsubscriptionsSubscriptionProviderConfig struct {
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
	// The Amazon Resource Name (ARN) of the Secrets Manager secret that contains the access token for the third-party Linux subscription provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanagerlinuxsubscriptions_subscription_provider#secret_arn LicensemanagerlinuxsubscriptionsSubscriptionProvider#secret_arn}
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The Linux subscription provider that the registration is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanagerlinuxsubscriptions_subscription_provider#subscription_provider_source LicensemanagerlinuxsubscriptionsSubscriptionProvider#subscription_provider_source}
	SubscriptionProviderSource *string `field:"required" json:"subscriptionProviderSource" yaml:"subscriptionProviderSource"`
	// The metadata tags assigned to the registered subscription provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/licensemanagerlinuxsubscriptions_subscription_provider#tags LicensemanagerlinuxsubscriptionsSubscriptionProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

