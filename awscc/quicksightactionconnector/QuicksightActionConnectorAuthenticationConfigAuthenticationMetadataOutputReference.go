// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightactionconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightactionconnector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadataOutputReference
	ApiKeyConnectionMetadataInput() interface{}
	AuthorizationCodeGrantMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadataOutputReference
	AuthorizationCodeGrantMetadataInput() interface{}
	BasicAuthConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadataOutputReference
	BasicAuthConnectionMetadataInput() interface{}
	ClientCredentialsGrantMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadataOutputReference
	ClientCredentialsGrantMetadataInput() interface{}
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
	IamConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadataOutputReference
	IamConnectionMetadataInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoneConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadataOutputReference
	NoneConnectionMetadataInput() interface{}
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
	PutApiKeyConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadata)
	PutAuthorizationCodeGrantMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadata)
	PutBasicAuthConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadata)
	PutClientCredentialsGrantMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadata)
	PutIamConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadata)
	PutNoneConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadata)
	ResetApiKeyConnectionMetadata()
	ResetAuthorizationCodeGrantMetadata()
	ResetBasicAuthConnectionMetadata()
	ResetClientCredentialsGrantMetadata()
	ResetIamConnectionMetadata()
	ResetNoneConnectionMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference
type jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ApiKeyConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadataOutputReference
	_jsii_.Get(
		j,
		"apiKeyConnectionMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ApiKeyConnectionMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyConnectionMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) AuthorizationCodeGrantMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadataOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeGrantMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) AuthorizationCodeGrantMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationCodeGrantMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) BasicAuthConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadataOutputReference
	_jsii_.Get(
		j,
		"basicAuthConnectionMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) BasicAuthConnectionMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConnectionMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ClientCredentialsGrantMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadataOutputReference
	_jsii_.Get(
		j,
		"clientCredentialsGrantMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ClientCredentialsGrantMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCredentialsGrantMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) IamConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadataOutputReference
	_jsii_.Get(
		j,
		"iamConnectionMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) IamConnectionMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamConnectionMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) NoneConnectionMetadata() QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadataOutputReference {
	var returns QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadataOutputReference
	_jsii_.Get(
		j,
		"noneConnectionMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) NoneConnectionMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noneConnectionMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightActionConnector.QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference_Override(q QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightActionConnector.QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutApiKeyConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataApiKeyConnectionMetadata) {
	if err := q.validatePutApiKeyConnectionMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putApiKeyConnectionMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutAuthorizationCodeGrantMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataAuthorizationCodeGrantMetadata) {
	if err := q.validatePutAuthorizationCodeGrantMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAuthorizationCodeGrantMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutBasicAuthConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataBasicAuthConnectionMetadata) {
	if err := q.validatePutBasicAuthConnectionMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putBasicAuthConnectionMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutClientCredentialsGrantMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataClientCredentialsGrantMetadata) {
	if err := q.validatePutClientCredentialsGrantMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putClientCredentialsGrantMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutIamConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataIamConnectionMetadata) {
	if err := q.validatePutIamConnectionMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIamConnectionMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) PutNoneConnectionMetadata(value *QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataNoneConnectionMetadata) {
	if err := q.validatePutNoneConnectionMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNoneConnectionMetadata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetApiKeyConnectionMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetApiKeyConnectionMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetAuthorizationCodeGrantMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetAuthorizationCodeGrantMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetBasicAuthConnectionMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetBasicAuthConnectionMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetClientCredentialsGrantMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetClientCredentialsGrantMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetIamConnectionMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetIamConnectionMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ResetNoneConnectionMetadata() {
	_jsii_.InvokeVoid(
		q,
		"resetNoneConnectionMetadata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightActionConnectorAuthenticationConfigAuthenticationMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

