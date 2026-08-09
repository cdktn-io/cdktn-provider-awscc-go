// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabriclink/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	HttpCode() *string
	SetHttpCode(val *string)
	HttpCodeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoggingTypes() *[]*string
	SetLoggingTypes(val *[]*string)
	LoggingTypesInput() *[]*string
	ResponseLoggingPercentage() *float64
	SetResponseLoggingPercentage(val *float64)
	ResponseLoggingPercentageInput() *float64
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
	ResetAction()
	ResetHttpCode()
	ResetLoggingTypes()
	ResetResponseLoggingPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference
type jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) HttpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) HttpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) LoggingTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loggingTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) LoggingTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loggingTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResponseLoggingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseLoggingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResponseLoggingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseLoggingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricLinkLinkAttributesResponderErrorMaskingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference_Override(r RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricLink.RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetHttpCode(val *string) {
	if err := j.validateSetHttpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCode",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetLoggingTypes(val *[]*string) {
	if err := j.validateSetLoggingTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingTypes",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetResponseLoggingPercentage(val *float64) {
	if err := j.validateSetResponseLoggingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseLoggingPercentage",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		r,
		"resetAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResetHttpCode() {
	_jsii_.InvokeVoid(
		r,
		"resetHttpCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResetLoggingTypes() {
	_jsii_.InvokeVoid(
		r,
		"resetLoggingTypes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ResetResponseLoggingPercentage() {
	_jsii_.InvokeVoid(
		r,
		"resetResponseLoggingPercentage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricLinkLinkAttributesResponderErrorMaskingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

