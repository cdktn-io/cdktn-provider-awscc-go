// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformInputRecordTables struct {
	// The database and table in the AWS Glue Data Catalog that is used for input or output data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_ml_transform#glue_tables GlueMlTransform#glue_tables}
	GlueTables interface{} `field:"optional" json:"glueTables" yaml:"glueTables"`
}

