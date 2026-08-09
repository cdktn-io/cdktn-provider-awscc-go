// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pinpointinapptemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PinpointInAppTemplateContentPrimaryBtnOutputReference interface {
	cdktn.ComplexObject
	Android() PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference
	AndroidInput() interface{}
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
	DefaultConfig() PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference
	DefaultConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ios() PinpointInAppTemplateContentPrimaryBtnIosOutputReference
	IosInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Web() PinpointInAppTemplateContentPrimaryBtnWebOutputReference
	WebInput() interface{}
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
	PutAndroid(value *PinpointInAppTemplateContentPrimaryBtnAndroid)
	PutDefaultConfig(value *PinpointInAppTemplateContentPrimaryBtnDefaultConfig)
	PutIos(value *PinpointInAppTemplateContentPrimaryBtnIos)
	PutWeb(value *PinpointInAppTemplateContentPrimaryBtnWeb)
	ResetAndroid()
	ResetDefaultConfig()
	ResetIos()
	ResetWeb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PinpointInAppTemplateContentPrimaryBtnOutputReference
type jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) Android() PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference {
	var returns PinpointInAppTemplateContentPrimaryBtnAndroidOutputReference
	_jsii_.Get(
		j,
		"android",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) AndroidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"androidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) DefaultConfig() PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference {
	var returns PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference
	_jsii_.Get(
		j,
		"defaultConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) DefaultConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) Ios() PinpointInAppTemplateContentPrimaryBtnIosOutputReference {
	var returns PinpointInAppTemplateContentPrimaryBtnIosOutputReference
	_jsii_.Get(
		j,
		"ios",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) IosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) Web() PinpointInAppTemplateContentPrimaryBtnWebOutputReference {
	var returns PinpointInAppTemplateContentPrimaryBtnWebOutputReference
	_jsii_.Get(
		j,
		"web",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) WebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webInput",
		&returns,
	)
	return returns
}


func NewPinpointInAppTemplateContentPrimaryBtnOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PinpointInAppTemplateContentPrimaryBtnOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointInAppTemplateContentPrimaryBtnOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPinpointInAppTemplateContentPrimaryBtnOutputReference_Override(p PinpointInAppTemplateContentPrimaryBtnOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) PutAndroid(value *PinpointInAppTemplateContentPrimaryBtnAndroid) {
	if err := p.validatePutAndroidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAndroid",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) PutDefaultConfig(value *PinpointInAppTemplateContentPrimaryBtnDefaultConfig) {
	if err := p.validatePutDefaultConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDefaultConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) PutIos(value *PinpointInAppTemplateContentPrimaryBtnIos) {
	if err := p.validatePutIosParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIos",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) PutWeb(value *PinpointInAppTemplateContentPrimaryBtnWeb) {
	if err := p.validatePutWebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWeb",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ResetAndroid() {
	_jsii_.InvokeVoid(
		p,
		"resetAndroid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ResetDefaultConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ResetIos() {
	_jsii_.InvokeVoid(
		p,
		"resetIos",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ResetWeb() {
	_jsii_.InvokeVoid(
		p,
		"resetWeb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

