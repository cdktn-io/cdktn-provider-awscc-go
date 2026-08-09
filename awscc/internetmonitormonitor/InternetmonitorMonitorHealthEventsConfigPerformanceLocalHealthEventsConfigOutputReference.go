// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/internetmonitormonitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HealthScoreThreshold() *float64
	SetHealthScoreThreshold(val *float64)
	HealthScoreThresholdInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinTrafficImpact() *float64
	SetMinTrafficImpact(val *float64)
	MinTrafficImpactInput() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	ResetHealthScoreThreshold()
	ResetMinTrafficImpact()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference
type jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) HealthScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthScoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) HealthScoreThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthScoreThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) MinTrafficImpact() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTrafficImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) MinTrafficImpactInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTrafficImpactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewInternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.internetmonitorMonitor.InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference_Override(i InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.internetmonitorMonitor.InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetHealthScoreThreshold(val *float64) {
	if err := j.validateSetHealthScoreThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthScoreThreshold",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetMinTrafficImpact(val *float64) {
	if err := j.validateSetMinTrafficImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTrafficImpact",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ResetHealthScoreThreshold() {
	_jsii_.InvokeVoid(
		i,
		"resetHealthScoreThreshold",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ResetMinTrafficImpact() {
	_jsii_.InvokeVoid(
		i,
		"resetMinTrafficImpact",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InternetmonitorMonitorHealthEventsConfigPerformanceLocalHealthEventsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

