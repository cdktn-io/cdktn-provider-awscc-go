// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorCsvFormatDescriptor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#charset LookoutmetricsAnomalyDetector#charset}.
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#contains_header LookoutmetricsAnomalyDetector#contains_header}.
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#delimiter LookoutmetricsAnomalyDetector#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#file_compression LookoutmetricsAnomalyDetector#file_compression}.
	FileCompression *string `field:"optional" json:"fileCompression" yaml:"fileCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#header_list LookoutmetricsAnomalyDetector#header_list}.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lookoutmetrics_anomaly_detector#quote_symbol LookoutmetricsAnomalyDetector#quote_symbol}.
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
}

