// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxytargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsdbproxytargetgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference interface {
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
	ConnectionBorrowTimeout() *float64
	SetConnectionBorrowTimeout(val *float64)
	ConnectionBorrowTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InitQuery() *string
	SetInitQuery(val *string)
	InitQueryInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConnectionsPercent() *float64
	SetMaxConnectionsPercent(val *float64)
	MaxConnectionsPercentInput() *float64
	MaxIdleConnectionsPercent() *float64
	SetMaxIdleConnectionsPercent(val *float64)
	MaxIdleConnectionsPercentInput() *float64
	SessionPinningFilters() *[]*string
	SetSessionPinningFilters(val *[]*string)
	SessionPinningFiltersInput() *[]*string
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
	ResetConnectionBorrowTimeout()
	ResetInitQuery()
	ResetMaxConnectionsPercent()
	ResetMaxIdleConnectionsPercent()
	ResetSessionPinningFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference
type jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ConnectionBorrowTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ConnectionBorrowTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) InitQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) InitQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) MaxConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) MaxConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) MaxIdleConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) MaxIdleConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) SessionPinningFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) SessionPinningFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference {
	_init_.Initialize()

	if err := validateNewRdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference_Override(r RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbProxyTargetGroup.RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetConnectionBorrowTimeout(val *float64) {
	if err := j.validateSetConnectionBorrowTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionBorrowTimeout",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetInitQuery(val *string) {
	if err := j.validateSetInitQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initQuery",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetMaxConnectionsPercent(val *float64) {
	if err := j.validateSetMaxConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetMaxIdleConnectionsPercent(val *float64) {
	if err := j.validateSetMaxIdleConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetSessionPinningFilters(val *[]*string) {
	if err := j.validateSetSessionPinningFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPinningFilters",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ResetConnectionBorrowTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetConnectionBorrowTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ResetInitQuery() {
	_jsii_.InvokeVoid(
		r,
		"resetInitQuery",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ResetMaxConnectionsPercent() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxConnectionsPercent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ResetMaxIdleConnectionsPercent() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxIdleConnectionsPercent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ResetSessionPinningFilters() {
	_jsii_.InvokeVoid(
		r,
		"resetSessionPinningFilters",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbProxyTargetGroupConnectionPoolConfigurationInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

