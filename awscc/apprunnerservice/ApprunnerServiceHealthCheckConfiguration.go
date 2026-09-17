// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceHealthCheckConfiguration struct {
	// Health check Healthy Threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#healthy_threshold ApprunnerService#healthy_threshold}
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Health check Interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#interval ApprunnerService#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Health check Path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#path ApprunnerService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Health Check Protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#protocol ApprunnerService#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Health check Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#timeout ApprunnerService#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Health check Unhealthy Threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apprunner_service#unhealthy_threshold ApprunnerService#unhealthy_threshold}
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

