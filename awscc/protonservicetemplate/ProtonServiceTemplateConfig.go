// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package protonservicetemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ProtonServiceTemplateConfig struct {
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
	// <p>A description of the service template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#description ProtonServiceTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// <p>The name of the service template as displayed in the developer interface.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#display_name ProtonServiceTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// <p>A customer provided encryption key that's used to encrypt data.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#encryption_key ProtonServiceTemplate#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#name ProtonServiceTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#pipeline_provisioning ProtonServiceTemplate#pipeline_provisioning}.
	PipelineProvisioning *string `field:"optional" json:"pipelineProvisioning" yaml:"pipelineProvisioning"`
	// <p>An optional list of metadata items that you can associate with the Proton service template.
	//
	// A tag is a key-value pair.</p>
	//          <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
	//         <i>Proton User Guide</i>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/proton_service_template#tags ProtonServiceTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

