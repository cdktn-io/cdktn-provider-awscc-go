// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetableoptimizer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluetableoptimizer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableOptimizerTableOptimizerConfigurationOutputReference interface {
	cdktn.ComplexObject
	CompactionConfiguration() GlueTableOptimizerTableOptimizerConfigurationCompactionConfigurationOutputReference
	CompactionConfigurationInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OrphanFileDeletionConfiguration() GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfigurationOutputReference
	OrphanFileDeletionConfigurationInput() interface{}
	RetentionConfiguration() GlueTableOptimizerTableOptimizerConfigurationRetentionConfigurationOutputReference
	RetentionConfigurationInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcConfiguration() GlueTableOptimizerTableOptimizerConfigurationVpcConfigurationOutputReference
	VpcConfigurationInput() interface{}
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
	PutCompactionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationCompactionConfiguration)
	PutOrphanFileDeletionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfiguration)
	PutRetentionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationRetentionConfiguration)
	PutVpcConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationVpcConfiguration)
	ResetCompactionConfiguration()
	ResetOrphanFileDeletionConfiguration()
	ResetRetentionConfiguration()
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableOptimizerTableOptimizerConfigurationOutputReference
type jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) CompactionConfiguration() GlueTableOptimizerTableOptimizerConfigurationCompactionConfigurationOutputReference {
	var returns GlueTableOptimizerTableOptimizerConfigurationCompactionConfigurationOutputReference
	_jsii_.Get(
		j,
		"compactionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) CompactionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compactionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) OrphanFileDeletionConfiguration() GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfigurationOutputReference {
	var returns GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfigurationOutputReference
	_jsii_.Get(
		j,
		"orphanFileDeletionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) OrphanFileDeletionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orphanFileDeletionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) RetentionConfiguration() GlueTableOptimizerTableOptimizerConfigurationRetentionConfigurationOutputReference {
	var returns GlueTableOptimizerTableOptimizerConfigurationRetentionConfigurationOutputReference
	_jsii_.Get(
		j,
		"retentionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) RetentionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) VpcConfiguration() GlueTableOptimizerTableOptimizerConfigurationVpcConfigurationOutputReference {
	var returns GlueTableOptimizerTableOptimizerConfigurationVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) VpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewGlueTableOptimizerTableOptimizerConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableOptimizerTableOptimizerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableOptimizerTableOptimizerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTableOptimizer.GlueTableOptimizerTableOptimizerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableOptimizerTableOptimizerConfigurationOutputReference_Override(g GlueTableOptimizerTableOptimizerConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTableOptimizer.GlueTableOptimizerTableOptimizerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) PutCompactionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationCompactionConfiguration) {
	if err := g.validatePutCompactionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCompactionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) PutOrphanFileDeletionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationOrphanFileDeletionConfiguration) {
	if err := g.validatePutOrphanFileDeletionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOrphanFileDeletionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) PutRetentionConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationRetentionConfiguration) {
	if err := g.validatePutRetentionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetentionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) PutVpcConfiguration(value *GlueTableOptimizerTableOptimizerConfigurationVpcConfiguration) {
	if err := g.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ResetCompactionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetCompactionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ResetOrphanFileDeletionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetOrphanFileDeletionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ResetRetentionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOptimizerTableOptimizerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

