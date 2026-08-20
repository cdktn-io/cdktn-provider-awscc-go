// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/inspectorv2connector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference interface {
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

// The jsii proxy struct for Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference
type jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ScopeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ScopeValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ScopeValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) StateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference_Override(i Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetScopeType(val *string) {
	if err := j.validateSetScopeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeType",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetScopeValues(val *[]*string) {
	if err := j.validateSetScopeValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeValues",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ResetScopeType() {
	_jsii_.InvokeVoid(
		i,
		"resetScopeType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ResetScopeValues() {
	_jsii_.InvokeVoid(
		i,
		"resetScopeValues",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

