// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudtrailEventDataStoreConfig struct {
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
	// The advanced event selectors that were used to select events for the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#advanced_event_selectors CloudtrailEventDataStore#advanced_event_selectors}
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// The mode that the event data store will use to charge for event storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#billing_mode CloudtrailEventDataStore#billing_mode}
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// An array that enriches event records in an existing event data store by including additional information specified in individual ContexKeySelector entries.
	//
	// If you add ContextKeySelectors, you must set MaxEventSize to Large.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#context_key_selectors CloudtrailEventDataStore#context_key_selectors}
	ContextKeySelectors interface{} `field:"optional" json:"contextKeySelectors" yaml:"contextKeySelectors"`
	// Indicates whether federation is enabled on an event data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#federation_enabled CloudtrailEventDataStore#federation_enabled}
	FederationEnabled interface{} `field:"optional" json:"federationEnabled" yaml:"federationEnabled"`
	// The ARN of the role used for event data store federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#federation_role_arn CloudtrailEventDataStore#federation_role_arn}
	FederationRoleArn *string `field:"optional" json:"federationRoleArn" yaml:"federationRoleArn"`
	// Indicates whether the event data store is ingesting events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#ingestion_enabled CloudtrailEventDataStore#ingestion_enabled}
	IngestionEnabled interface{} `field:"optional" json:"ingestionEnabled" yaml:"ingestionEnabled"`
	// Specifies the ARN of the event data store that will collect Insights events.
	//
	// Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#insights_destination CloudtrailEventDataStore#insights_destination}
	InsightsDestination *string `field:"optional" json:"insightsDestination" yaml:"insightsDestination"`
	// Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing event data store.
	//
	// Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#insight_selectors CloudtrailEventDataStore#insight_selectors}
	InsightSelectors interface{} `field:"optional" json:"insightSelectors" yaml:"insightSelectors"`
	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#kms_key_id CloudtrailEventDataStore#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the maximum size allowed for the event.
	//
	// Valid values are Standard and Large. If you add ContextKeySelectors, this value must be set to Large.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#max_event_size CloudtrailEventDataStore#max_event_size}
	MaxEventSize *string `field:"optional" json:"maxEventSize" yaml:"maxEventSize"`
	// Indicates whether the event data store includes events from all regions, or only from the region in which it was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#multi_region_enabled CloudtrailEventDataStore#multi_region_enabled}
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// The name of the event data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#name CloudtrailEventDataStore#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates that an event data store is collecting logged events for an organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#organization_enabled CloudtrailEventDataStore#organization_enabled}
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// The retention period, in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#retention_period CloudtrailEventDataStore#retention_period}
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#tags CloudtrailEventDataStore#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the event data store is protected from termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_event_data_store#termination_protection_enabled CloudtrailEventDataStore#termination_protection_enabled}
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

