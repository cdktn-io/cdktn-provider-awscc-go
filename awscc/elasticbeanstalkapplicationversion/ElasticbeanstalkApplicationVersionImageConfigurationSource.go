// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplicationversion


type ElasticbeanstalkApplicationVersionImageConfigurationSource struct {
	// The URI of the container image, e.g. an ECR image URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#uri ElasticbeanstalkApplicationVersion#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

