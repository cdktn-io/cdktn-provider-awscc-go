// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionCode struct {
	// URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#image_uri LambdaFunction#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// An Amazon S3 bucket in the same AWS-Region as your function. The bucket can be in a different AWS-account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#s3_bucket LambdaFunction#s3_bucket}
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the deployment package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#s3_key LambdaFunction#s3_key}
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#s3_object_storage_mode LambdaFunction#s3_object_storage_mode}.
	S3ObjectStorageMode *string `field:"optional" json:"s3ObjectStorageMode" yaml:"s3ObjectStorageMode"`
	// For versioned objects, the version of the deployment package object to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#s3_object_version LambdaFunction#s3_object_version}
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// The ARN of the KMSlong (KMS) customer managed key that's used to encrypt your function's .zip deployment package. If you don't provide a customer managed key, Lambda uses an [owned key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#source_kms_key_arn LambdaFunction#source_kms_key_arn}
	SourceKmsKeyArn *string `field:"optional" json:"sourceKmsKeyArn" yaml:"sourceKmsKeyArn"`
	// (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, CFN places it in a file named ``index`` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html). This zip file cannot exceed 4MB. For the ``Handler`` property, the first part of the handler identifier must be ``index``. For example, ``index.handler``.   When you specify source code inline for a Node.js function, the ``index`` file that CFN creates uses the extension ``.js``. This means that Node.js treats the file as a CommonJS module.  When using Node.js 24 or later, Node.js can automatically detect if a ``.js`` file should be treated as CommonJS or as an ES module. To enable auto-detection, add the ``--experimental-detect-module`` flag to the ``NODE_OPTIONS`` environment variable. For more information, see [Experimental Node.js features](https://docs.aws.amazon.com//lambda/latest/dg/lambda-nodejs.html#nodejs-experimental-features).    For JSON, you must escape quotes and special characters such as newline (``\n``) with a backslash.  If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ([cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html)) that simplifies sending responses. See [Using Lambda with CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function#zip_file LambdaFunction#zip_file}
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

