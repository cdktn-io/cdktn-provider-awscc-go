// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsingestconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IvsIngestConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Ingest Protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#ingest_protocol IvsIngestConfiguration#ingest_protocol}
	IngestProtocol *string `field:"optional" json:"ingestProtocol" yaml:"ingestProtocol"`
	// Whether ingest configuration allows insecure ingest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#insecure_ingest IvsIngestConfiguration#insecure_ingest}
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// IngestConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#name IvsIngestConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Stage ARN.
	//
	// A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: "" (recording is disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#stage_arn IvsIngestConfiguration#stage_arn}
	StageArn *string `field:"optional" json:"stageArn" yaml:"stageArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#tags IvsIngestConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// User defined indentifier for participant associated with IngestConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_ingest_configuration#user_id IvsIngestConfiguration#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

