// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication


type ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfig struct {
	// Specify a max age rule to restrict the length of time that application versions are retained for an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticbeanstalk_application#max_age_rule ElasticbeanstalkApplication#max_age_rule}
	MaxAgeRule *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule `field:"optional" json:"maxAgeRule" yaml:"maxAgeRule"`
	// Specify a max count rule to restrict the number of application versions that are retained for an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticbeanstalk_application#max_count_rule ElasticbeanstalkApplication#max_count_rule}
	MaxCountRule *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRule `field:"optional" json:"maxCountRule" yaml:"maxCountRule"`
}

