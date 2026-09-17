// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcustommetric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotCustomMetricConfig struct {
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
	// The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_custom_metric#metric_type IotCustomMetric#metric_type}
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// Field represents a friendly name in the console for the custom metric;
	//
	// it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_custom_metric#display_name IotCustomMetric#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_custom_metric#metric_name IotCustomMetric#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_custom_metric#tags IotCustomMetric#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

