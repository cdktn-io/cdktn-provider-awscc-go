// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplicationversion


type ElasticbeanstalkApplicationVersionImageConfigurationBuild struct {
	// The target architecture for the built container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#architecture ElasticbeanstalkApplicationVersion#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// The buildpack to use for building the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#buildpack ElasticbeanstalkApplicationVersion#buildpack}
	Buildpack *string `field:"optional" json:"buildpack" yaml:"buildpack"`
	// The ARN of the IAM role that AWS CodeBuild assumes to build the application version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#code_build_service_role ElasticbeanstalkApplicationVersion#code_build_service_role}
	CodeBuildServiceRole *string `field:"optional" json:"codeBuildServiceRole" yaml:"codeBuildServiceRole"`
	// The compute type for the CodeBuild build environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#compute_type ElasticbeanstalkApplicationVersion#compute_type}
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The path to the Dockerfile, relative to the source root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#dockerfile_location ElasticbeanstalkApplicationVersion#dockerfile_location}
	DockerfileLocation *string `field:"optional" json:"dockerfileLocation" yaml:"dockerfileLocation"`
	// The timeout for the CodeBuild build, in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#timeout_in_minutes ElasticbeanstalkApplicationVersion#timeout_in_minutes}
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// The type of image build: docker or buildpack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticbeanstalk_application_version#type ElasticbeanstalkApplicationVersion#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

