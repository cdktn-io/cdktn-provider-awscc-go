// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitouserpoolregionalconfigurationattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference interface {
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
	EumsSms() CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference
	EumsSmsInput() interface{}
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SnsCallerArn() *string
	SetSnsCallerArn(val *string)
	SnsCallerArnInput() *string
	SnsRegion() *string
	SetSnsRegion(val *string)
	SnsRegionInput() *string
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
	PutEumsSms(value *CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSms)
	ResetEumsSms()
	ResetExternalId()
	ResetSnsCallerArn()
	ResetSnsRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference
type jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) EumsSms() CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference {
	var returns CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSmsOutputReference
	_jsii_.Get(
		j,
		"eumsSms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) EumsSmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eumsSmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) SnsCallerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsCallerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) SnsCallerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsCallerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) SnsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) SnsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference_Override(c CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetSnsCallerArn(val *string) {
	if err := j.validateSetSnsCallerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsCallerArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetSnsRegion(val *string) {
	if err := j.validateSetSnsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsRegion",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) PutEumsSms(value *CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationEumsSms) {
	if err := c.validatePutEumsSmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEumsSms",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ResetEumsSms() {
	_jsii_.InvokeVoid(
		c,
		"resetEumsSms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ResetSnsCallerArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsCallerArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ResetSnsRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetSnsRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentSmsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

