// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftusagelimit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftUsageLimitConfig struct {
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
	// The limit amount.
	//
	// If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#amount RedshiftUsageLimit#amount}
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The identifier of the cluster that you want to limit usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#cluster_identifier RedshiftUsageLimit#cluster_identifier}
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The Amazon Redshift feature that you want to limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#feature_type RedshiftUsageLimit#feature_type}
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
	// The type of limit.
	//
	// Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#limit_type RedshiftUsageLimit#limit_type}
	LimitType *string `field:"required" json:"limitType" yaml:"limitType"`
	// The action that Amazon Redshift takes when the limit is reached. The default is log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#breach_action RedshiftUsageLimit#breach_action}
	BreachAction *string `field:"optional" json:"breachAction" yaml:"breachAction"`
	// The time period that the amount applies to. A weekly period begins on Sunday. The default is monthly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#period RedshiftUsageLimit#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// A list of tag instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_usage_limit#tags RedshiftUsageLimit#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

