// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph


type NeptunegraphGraphImportTaskImportOptionsNeptune struct {
	// Neptune Analytics supports label-less vertices and no labels are assigned unless one is explicitly provided.
	//
	// Neptune assigns default labels when none is explicitly provided. When importing the data into Neptune Analytics, the default vertex labels can be omitted by setting preserveDefaultVertexLabels to false. Note that if the vertex only has default labels, and has no other properties or edges, then the vertex will effectively not get imported into Neptune Analytics when preserveDefaultVertexLabels is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#preserve_default_vertex_labels NeptunegraphGraph#preserve_default_vertex_labels}
	PreserveDefaultVertexLabels interface{} `field:"optional" json:"preserveDefaultVertexLabels" yaml:"preserveDefaultVertexLabels"`
	// Neptune Analytics currently does not support user defined edge ids.
	//
	// The edge ids are not imported by default. They are imported if preserveEdgeIds is set to true, and ids are stored as properties on the relationships with the property name neptuneEdgeId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#preserve_edge_ids NeptunegraphGraph#preserve_edge_ids}
	PreserveEdgeIds interface{} `field:"optional" json:"preserveEdgeIds" yaml:"preserveEdgeIds"`
	// The KMS key to use to encrypt data in the S3 bucket where the graph data is exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#s3_export_kms_key_id NeptunegraphGraph#s3_export_kms_key_id}
	S3ExportKmsKeyId *string `field:"optional" json:"s3ExportKmsKeyId" yaml:"s3ExportKmsKeyId"`
	// The path to an S3 bucket from which to import data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptunegraph_graph#s3_export_path NeptunegraphGraph#s3_export_path}
	S3ExportPath *string `field:"optional" json:"s3ExportPath" yaml:"s3ExportPath"`
}

