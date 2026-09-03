// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2networkperformancemetricsubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2NetworkPerformanceMetricSubscriptionConfig struct {
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
	// The target Region or Availability Zone for the metric to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_network_performance_metric_subscription#destination Ec2NetworkPerformanceMetricSubscription#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The metric type to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_network_performance_metric_subscription#metric Ec2NetworkPerformanceMetricSubscription#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The starting Region or Availability Zone for metric to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_network_performance_metric_subscription#source Ec2NetworkPerformanceMetricSubscription#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The statistic to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_network_performance_metric_subscription#statistic Ec2NetworkPerformanceMetricSubscription#statistic}
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

