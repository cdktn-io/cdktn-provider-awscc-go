// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lookoutequipmentinferencescheduler/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InferenceInputNameConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference
	InferenceInputNameConfigurationInput() interface{}
	InputTimeZoneOffset() *string
	SetInputTimeZoneOffset(val *string)
	InputTimeZoneOffsetInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3InputConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfigurationOutputReference
	S3InputConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutInferenceInputNameConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfiguration)
	PutS3InputConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfiguration)
	ResetInferenceInputNameConfiguration()
	ResetInputTimeZoneOffset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference
type jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InferenceInputNameConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference {
	var returns LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfigurationOutputReference
	_jsii_.Get(
		j,
		"inferenceInputNameConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InferenceInputNameConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceInputNameConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InputTimeZoneOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTimeZoneOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InputTimeZoneOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTimeZoneOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) S3InputConfiguration() LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfigurationOutputReference {
	var returns LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3InputConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) S3InputConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3InputConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLookoutequipmentInferenceSchedulerDataInputConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference_Override(l LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lookoutequipmentInferenceScheduler.LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetInputTimeZoneOffset(val *string) {
	if err := j.validateSetInputTimeZoneOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputTimeZoneOffset",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) PutInferenceInputNameConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfiguration) {
	if err := l.validatePutInferenceInputNameConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInferenceInputNameConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) PutS3InputConfiguration(value *LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfiguration) {
	if err := l.validatePutS3InputConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putS3InputConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ResetInferenceInputNameConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetInferenceInputNameConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ResetInputTimeZoneOffset() {
	_jsii_.InvokeVoid(
		l,
		"resetInputTimeZoneOffset",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookoutequipmentInferenceSchedulerDataInputConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

