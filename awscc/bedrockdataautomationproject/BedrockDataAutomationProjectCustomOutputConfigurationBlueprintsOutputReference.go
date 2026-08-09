// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference interface {
	cdktn.ComplexObject
	BlueprintArn() *string
	SetBlueprintArn(val *string)
	BlueprintArnInput() *string
	BlueprintStage() *string
	SetBlueprintStage(val *string)
	BlueprintStageInput() *string
	BlueprintVersion() *string
	SetBlueprintVersion(val *string)
	BlueprintVersionInput() *string
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
	ResetBlueprintArn()
	ResetBlueprintStage()
	ResetBlueprintVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference
type jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) BlueprintVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference_Override(b BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetBlueprintArn(val *string) {
	if err := j.validateSetBlueprintArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blueprintArn",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetBlueprintStage(val *string) {
	if err := j.validateSetBlueprintStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blueprintStage",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetBlueprintVersion(val *string) {
	if err := j.validateSetBlueprintVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blueprintVersion",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ResetBlueprintArn() {
	_jsii_.InvokeVoid(
		b,
		"resetBlueprintArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ResetBlueprintStage() {
	_jsii_.InvokeVoid(
		b,
		"resetBlueprintStage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ResetBlueprintVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetBlueprintVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectCustomOutputConfigurationBlueprintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

