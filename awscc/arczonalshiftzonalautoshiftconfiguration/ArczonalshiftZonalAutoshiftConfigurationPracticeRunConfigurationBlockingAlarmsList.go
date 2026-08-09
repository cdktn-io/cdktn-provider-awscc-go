// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/arczonalshiftzonalautoshiftconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList
type jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList {
	_init_.Initialize()

	if err := validateNewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList_Override(a ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.arczonalshiftZonalAutoshiftConfiguration.ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) Get(index *float64) ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfigurationBlockingAlarmsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

