// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication


type ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRule struct {
	// Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#delete_source_from_s3 ElasticbeanstalkApplication#delete_source_from_s3}
	DeleteSourceFromS3 interface{} `field:"optional" json:"deleteSourceFromS3" yaml:"deleteSourceFromS3"`
	// Specify true to apply the rule, or false to disable it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#enabled ElasticbeanstalkApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the maximum number of application versions to retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#max_count ElasticbeanstalkApplication#max_count}
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
}

