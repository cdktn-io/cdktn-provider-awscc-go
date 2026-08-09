// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricoutboundexternallink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rtbfabricoutboundexternallink/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference interface {
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

// The jsii proxy struct for RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference
type jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) HttpCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) HttpCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) LoggingTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loggingTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) LoggingTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loggingTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResponseLoggingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseLoggingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResponseLoggingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseLoggingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference {
	_init_.Initialize()

	if err := validateNewRtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricOutboundExternalLink.RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference_Override(r RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rtbfabricOutboundExternalLink.RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetHttpCode(val *string) {
	if err := j.validateSetHttpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpCode",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetLoggingTypes(val *[]*string) {
	if err := j.validateSetLoggingTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingTypes",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetResponseLoggingPercentage(val *float64) {
	if err := j.validateSetResponseLoggingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseLoggingPercentage",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		r,
		"resetAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResetHttpCode() {
	_jsii_.InvokeVoid(
		r,
		"resetHttpCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResetLoggingTypes() {
	_jsii_.InvokeVoid(
		r,
		"resetLoggingTypes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ResetResponseLoggingPercentage() {
	_jsii_.InvokeVoid(
		r,
		"resetResponseLoggingPercentage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_RtbfabricOutboundExternalLinkLinkAttributesResponderErrorMaskingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

