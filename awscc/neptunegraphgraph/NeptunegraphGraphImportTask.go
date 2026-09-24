// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph


type NeptunegraphGraphImportTask struct {
	// The method to handle blank nodes in the dataset.
	//
	// Currently, only convertToIri is supported, meaning blank nodes are converted to unique IRIs at load time. Must be provided when format is NTRIPLES
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#blank_node_handling NeptunegraphGraph#blank_node_handling}
	BlankNodeHandling *string `field:"optional" json:"blankNodeHandling" yaml:"blankNodeHandling"`
	// If set to true, the task halts when an import error is encountered.
	//
	// If set to false, the task skips the data that caused the error and continues if possible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#fail_on_error NeptunegraphGraph#fail_on_error}
	FailOnError interface{} `field:"optional" json:"failOnError" yaml:"failOnError"`
	// Specifies the format of S3 data to be imported.
	//
	// Valid values are CSV, which identifies the Gremlin CSV format, OPEN_CYPHER, which identifies the openCypher load format, or NTRIPLES, which identifies the RDF n-triples format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#format NeptunegraphGraph#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Contains options for controlling the import process.
	//
	// For example, if the failOnError key is set to false, the import skips the data that caused the error and continues if possible (whereas if set to true, the default, or if omitted, the import operation halts immediately when an error is encountered).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#import_options NeptunegraphGraph#import_options}
	ImportOptions *NeptunegraphGraphImportTaskImportOptions `field:"optional" json:"importOptions" yaml:"importOptions"`
	// The maximum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Default: 1024, or the approved upper limit for your account. If both the minimum and maximum values are specified, the final provisioned-memory will be chosen per the actual size of your imported data. If neither value is specified, 128 m-NCUs are used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#max_provisioned_memory NeptunegraphGraph#max_provisioned_memory}
	MaxProvisionedMemory *float64 `field:"optional" json:"maxProvisionedMemory" yaml:"maxProvisionedMemory"`
	// The minimum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph. Default: 16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#min_provisioned_memory NeptunegraphGraph#min_provisioned_memory}
	MinProvisionedMemory *float64 `field:"optional" json:"minProvisionedMemory" yaml:"minProvisionedMemory"`
	// The parquet type of the import task. Required when Format is PARQUET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#parquet_type NeptunegraphGraph#parquet_type}
	ParquetType *string `field:"optional" json:"parquetType" yaml:"parquetType"`
	// The ARN of the IAM role that will allow access to the data that is to be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#role_arn NeptunegraphGraph#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A URL identifying to the location of the data to be imported.
	//
	// This can be an Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#source NeptunegraphGraph#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

