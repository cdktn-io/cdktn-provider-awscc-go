// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/inspectorv2connector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScopeType() *string
	SetScopeType(val *string)
	ScopeTypeInput() *string
	ScopeValues() *[]*string
	SetScopeValues(val *[]*string)
	ScopeValuesInput() *[]*string
	State() *string
	StateReason() *string
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
	ResetScopeType()
	ResetScopeValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference
type jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ScopeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ScopeValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ScopeValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) StateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference_Override(i Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetScopeType(val *string) {
	if err := j.validateSetScopeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeType",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetScopeValues(val *[]*string) {
	if err := j.validateSetScopeValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeValues",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ResetScopeType() {
	_jsii_.InvokeVoid(
		i,
		"resetScopeType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ResetScopeValues() {
	_jsii_.InvokeVoid(
		i,
		"resetScopeValues",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

