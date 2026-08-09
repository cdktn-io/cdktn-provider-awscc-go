// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pinpointinapptemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PinpointInAppTemplateContentSecondaryBtnOutputReference interface {
	cdktn.ComplexObject
	Android() PinpointInAppTemplateContentSecondaryBtnAndroidOutputReference
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
	DefaultConfig() PinpointInAppTemplateContentSecondaryBtnDefaultConfigOutputReference
	DefaultConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ios() PinpointInAppTemplateContentSecondaryBtnIosOutputReference
	IosInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Web() PinpointInAppTemplateContentSecondaryBtnWebOutputReference
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
	PutAndroid(value *PinpointInAppTemplateContentSecondaryBtnAndroid)
	PutDefaultConfig(value *PinpointInAppTemplateContentSecondaryBtnDefaultConfig)
	PutIos(value *PinpointInAppTemplateContentSecondaryBtnIos)
	PutWeb(value *PinpointInAppTemplateContentSecondaryBtnWeb)
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

// The jsii proxy struct for PinpointInAppTemplateContentSecondaryBtnOutputReference
type jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) Android() PinpointInAppTemplateContentSecondaryBtnAndroidOutputReference {
	var returns PinpointInAppTemplateContentSecondaryBtnAndroidOutputReference
	_jsii_.Get(
		j,
		"android",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) AndroidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"androidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) DefaultConfig() PinpointInAppTemplateContentSecondaryBtnDefaultConfigOutputReference {
	var returns PinpointInAppTemplateContentSecondaryBtnDefaultConfigOutputReference
	_jsii_.Get(
		j,
		"defaultConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) DefaultConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) Ios() PinpointInAppTemplateContentSecondaryBtnIosOutputReference {
	var returns PinpointInAppTemplateContentSecondaryBtnIosOutputReference
	_jsii_.Get(
		j,
		"ios",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) IosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) Web() PinpointInAppTemplateContentSecondaryBtnWebOutputReference {
	var returns PinpointInAppTemplateContentSecondaryBtnWebOutputReference
	_jsii_.Get(
		j,
		"web",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) WebInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webInput",
		&returns,
	)
	return returns
}


func NewPinpointInAppTemplateContentSecondaryBtnOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PinpointInAppTemplateContentSecondaryBtnOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointInAppTemplateContentSecondaryBtnOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentSecondaryBtnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPinpointInAppTemplateContentSecondaryBtnOutputReference_Override(p PinpointInAppTemplateContentSecondaryBtnOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentSecondaryBtnOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) PutAndroid(value *PinpointInAppTemplateContentSecondaryBtnAndroid) {
	if err := p.validatePutAndroidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAndroid",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) PutDefaultConfig(value *PinpointInAppTemplateContentSecondaryBtnDefaultConfig) {
	if err := p.validatePutDefaultConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDefaultConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) PutIos(value *PinpointInAppTemplateContentSecondaryBtnIos) {
	if err := p.validatePutIosParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIos",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) PutWeb(value *PinpointInAppTemplateContentSecondaryBtnWeb) {
	if err := p.validatePutWebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWeb",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ResetAndroid() {
	_jsii_.InvokeVoid(
		p,
		"resetAndroid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ResetDefaultConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ResetIos() {
	_jsii_.InvokeVoid(
		p,
		"resetIos",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ResetWeb() {
	_jsii_.InvokeVoid(
		p,
		"resetWeb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PinpointInAppTemplateContentSecondaryBtnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

