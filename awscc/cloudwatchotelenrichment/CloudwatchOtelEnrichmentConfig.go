// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchotelenrichment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchOtelEnrichmentConfig struct {
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
	// Removes metrics from the include set.
	//
	// Absent or empty means nothing is removed. Evaluated after IncludeFilters, so ExcludeFilters always wins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_otel_enrichment#exclude_filters CloudwatchOtelEnrichment#exclude_filters}
	ExcludeFilters interface{} `field:"optional" json:"excludeFilters" yaml:"excludeFilters"`
	// Scopes enrichment to a subset of the account's telemetry.
	//
	// Absent or empty means all namespaces are in scope. Present means only these are. The service enforces a combined cap of 100 selectors across IncludeFilters and ExcludeFilters, and rejects more than one selector for the same namespace within a direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/cloudwatch_otel_enrichment#include_filters CloudwatchOtelEnrichment#include_filters}
	IncludeFilters interface{} `field:"optional" json:"includeFilters" yaml:"includeFilters"`
}

