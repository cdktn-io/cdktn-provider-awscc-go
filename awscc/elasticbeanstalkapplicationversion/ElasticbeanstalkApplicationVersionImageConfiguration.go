// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplicationversion


type ElasticbeanstalkApplicationVersionImageConfiguration struct {
	// Configuration for building a container image from source code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#build ElasticbeanstalkApplicationVersion#build}
	BuildAttribute *ElasticbeanstalkApplicationVersionImageConfigurationBuild `field:"optional" json:"buildAttribute" yaml:"buildAttribute"`
	// The container image source for this version, as an ECR image URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#source ElasticbeanstalkApplicationVersion#source}
	Source *ElasticbeanstalkApplicationVersionImageConfigurationSource `field:"optional" json:"source" yaml:"source"`
}

