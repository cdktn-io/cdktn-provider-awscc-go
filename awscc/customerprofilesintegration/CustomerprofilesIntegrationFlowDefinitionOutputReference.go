// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/customerprofilesintegration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesIntegrationFlowDefinitionOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FlowName() *string
	SetFlowName(val *string)
	FlowNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsArn() *string
	SetKmsArn(val *string)
	KmsArnInput() *string
	SourceFlowConfig() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference
	SourceFlowConfigInput() interface{}
	Tasks() CustomerprofilesIntegrationFlowDefinitionTasksList
	TasksInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggerConfig() CustomerprofilesIntegrationFlowDefinitionTriggerConfigOutputReference
	TriggerConfigInput() interface{}
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
	PutSourceFlowConfig(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig)
	PutTasks(value interface{})
	PutTriggerConfig(value *CustomerprofilesIntegrationFlowDefinitionTriggerConfig)
	ResetDescription()
	ResetFlowName()
	ResetKmsArn()
	ResetSourceFlowConfig()
	ResetTasks()
	ResetTriggerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesIntegrationFlowDefinitionOutputReference
type jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) FlowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) FlowNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) KmsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) SourceFlowConfig() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference {
	var returns CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference
	_jsii_.Get(
		j,
		"sourceFlowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) SourceFlowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFlowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) Tasks() CustomerprofilesIntegrationFlowDefinitionTasksList {
	var returns CustomerprofilesIntegrationFlowDefinitionTasksList
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) TasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) TriggerConfig() CustomerprofilesIntegrationFlowDefinitionTriggerConfigOutputReference {
	var returns CustomerprofilesIntegrationFlowDefinitionTriggerConfigOutputReference
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) TriggerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfigInput",
		&returns,
	)
	return returns
}


func NewCustomerprofilesIntegrationFlowDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CustomerprofilesIntegrationFlowDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesIntegrationFlowDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesIntegrationFlowDefinitionOutputReference_Override(c CustomerprofilesIntegrationFlowDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetFlowName(val *string) {
	if err := j.validateSetFlowNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetKmsArn(val *string) {
	if err := j.validateSetKmsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) PutSourceFlowConfig(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig) {
	if err := c.validatePutSourceFlowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceFlowConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) PutTasks(value interface{}) {
	if err := c.validatePutTasksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTasks",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) PutTriggerConfig(value *CustomerprofilesIntegrationFlowDefinitionTriggerConfig) {
	if err := c.validatePutTriggerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTriggerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetFlowName() {
	_jsii_.InvokeVoid(
		c,
		"resetFlowName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetKmsArn() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetSourceFlowConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceFlowConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetTasks() {
	_jsii_.InvokeVoid(
		c,
		"resetTasks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ResetTriggerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

