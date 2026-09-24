// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package piperfreports

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PiPerfReportsConfig struct {
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
	// The end time defined for the analysis report in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pi_perf_reports#end_time PiPerfReports#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// An immutable, AWS Region-unique identifier for a data source.
	//
	// Performance Insights gathers metrics from this data source. To use an Amazon RDS instance as a data source, specify its DbiResourceId value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pi_perf_reports#identifier PiPerfReports#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The AWS service for which Performance Insights returns metrics. Valid values are RDS and DOCDB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pi_perf_reports#service_type PiPerfReports#service_type}
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// The start time defined for the analysis report in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pi_perf_reports#start_time PiPerfReports#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pi_perf_reports#tags PiPerfReports#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

