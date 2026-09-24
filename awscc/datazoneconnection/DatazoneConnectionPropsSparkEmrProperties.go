// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsSparkEmrProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#compute_arn DatazoneConnection#compute_arn}.
	ComputeArn *string `field:"optional" json:"computeArn" yaml:"computeArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#instance_profile_arn DatazoneConnection#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#java_virtual_env DatazoneConnection#java_virtual_env}.
	JavaVirtualEnv *string `field:"optional" json:"javaVirtualEnv" yaml:"javaVirtualEnv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#log_uri DatazoneConnection#log_uri}.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#managed_endpoint_arn DatazoneConnection#managed_endpoint_arn}.
	ManagedEndpointArn *string `field:"optional" json:"managedEndpointArn" yaml:"managedEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#python_virtual_env DatazoneConnection#python_virtual_env}.
	PythonVirtualEnv *string `field:"optional" json:"pythonVirtualEnv" yaml:"pythonVirtualEnv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#runtime_role DatazoneConnection#runtime_role}.
	RuntimeRole *string `field:"optional" json:"runtimeRole" yaml:"runtimeRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#trusted_certificates_s3_uri DatazoneConnection#trusted_certificates_s3_uri}.
	TrustedCertificatesS3Uri *string `field:"optional" json:"trustedCertificatesS3Uri" yaml:"trustedCertificatesS3Uri"`
}

