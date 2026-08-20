// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationaggregator

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigConfigurationAggregatorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#account_aggregation_sources ConfigConfigurationAggregator#account_aggregation_sources}.
	AccountAggregationSources interface{} `field:"optional" json:"accountAggregationSources" yaml:"accountAggregationSources"`
	// The name of the aggregator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#configuration_aggregator_name ConfigConfigurationAggregator#configuration_aggregator_name}
	ConfigurationAggregatorName *string `field:"optional" json:"configurationAggregatorName" yaml:"configurationAggregatorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#organization_aggregation_source ConfigConfigurationAggregator#organization_aggregation_source}.
	OrganizationAggregationSource *ConfigConfigurationAggregatorOrganizationAggregationSource `field:"optional" json:"organizationAggregationSource" yaml:"organizationAggregationSource"`
	// The tags for the configuration aggregator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#tags ConfigConfigurationAggregator#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

