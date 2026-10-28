// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakefhirdatastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HealthlakeFhirDatastoreConfig struct {
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
	// The FHIR version. Only R4 version data is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#datastore_type_version HealthlakeFhirDatastore#datastore_type_version}
	DatastoreTypeVersion *string `field:"required" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The analytics configuration for the Data Store. Omitted when the Data Store can't carry analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#analytics_configuration HealthlakeFhirDatastore#analytics_configuration}
	AnalyticsConfiguration *HealthlakeFhirDatastoreAnalyticsConfiguration `field:"optional" json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// The user-generated name for the Data Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#datastore_name HealthlakeFhirDatastore#datastore_name}
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The identity provider configuration for the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#identity_provider_configuration HealthlakeFhirDatastore#identity_provider_configuration}
	IdentityProviderConfiguration *HealthlakeFhirDatastoreIdentityProviderConfiguration `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// The natural language processing (NLP) configuration for the Data Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#nlp_configuration HealthlakeFhirDatastore#nlp_configuration}
	NlpConfiguration *HealthlakeFhirDatastoreNlpConfiguration `field:"optional" json:"nlpConfiguration" yaml:"nlpConfiguration"`
	// The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#preload_data_config HealthlakeFhirDatastore#preload_data_config}
	PreloadDataConfig *HealthlakeFhirDatastorePreloadDataConfig `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The profile configuration for the Data Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#profile_configuration HealthlakeFhirDatastore#profile_configuration}
	ProfileConfiguration *HealthlakeFhirDatastoreProfileConfiguration `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// The server-side encryption key configuration for a customer provided encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#sse_configuration HealthlakeFhirDatastore#sse_configuration}
	SseConfiguration *HealthlakeFhirDatastoreSseConfiguration `field:"optional" json:"sseConfiguration" yaml:"sseConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/healthlake_fhir_datastore#tags HealthlakeFhirDatastore#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

