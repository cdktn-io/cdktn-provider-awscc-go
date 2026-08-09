// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/applicationsignalsservicelevelobjective/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference interface {
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
	CompositeSliConfig() ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfigOutputReference
	CompositeSliConfigInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependencyConfig() ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfigOutputReference
	DependencyConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyAttributes() *map[string]*string
	SetKeyAttributes(val *map[string]*string)
	KeyAttributesInput() *map[string]*string
	MetricDataQueries() ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricDataQueriesList
	MetricDataQueriesInput() interface{}
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricSource() ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSourceOutputReference
	MetricSourceInput() interface{}
	MetricType() *string
	SetMetricType(val *string)
	MetricTypeInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
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
	PutCompositeSliConfig(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfig)
	PutDependencyConfig(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfig)
	PutMetricDataQueries(value interface{})
	PutMetricSource(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSource)
	ResetCompositeSliConfig()
	ResetDependencyConfig()
	ResetKeyAttributes()
	ResetMetricDataQueries()
	ResetMetricName()
	ResetMetricSource()
	ResetMetricType()
	ResetOperationName()
	ResetPeriodSeconds()
	ResetStatistic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference
type jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) CompositeSliConfig() ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfigOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfigOutputReference
	_jsii_.Get(
		j,
		"compositeSliConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) CompositeSliConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeSliConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) DependencyConfig() ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfigOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfigOutputReference
	_jsii_.Get(
		j,
		"dependencyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) DependencyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) KeyAttributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) KeyAttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricDataQueries() ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricDataQueriesList {
	var returns ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricDataQueriesList
	_jsii_.Get(
		j,
		"metricDataQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricDataQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDataQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricSource() ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSourceOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSourceOutputReference
	_jsii_.Get(
		j,
		"metricSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) MetricTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference_Override(a ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetKeyAttributes(val *map[string]*string) {
	if err := j.validateSetKeyAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAttributes",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetMetricType(val *string) {
	if err := j.validateSetMetricTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricType",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PutCompositeSliConfig(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricCompositeSliConfig) {
	if err := a.validatePutCompositeSliConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCompositeSliConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PutDependencyConfig(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfig) {
	if err := a.validatePutDependencyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDependencyConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PutMetricDataQueries(value interface{}) {
	if err := a.validatePutMetricDataQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricDataQueries",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) PutMetricSource(value *ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricSource) {
	if err := a.validatePutMetricSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetCompositeSliConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCompositeSliConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetDependencyConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDependencyConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetKeyAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetMetricDataQueries() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricDataQueries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetMetricName() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetMetricSource() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetMetricType() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ResetStatistic() {
	_jsii_.InvokeVoid(
		a,
		"resetStatistic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveSliSliMetricOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

