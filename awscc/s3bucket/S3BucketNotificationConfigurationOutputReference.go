// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3bucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3BucketNotificationConfigurationOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventBridgeConfiguration() S3BucketNotificationConfigurationEventBridgeConfigurationOutputReference
	EventBridgeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LambdaConfigurations() S3BucketNotificationConfigurationLambdaConfigurationsList
	LambdaConfigurationsInput() interface{}
	QueueConfigurations() S3BucketNotificationConfigurationQueueConfigurationsList
	QueueConfigurationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TopicConfigurations() S3BucketNotificationConfigurationTopicConfigurationsList
	TopicConfigurationsInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutEventBridgeConfiguration(value *S3BucketNotificationConfigurationEventBridgeConfiguration)
	PutLambdaConfigurations(value interface{})
	PutQueueConfigurations(value interface{})
	PutTopicConfigurations(value interface{})
	ResetEventBridgeConfiguration()
	ResetLambdaConfigurations()
	ResetQueueConfigurations()
	ResetTopicConfigurations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketNotificationConfigurationOutputReference
type jsiiProxy_S3BucketNotificationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) EventBridgeConfiguration() S3BucketNotificationConfigurationEventBridgeConfigurationOutputReference {
	var returns S3BucketNotificationConfigurationEventBridgeConfigurationOutputReference
	_jsii_.Get(
		j,
		"eventBridgeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) EventBridgeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventBridgeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) LambdaConfigurations() S3BucketNotificationConfigurationLambdaConfigurationsList {
	var returns S3BucketNotificationConfigurationLambdaConfigurationsList
	_jsii_.Get(
		j,
		"lambdaConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) LambdaConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) QueueConfigurations() S3BucketNotificationConfigurationQueueConfigurationsList {
	var returns S3BucketNotificationConfigurationQueueConfigurationsList
	_jsii_.Get(
		j,
		"queueConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) QueueConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) TopicConfigurations() S3BucketNotificationConfigurationTopicConfigurationsList {
	var returns S3BucketNotificationConfigurationTopicConfigurationsList
	_jsii_.Get(
		j,
		"topicConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference) TopicConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"topicConfigurationsInput",
		&returns,
	)
	return returns
}


func NewS3BucketNotificationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3BucketNotificationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketNotificationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketNotificationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3Bucket.S3BucketNotificationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketNotificationConfigurationOutputReference_Override(s S3BucketNotificationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3Bucket.S3BucketNotificationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketNotificationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) PutEventBridgeConfiguration(value *S3BucketNotificationConfigurationEventBridgeConfiguration) {
	if err := s.validatePutEventBridgeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventBridgeConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) PutLambdaConfigurations(value interface{}) {
	if err := s.validatePutLambdaConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLambdaConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) PutQueueConfigurations(value interface{}) {
	if err := s.validatePutQueueConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQueueConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) PutTopicConfigurations(value interface{}) {
	if err := s.validatePutTopicConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTopicConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ResetEventBridgeConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetEventBridgeConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ResetLambdaConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ResetQueueConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetQueueConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ResetTopicConfigurations() {
	_jsii_.InvokeVoid(
		s,
		"resetTopicConfigurations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketNotificationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

