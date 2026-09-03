// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigDataCaptureConfigOutputReference interface {
	cdktn.ComplexObject
	CaptureContentTypeHeader() SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeaderOutputReference
	CaptureContentTypeHeaderInput() interface{}
	CaptureOptions() SagemakerEndpointConfigDataCaptureConfigCaptureOptionsList
	CaptureOptionsInput() interface{}
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
	DestinationS3Uri() *string
	SetDestinationS3Uri(val *string)
	DestinationS3UriInput() *string
	EnableCapture() interface{}
	SetEnableCapture(val interface{})
	EnableCaptureInput() interface{}
	// Experimental.
	Fqn() *string
	InitialSamplingPercentage() *float64
	SetInitialSamplingPercentage(val *float64)
	InitialSamplingPercentageInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
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
	PutCaptureContentTypeHeader(value *SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeader)
	PutCaptureOptions(value interface{})
	ResetCaptureContentTypeHeader()
	ResetCaptureOptions()
	ResetDestinationS3Uri()
	ResetEnableCapture()
	ResetInitialSamplingPercentage()
	ResetKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigDataCaptureConfigOutputReference
type jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) CaptureContentTypeHeader() SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeaderOutputReference {
	var returns SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeaderOutputReference
	_jsii_.Get(
		j,
		"captureContentTypeHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) CaptureContentTypeHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureContentTypeHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) CaptureOptions() SagemakerEndpointConfigDataCaptureConfigCaptureOptionsList {
	var returns SagemakerEndpointConfigDataCaptureConfigCaptureOptionsList
	_jsii_.Get(
		j,
		"captureOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) CaptureOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) DestinationS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) DestinationS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) EnableCapture() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCapture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) EnableCaptureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCaptureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) InitialSamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSamplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) InitialSamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSamplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigDataCaptureConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigDataCaptureConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigDataCaptureConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigDataCaptureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigDataCaptureConfigOutputReference_Override(s SagemakerEndpointConfigDataCaptureConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigDataCaptureConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetDestinationS3Uri(val *string) {
	if err := j.validateSetDestinationS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationS3Uri",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetEnableCapture(val interface{}) {
	if err := j.validateSetEnableCaptureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCapture",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetInitialSamplingPercentage(val *float64) {
	if err := j.validateSetInitialSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialSamplingPercentage",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) PutCaptureContentTypeHeader(value *SagemakerEndpointConfigDataCaptureConfigCaptureContentTypeHeader) {
	if err := s.validatePutCaptureContentTypeHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCaptureContentTypeHeader",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) PutCaptureOptions(value interface{}) {
	if err := s.validatePutCaptureOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCaptureOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetCaptureContentTypeHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetCaptureContentTypeHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetCaptureOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetCaptureOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetDestinationS3Uri() {
	_jsii_.InvokeVoid(
		s,
		"resetDestinationS3Uri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetEnableCapture() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableCapture",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetInitialSamplingPercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialSamplingPercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigDataCaptureConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

