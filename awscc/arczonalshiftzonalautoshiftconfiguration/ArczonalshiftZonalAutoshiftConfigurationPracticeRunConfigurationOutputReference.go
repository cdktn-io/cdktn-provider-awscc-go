// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/arczonalshiftzonalautoshiftconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference interface {
	cdktn.ComplexObject
	BlockedDates() *[]*string
	SetBlockedDates(val *[]*string)
	BlockedDatesInput() *[]*string
	BlockedWindows() *[]*string
	SetBlockedWindows(val *[]*string)
	BlockedWindowsInput() *[]*string
	BlockingAlarms() ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList
	BlockingAlarmsInput() interface{}
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
	OutcomeAlarms() ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutcomeAlarmsList
	OutcomeAlarmsInput() interface{}
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
	PutBlockingAlarms(value interface{})
	PutOutcomeAlarms(value interface{})
	ResetBlockedDates()
	ResetBlockedWindows()
	ResetBlockingAlarms()
	ResetOutcomeAlarms()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference
type jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockedDates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockedDatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedDatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockedWindows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockedWindowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockingAlarms() ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList {
	var returns ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList
	_jsii_.Get(
		j,
		"blockingAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) BlockingAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockingAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) OutcomeAlarms() ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutcomeAlarmsList {
	var returns ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutcomeAlarmsList
	_jsii_.Get(
		j,
		"outcomeAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) OutcomeAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outcomeAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference_Override(a ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetBlockedDates(val *[]*string) {
	if err := j.validateSetBlockedDatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedDates",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetBlockedWindows(val *[]*string) {
	if err := j.validateSetBlockedWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedWindows",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) PutBlockingAlarms(value interface{}) {
	if err := a.validatePutBlockingAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockingAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) PutOutcomeAlarms(value interface{}) {
	if err := a.validatePutOutcomeAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutcomeAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ResetBlockedDates() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedDates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ResetBlockedWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockedWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ResetBlockingAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockingAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ResetOutcomeAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetOutcomeAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

