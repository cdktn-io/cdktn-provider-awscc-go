// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectordetector


type FrauddetectorDetectorEventTypeEventVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#arn FrauddetectorDetector#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The time when the event variable was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#created_time FrauddetectorDetector#created_time}
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#data_source FrauddetectorDetector#data_source}.
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#data_type FrauddetectorDetector#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#default_value FrauddetectorDetector#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#description FrauddetectorDetector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#inline FrauddetectorDetector#inline}.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// The time when the event variable was last updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#last_updated_time FrauddetectorDetector#last_updated_time}
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#name FrauddetectorDetector#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with this event variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#tags FrauddetectorDetector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_detector#variable_type FrauddetectorDetector#variable_type}.
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

