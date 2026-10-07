// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationSqsParametersMessageSystemAttributes struct {
	// The attribute value for the Binary data type, Base64-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#binary_value Eventsv2Subscriber#binary_value}
	BinaryValue *string `field:"optional" json:"binaryValue" yaml:"binaryValue"`
	// The attribute data type.
	//
	// For Amazon SQS targets, specify String, Number, or Binary, optionally with a custom label suffix such as Number.float. For Amazon SNS targets, specify String, String.Array, Number, or Binary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#data_type Eventsv2Subscriber#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The attribute value for the String and Number data types (and String.Array for Amazon SNS targets).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#string_value Eventsv2Subscriber#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

