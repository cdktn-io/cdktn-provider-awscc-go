// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	Database() *string
	SetDatabase(val *string)
	DatabaseAccessControlRole() *string
	SetDatabaseAccessControlRole(val *string)
	DatabaseAccessControlRoleInput() *string
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OAuthParameters() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference
	OAuthParametersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
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
	PutOAuthParameters(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParameters)
	ResetAuthenticationType()
	ResetDatabase()
	ResetDatabaseAccessControlRole()
	ResetHost()
	ResetOAuthParameters()
	ResetWarehouse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference
type jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) DatabaseAccessControlRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseAccessControlRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) DatabaseAccessControlRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseAccessControlRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) OAuthParameters() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference
	_jsii_.Get(
		j,
		"oAuthParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) OAuthParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuthParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference_Override(q QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetDatabaseAccessControlRole(val *string) {
	if err := j.validateSetDatabaseAccessControlRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseAccessControlRole",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference)SetWarehouse(val *string) {
	if err := j.validateSetWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) PutOAuthParameters(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParameters) {
	if err := q.validatePutOAuthParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putOAuthParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		q,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetDatabaseAccessControlRole() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabaseAccessControlRole",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		q,
		"resetHost",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetOAuthParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetOAuthParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ResetWarehouse() {
	_jsii_.InvokeVoid(
		q,
		"resetWarehouse",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

