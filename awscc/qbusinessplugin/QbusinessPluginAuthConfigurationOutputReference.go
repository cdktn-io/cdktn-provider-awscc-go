// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessplugin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinessplugin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessPluginAuthConfigurationOutputReference interface {
	cdktn.ComplexObject
	BasicAuthConfiguration() QbusinessPluginAuthConfigurationBasicAuthConfigurationOutputReference
	BasicAuthConfigurationInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoAuthConfiguration() *string
	SetNoAuthConfiguration(val *string)
	NoAuthConfigurationInput() *string
	OAuth2ClientCredentialConfiguration() QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfigurationOutputReference
	OAuth2ClientCredentialConfigurationInput() interface{}
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
	PutBasicAuthConfiguration(value *QbusinessPluginAuthConfigurationBasicAuthConfiguration)
	PutOAuth2ClientCredentialConfiguration(value *QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfiguration)
	ResetBasicAuthConfiguration()
	ResetNoAuthConfiguration()
	ResetOAuth2ClientCredentialConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessPluginAuthConfigurationOutputReference
type jsiiProxy_QbusinessPluginAuthConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) BasicAuthConfiguration() QbusinessPluginAuthConfigurationBasicAuthConfigurationOutputReference {
	var returns QbusinessPluginAuthConfigurationBasicAuthConfigurationOutputReference
	_jsii_.Get(
		j,
		"basicAuthConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) BasicAuthConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) NoAuthConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noAuthConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) NoAuthConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noAuthConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) OAuth2ClientCredentialConfiguration() QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfigurationOutputReference {
	var returns QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfigurationOutputReference
	_jsii_.Get(
		j,
		"oAuth2ClientCredentialConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) OAuth2ClientCredentialConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuth2ClientCredentialConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQbusinessPluginAuthConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QbusinessPluginAuthConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessPluginAuthConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessPluginAuthConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessPlugin.QbusinessPluginAuthConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessPluginAuthConfigurationOutputReference_Override(q QbusinessPluginAuthConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessPlugin.QbusinessPluginAuthConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetNoAuthConfiguration(val *string) {
	if err := j.validateSetNoAuthConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAuthConfiguration",
		val,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) PutBasicAuthConfiguration(value *QbusinessPluginAuthConfigurationBasicAuthConfiguration) {
	if err := q.validatePutBasicAuthConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putBasicAuthConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) PutOAuth2ClientCredentialConfiguration(value *QbusinessPluginAuthConfigurationOAuth2ClientCredentialConfiguration) {
	if err := q.validatePutOAuth2ClientCredentialConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOAuth2ClientCredentialConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ResetBasicAuthConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetBasicAuthConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ResetNoAuthConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetNoAuthConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ResetOAuth2ClientCredentialConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetOAuth2ClientCredentialConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessPluginAuthConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

