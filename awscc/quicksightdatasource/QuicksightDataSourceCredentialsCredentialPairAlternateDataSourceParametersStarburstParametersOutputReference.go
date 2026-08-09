// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
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
	DatabaseAccessControlRole() *string
	SetDatabaseAccessControlRole(val *string)
	DatabaseAccessControlRoleInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OAuthParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference
	OAuthParametersInput() interface{}
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProductType() *string
	SetProductType(val *string)
	ProductTypeInput() *string
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
	PutOAuthParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParameters)
	ResetAuthenticationType()
	ResetCatalog()
	ResetDatabaseAccessControlRole()
	ResetHost()
	ResetOAuthParameters()
	ResetPort()
	ResetProductType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) DatabaseAccessControlRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseAccessControlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) DatabaseAccessControlRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseAccessControlRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) OAuthParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference
	_jsii_.Get(
		j,
		"oAuthParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) OAuthParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuthParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ProductType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ProductTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference_Override(q QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetDatabaseAccessControlRole(val *string) {
	if err := j.validateSetDatabaseAccessControlRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseAccessControlRole",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetProductType(val *string) {
	if err := j.validateSetProductTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productType",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) PutOAuthParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParameters) {
	if err := q.validatePutOAuthParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOAuthParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		q,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		q,
		"resetCatalog",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetDatabaseAccessControlRole() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabaseAccessControlRole",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		q,
		"resetHost",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetOAuthParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetOAuthParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		q,
		"resetPort",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ResetProductType() {
	_jsii_.InvokeVoid(
		q,
		"resetProductType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

