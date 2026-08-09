// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionidnamespace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/entityresolutionidnamespace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference interface {
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
	ProviderConfiguration() *map[string]*string
	SetProviderConfiguration(val *map[string]*string)
	ProviderConfigurationInput() *map[string]*string
	ProviderServiceArn() *string
	SetProviderServiceArn(val *string)
	ProviderServiceArnInput() *string
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
	ResetProviderConfiguration()
	ResetProviderServiceArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference
type jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ProviderConfiguration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ProviderConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ProviderServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ProviderServiceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerServiceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewEntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionIdNamespace.EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference_Override(e EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionIdNamespace.EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetProviderConfiguration(val *map[string]*string) {
	if err := j.validateSetProviderConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerConfiguration",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetProviderServiceArn(val *string) {
	if err := j.validateSetProviderServiceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerServiceArn",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ResetProviderConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetProviderConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ResetProviderServiceArn() {
	_jsii_.InvokeVoid(
		e,
		"resetProviderServiceArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

