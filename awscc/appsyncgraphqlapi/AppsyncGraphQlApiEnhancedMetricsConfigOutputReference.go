// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncgraphqlapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appsyncgraphqlapi/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncGraphQlApiEnhancedMetricsConfigOutputReference interface {
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
	DataSourceLevelMetricsBehavior() *string
	SetDataSourceLevelMetricsBehavior(val *string)
	DataSourceLevelMetricsBehaviorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperationLevelMetricsConfig() *string
	SetOperationLevelMetricsConfig(val *string)
	OperationLevelMetricsConfigInput() *string
	ResolverLevelMetricsBehavior() *string
	SetResolverLevelMetricsBehavior(val *string)
	ResolverLevelMetricsBehaviorInput() *string
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
	ResetDataSourceLevelMetricsBehavior()
	ResetOperationLevelMetricsConfig()
	ResetResolverLevelMetricsBehavior()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncGraphQlApiEnhancedMetricsConfigOutputReference
type jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) DataSourceLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) DataSourceLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) OperationLevelMetricsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) OperationLevelMetricsConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ResolverLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ResolverLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncGraphQlApiEnhancedMetricsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppsyncGraphQlApiEnhancedMetricsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncGraphQlApiEnhancedMetricsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncGraphQlApi.AppsyncGraphQlApiEnhancedMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsyncGraphQlApiEnhancedMetricsConfigOutputReference_Override(a AppsyncGraphQlApiEnhancedMetricsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appsyncGraphQlApi.AppsyncGraphQlApiEnhancedMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetDataSourceLevelMetricsBehavior(val *string) {
	if err := j.validateSetDataSourceLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetOperationLevelMetricsConfig(val *string) {
	if err := j.validateSetOperationLevelMetricsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationLevelMetricsConfig",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetResolverLevelMetricsBehavior(val *string) {
	if err := j.validateSetResolverLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ResetDataSourceLevelMetricsBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSourceLevelMetricsBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ResetOperationLevelMetricsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationLevelMetricsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ResetResolverLevelMetricsBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetResolverLevelMetricsBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphQlApiEnhancedMetricsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

