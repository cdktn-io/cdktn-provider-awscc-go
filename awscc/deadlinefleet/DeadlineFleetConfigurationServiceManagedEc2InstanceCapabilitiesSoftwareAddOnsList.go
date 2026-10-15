// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList interface {
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
	Get(index *float64) DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList
type jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList {
	_init_.Initialize()

	if err := validateNewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList_Override(d DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) Get(index *float64) DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesSoftwareAddOnsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

