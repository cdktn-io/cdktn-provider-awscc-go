// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLifecycleConfiguration struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Indicates which default minimum object size behavior is applied to the lifecycle configuration.
	//
	// This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.
	//    +   ``all_storage_classes_128K`` - Objects smaller than 128 KB will not transition to any storage class by default.
	//   +   ``varies_by_storage_class`` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB.
	//
	//  To customize the minimum object size for any transition you can add a filter that specifies a custom ``ObjectSizeGreaterThan`` or ``ObjectSizeLessThan`` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3_bucket#transition_default_minimum_object_size S3Bucket#transition_default_minimum_object_size}
	TransitionDefaultMinimumObjectSize *string `field:"optional" json:"transitionDefaultMinimumObjectSize" yaml:"transitionDefaultMinimumObjectSize"`
}

