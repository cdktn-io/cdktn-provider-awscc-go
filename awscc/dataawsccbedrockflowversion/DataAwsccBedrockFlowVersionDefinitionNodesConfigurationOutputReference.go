// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbedrockflowversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbedrockflowversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference interface {
	cdktn.ComplexObject
	Agent() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationAgentOutputReference
	Collector() *string
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
	Condition() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationConditionOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InlineCode() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationInlineCodeOutputReference
	Input() *string
	InternalValue() *DataAwsccBedrockFlowVersionDefinitionNodesConfiguration
	SetInternalValue(val *DataAwsccBedrockFlowVersionDefinitionNodesConfiguration)
	Iterator() *string
	KnowledgeBase() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationKnowledgeBaseOutputReference
	LambdaFunction() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLambdaFunctionOutputReference
	Lex() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLexOutputReference
	Output() *string
	Prompt() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationPromptOutputReference
	Retrieval() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationRetrievalOutputReference
	Storage() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationStorageOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference
type jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Agent() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationAgentOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationAgentOutputReference
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Collector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Condition() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationConditionOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) InlineCode() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationInlineCodeOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationInlineCodeOutputReference
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) InternalValue() *DataAwsccBedrockFlowVersionDefinitionNodesConfiguration {
	var returns *DataAwsccBedrockFlowVersionDefinitionNodesConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Iterator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iterator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) KnowledgeBase() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationKnowledgeBaseOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationKnowledgeBaseOutputReference
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) LambdaFunction() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLambdaFunctionOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLambdaFunctionOutputReference
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Lex() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLexOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationLexOutputReference
	_jsii_.Get(
		j,
		"lex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Output() *string {
	var returns *string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Prompt() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationPromptOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationPromptOutputReference
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Retrieval() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationRetrievalOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationRetrievalOutputReference
	_jsii_.Get(
		j,
		"retrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Storage() DataAwsccBedrockFlowVersionDefinitionNodesConfigurationStorageOutputReference {
	var returns DataAwsccBedrockFlowVersionDefinitionNodesConfigurationStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockFlowVersion.DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference_Override(d DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockFlowVersion.DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference)SetInternalValue(val *DataAwsccBedrockFlowVersionDefinitionNodesConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockFlowVersionDefinitionNodesConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

