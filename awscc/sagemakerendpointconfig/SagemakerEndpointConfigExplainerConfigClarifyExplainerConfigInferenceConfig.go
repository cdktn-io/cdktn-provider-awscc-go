// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfig struct {
	// A template string used to format a JSON record into an acceptable model container input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#content_template SagemakerEndpointConfigA#content_template}
	ContentTemplate *string `field:"optional" json:"contentTemplate" yaml:"contentTemplate"`
	// The names of the features.
	//
	// If provided, these are included in the endpoint response payload to help readability of the InvokeEndpoint output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#feature_headers SagemakerEndpointConfigA#feature_headers}
	FeatureHeaders *[]*string `field:"optional" json:"featureHeaders" yaml:"featureHeaders"`
	// Provides the JMESPath expression to extract the features from a model container input in JSON Lines format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#features_attribute SagemakerEndpointConfigA#features_attribute}
	FeaturesAttribute *string `field:"optional" json:"featuresAttribute" yaml:"featuresAttribute"`
	// A list of data types of the features (optional).
	//
	// Applicable only to NLP explainability. If provided, FeatureTypes must have at least one 'text' string (for example, ['text']). If FeatureTypes is not provided, the explainer infers the feature types based on the baseline data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#feature_types SagemakerEndpointConfigA#feature_types}
	FeatureTypes *[]*string `field:"optional" json:"featureTypes" yaml:"featureTypes"`
	// A JMESPath expression used to locate the list of label headers in the model container output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#label_attribute SagemakerEndpointConfigA#label_attribute}
	LabelAttribute *string `field:"optional" json:"labelAttribute" yaml:"labelAttribute"`
	// For multiclass classification problems, the label headers are the names of the classes.
	//
	// Otherwise, the label header is the name of the predicted label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#label_headers SagemakerEndpointConfigA#label_headers}
	LabelHeaders *[]*string `field:"optional" json:"labelHeaders" yaml:"labelHeaders"`
	// A zero-based index used to extract a label header or list of label headers from model container output in CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#label_index SagemakerEndpointConfigA#label_index}
	LabelIndex *float64 `field:"optional" json:"labelIndex" yaml:"labelIndex"`
	// The maximum payload size (MB) allowed of a request from the explainer to the model container.
	//
	// Defaults to 6 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#max_payload_in_mb SagemakerEndpointConfigA#max_payload_in_mb}
	MaxPayloadInMb *float64 `field:"optional" json:"maxPayloadInMb" yaml:"maxPayloadInMb"`
	// The maximum number of records in a request that the model container can process when querying the model container for the predictions of a synthetic dataset.
	//
	// A record is a unit of input data that inference can be made on, for example, a single line in CSV data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#max_record_count SagemakerEndpointConfigA#max_record_count}
	MaxRecordCount *float64 `field:"optional" json:"maxRecordCount" yaml:"maxRecordCount"`
	// A JMESPath expression used to extract the probability (or score) from the model container output if the model container is in JSON Lines format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#probability_attribute SagemakerEndpointConfigA#probability_attribute}
	ProbabilityAttribute *string `field:"optional" json:"probabilityAttribute" yaml:"probabilityAttribute"`
	// A zero-based index used to extract a probability value (score) or list from model container output in CSV format.
	//
	// If this value is not provided, the entire model container output will be treated as a probability value (score) or list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#probability_index SagemakerEndpointConfigA#probability_index}
	ProbabilityIndex *float64 `field:"optional" json:"probabilityIndex" yaml:"probabilityIndex"`
}

