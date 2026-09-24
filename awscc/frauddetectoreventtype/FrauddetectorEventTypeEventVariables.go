// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectoreventtype


type FrauddetectorEventTypeEventVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#arn FrauddetectorEventType#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The time when the event type was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#created_time FrauddetectorEventType#created_time}
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#data_source FrauddetectorEventType#data_source}.
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#data_type FrauddetectorEventType#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#default_value FrauddetectorEventType#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#description FrauddetectorEventType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#inline FrauddetectorEventType#inline}.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// The time when the event type was last updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#last_updated_time FrauddetectorEventType#last_updated_time}
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#name FrauddetectorEventType#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#tags FrauddetectorEventType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/frauddetector_event_type#variable_type FrauddetectorEventType#variable_type}.
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

