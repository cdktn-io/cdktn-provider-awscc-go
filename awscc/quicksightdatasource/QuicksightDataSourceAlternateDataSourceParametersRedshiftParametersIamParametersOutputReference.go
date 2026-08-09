// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference interface {
	cdktn.ComplexObject
	AutoCreateDatabaseUser() interface{}
	SetAutoCreateDatabaseUser(val interface{})
	AutoCreateDatabaseUserInput() interface{}
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
	DatabaseGroups() *[]*string
	SetDatabaseGroups(val *[]*string)
	DatabaseGroupsInput() *[]*string
	DatabaseUser() *string
	SetDatabaseUser(val *string)
	DatabaseUserInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetAutoCreateDatabaseUser()
	ResetDatabaseGroups()
	ResetDatabaseUser()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference
type jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) AutoCreateDatabaseUser() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateDatabaseUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) AutoCreateDatabaseUserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateDatabaseUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) DatabaseGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) DatabaseGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) DatabaseUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) DatabaseUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference_Override(q QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetAutoCreateDatabaseUser(val interface{}) {
	if err := j.validateSetAutoCreateDatabaseUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreateDatabaseUser",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetDatabaseGroups(val *[]*string) {
	if err := j.validateSetDatabaseGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseGroups",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetDatabaseUser(val *string) {
	if err := j.validateSetDatabaseUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseUser",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ResetAutoCreateDatabaseUser() {
	_jsii_.InvokeVoid(
		q,
		"resetAutoCreateDatabaseUser",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ResetDatabaseGroups() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabaseGroups",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ResetDatabaseUser() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabaseUser",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		q,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersRedshiftParametersIamParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

