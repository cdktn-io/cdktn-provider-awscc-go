// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2applicationstatuscheck

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2ApplicationStatusCheckConfig struct {
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
	// The port used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#port Ec2ApplicationStatusCheck#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network protocol used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#protocol Ec2ApplicationStatusCheck#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Whether this check is included in the rolled-up application status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#aggregation Ec2ApplicationStatusCheck#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The network interface device index used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#device_index Ec2ApplicationStatusCheck#device_index}
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The number of consecutive failed probes required to mark the instance unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#failure_threshold Ec2ApplicationStatusCheck#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The source/destination network paths used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#health_check_paths Ec2ApplicationStatusCheck#health_check_paths}
	HealthCheckPaths interface{} `field:"optional" json:"healthCheckPaths" yaml:"healthCheckPaths"`
	// Seconds to wait after instance launch before beginning health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#initialization_grace_period_seconds Ec2ApplicationStatusCheck#initialization_grace_period_seconds}
	InitializationGracePeriodSeconds *float64 `field:"optional" json:"initializationGracePeriodSeconds" yaml:"initializationGracePeriodSeconds"`
	// The interval, in seconds, between health check probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#interval Ec2ApplicationStatusCheck#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The IP scope used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#ip_scope Ec2ApplicationStatusCheck#ip_scope}
	IpScope *string `field:"optional" json:"ipScope" yaml:"ipScope"`
	// The IP version used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#ip_version Ec2ApplicationStatusCheck#ip_version}
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// The HTTP path used for the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#path Ec2ApplicationStatusCheck#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP status codes considered successful (e.g., "200-299").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#status_code_matcher Ec2ApplicationStatusCheck#status_code_matcher}
	StatusCodeMatcher *string `field:"optional" json:"statusCodeMatcher" yaml:"statusCodeMatcher"`
	// The number of consecutive successful probes required to mark the instance healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#success_threshold Ec2ApplicationStatusCheck#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Tags to apply to the application status check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#tags Ec2ApplicationStatusCheck#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The timeout, in seconds, for each health check probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#timeout Ec2ApplicationStatusCheck#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

