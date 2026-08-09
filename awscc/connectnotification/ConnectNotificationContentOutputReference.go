// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectnotification

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectnotification/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectNotificationContentOutputReference interface {
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
	DeDe() *string
	SetDeDe(val *string)
	DeDeInput() *string
	EnUs() *string
	SetEnUs(val *string)
	EnUsInput() *string
	EsEs() *string
	SetEsEs(val *string)
	EsEsInput() *string
	// Experimental.
	Fqn() *string
	FrFr() *string
	SetFrFr(val *string)
	FrFrInput() *string
	IdId() *string
	SetIdId(val *string)
	IdIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ItIt() *string
	SetItIt(val *string)
	ItItInput() *string
	JaJp() *string
	SetJaJp(val *string)
	JaJpInput() *string
	KoKr() *string
	SetKoKr(val *string)
	KoKrInput() *string
	PtBr() *string
	SetPtBr(val *string)
	PtBrInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ZhCn() *string
	SetZhCn(val *string)
	ZhCnInput() *string
	ZhTw() *string
	SetZhTw(val *string)
	ZhTwInput() *string
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
	ResetDeDe()
	ResetEnUs()
	ResetEsEs()
	ResetFrFr()
	ResetIdId()
	ResetItIt()
	ResetJaJp()
	ResetKoKr()
	ResetPtBr()
	ResetZhCn()
	ResetZhTw()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectNotificationContentOutputReference
type jsiiProxy_ConnectNotificationContentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) DeDe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deDe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) DeDeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deDeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) EnUs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enUs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) EnUsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enUsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) EsEs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esEs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) EsEsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esEsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) FrFr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frFr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) FrFrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frFrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) IdId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) IdIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ItIt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itIt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ItItInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itItInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) JaJp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaJp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) JaJpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jaJpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) KoKr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"koKr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) KoKrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"koKrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) PtBr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ptBr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) PtBrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ptBrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ZhCn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zhCn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ZhCnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zhCnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ZhTw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zhTw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference) ZhTwInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zhTwInput",
		&returns,
	)
	return returns
}


func NewConnectNotificationContentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectNotificationContentOutputReference {
	_init_.Initialize()

	if err := validateNewConnectNotificationContentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectNotificationContentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectNotification.ConnectNotificationContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectNotificationContentOutputReference_Override(c ConnectNotificationContentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectNotification.ConnectNotificationContentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetDeDe(val *string) {
	if err := j.validateSetDeDeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deDe",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetEnUs(val *string) {
	if err := j.validateSetEnUsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enUs",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetEsEs(val *string) {
	if err := j.validateSetEsEsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"esEs",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetFrFr(val *string) {
	if err := j.validateSetFrFrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frFr",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetIdId(val *string) {
	if err := j.validateSetIdIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idId",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetItIt(val *string) {
	if err := j.validateSetItItParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"itIt",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetJaJp(val *string) {
	if err := j.validateSetJaJpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jaJp",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetKoKr(val *string) {
	if err := j.validateSetKoKrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"koKr",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetPtBr(val *string) {
	if err := j.validateSetPtBrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ptBr",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetZhCn(val *string) {
	if err := j.validateSetZhCnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zhCn",
		val,
	)
}

func (j *jsiiProxy_ConnectNotificationContentOutputReference)SetZhTw(val *string) {
	if err := j.validateSetZhTwParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zhTw",
		val,
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetDeDe() {
	_jsii_.InvokeVoid(
		c,
		"resetDeDe",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetEnUs() {
	_jsii_.InvokeVoid(
		c,
		"resetEnUs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetEsEs() {
	_jsii_.InvokeVoid(
		c,
		"resetEsEs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetFrFr() {
	_jsii_.InvokeVoid(
		c,
		"resetFrFr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetIdId() {
	_jsii_.InvokeVoid(
		c,
		"resetIdId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetItIt() {
	_jsii_.InvokeVoid(
		c,
		"resetItIt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetJaJp() {
	_jsii_.InvokeVoid(
		c,
		"resetJaJp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetKoKr() {
	_jsii_.InvokeVoid(
		c,
		"resetKoKr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetPtBr() {
	_jsii_.InvokeVoid(
		c,
		"resetPtBr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetZhCn() {
	_jsii_.InvokeVoid(
		c,
		"resetZhCn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ResetZhTw() {
	_jsii_.InvokeVoid(
		c,
		"resetZhTw",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectNotificationContentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

