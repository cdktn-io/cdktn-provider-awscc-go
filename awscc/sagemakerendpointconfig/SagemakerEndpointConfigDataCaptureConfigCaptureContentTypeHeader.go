// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeader struct {
	// A list of the CSV content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#csv_content_types SagemakerEndpointConfigA#csv_content_types}
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// A list of the JSON content types of the data that the endpoint captures.
	//
	// For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#json_content_types SagemakerEndpointConfigA#json_content_types}
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

