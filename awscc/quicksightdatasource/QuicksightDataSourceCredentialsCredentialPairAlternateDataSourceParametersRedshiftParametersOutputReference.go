// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference interface {
	cdktn.ComplexObject
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IamParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference
	IamParametersInput() interface{}
	IdentityCenterConfiguration() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfigurationOutputReference
	IdentityCenterConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	PutIamParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParameters)
	PutIdentityCenterConfiguration(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfiguration)
	ResetClusterId()
	ResetDatabase()
	ResetHost()
	ResetIamParameters()
	ResetIdentityCenterConfiguration()
	ResetPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) IamParameters() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference
	_jsii_.Get(
		j,
		"iamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) IamParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) IdentityCenterConfiguration() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfigurationOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfigurationOutputReference
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) IdentityCenterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityCenterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference_Override(q QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) PutIamParameters(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIamParameters) {
	if err := q.validatePutIamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIamParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) PutIdentityCenterConfiguration(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersIdentityCenterConfiguration) {
	if err := q.validatePutIdentityCenterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIdentityCenterConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		q,
		"resetClusterId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		q,
		"resetHost",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetIamParameters() {
	_jsii_.InvokeVoid(
		q,
		"resetIamParameters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetIdentityCenterConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityCenterConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		q,
		"resetPort",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersRedshiftParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

