// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/customerprofilesintegration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesIntegrationFlowDefinitionTasksOutputReference interface {
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
	ConnectorOperator() CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperatorOutputReference
	ConnectorOperatorInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationField() *string
	SetDestinationField(val *string)
	DestinationFieldInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceFields() *[]*string
	SetSourceFields(val *[]*string)
	SourceFieldsInput() *[]*string
	TaskProperties() CustomerprofilesIntegrationFlowDefinitionTasksTaskPropertiesList
	TaskPropertiesInput() interface{}
	TaskType() *string
	SetTaskType(val *string)
	TaskTypeInput() *string
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
	PutConnectorOperator(value *CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperator)
	PutTaskProperties(value interface{})
	ResetConnectorOperator()
	ResetDestinationField()
	ResetSourceFields()
	ResetTaskProperties()
	ResetTaskType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesIntegrationFlowDefinitionTasksOutputReference
type jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ConnectorOperator() CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperatorOutputReference {
	var returns CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperatorOutputReference
	_jsii_.Get(
		j,
		"connectorOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ConnectorOperatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) DestinationField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) DestinationFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) SourceFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) SourceFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TaskProperties() CustomerprofilesIntegrationFlowDefinitionTasksTaskPropertiesList {
	var returns CustomerprofilesIntegrationFlowDefinitionTasksTaskPropertiesList
	_jsii_.Get(
		j,
		"taskProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TaskPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesIntegrationFlowDefinitionTasksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CustomerprofilesIntegrationFlowDefinitionTasksOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesIntegrationFlowDefinitionTasksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionTasksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCustomerprofilesIntegrationFlowDefinitionTasksOutputReference_Override(c CustomerprofilesIntegrationFlowDefinitionTasksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionTasksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetDestinationField(val *string) {
	if err := j.validateSetDestinationFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationField",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetSourceFields(val *[]*string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetTaskType(val *string) {
	if err := j.validateSetTaskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) PutConnectorOperator(value *CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperator) {
	if err := c.validatePutConnectorOperatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConnectorOperator",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) PutTaskProperties(value interface{}) {
	if err := c.validatePutTaskPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaskProperties",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ResetConnectorOperator() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectorOperator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ResetDestinationField() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ResetSourceFields() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ResetTaskProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ResetTaskType() {
	_jsii_.InvokeVoid(
		c,
		"resetTaskType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionTasksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

