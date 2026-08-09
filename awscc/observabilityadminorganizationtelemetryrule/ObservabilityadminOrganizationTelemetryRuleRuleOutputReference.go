// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationtelemetryrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/observabilityadminorganizationtelemetryrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminOrganizationTelemetryRuleRuleOutputReference interface {
	cdktn.ComplexObject
	AllowFieldUpdates() interface{}
	SetAllowFieldUpdates(val interface{})
	AllowFieldUpdatesInput() interface{}
	AllRegions() interface{}
	SetAllRegions(val interface{})
	AllRegionsInput() interface{}
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
	DestinationConfiguration() ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationOutputReference
	DestinationConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SelectionCriteria() *string
	SetSelectionCriteria(val *string)
	SelectionCriteriaInput() *string
	TelemetrySourceTypes() *[]*string
	SetTelemetrySourceTypes(val *[]*string)
	TelemetrySourceTypesInput() *[]*string
	TelemetryType() *string
	SetTelemetryType(val *string)
	TelemetryTypeInput() *string
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
	PutDestinationConfiguration(value *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfiguration)
	ResetAllowFieldUpdates()
	ResetAllRegions()
	ResetDestinationConfiguration()
	ResetRegions()
	ResetScope()
	ResetSelectionCriteria()
	ResetTelemetrySourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminOrganizationTelemetryRuleRuleOutputReference
type jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) AllowFieldUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFieldUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) AllowFieldUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFieldUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) AllRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) AllRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) DestinationConfiguration() ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationOutputReference {
	var returns ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) DestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) SelectionCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) SelectionCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TelemetrySourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetrySourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TelemetrySourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetrySourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TelemetryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telemetryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TelemetryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telemetryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityadminOrganizationTelemetryRuleRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ObservabilityadminOrganizationTelemetryRuleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminOrganizationTelemetryRuleRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationTelemetryRule.ObservabilityadminOrganizationTelemetryRuleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityadminOrganizationTelemetryRuleRuleOutputReference_Override(o ObservabilityadminOrganizationTelemetryRuleRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminOrganizationTelemetryRule.ObservabilityadminOrganizationTelemetryRuleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetAllowFieldUpdates(val interface{}) {
	if err := j.validateSetAllowFieldUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFieldUpdates",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetAllRegions(val interface{}) {
	if err := j.validateSetAllRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allRegions",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetSelectionCriteria(val *string) {
	if err := j.validateSetSelectionCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionCriteria",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetTelemetrySourceTypes(val *[]*string) {
	if err := j.validateSetTelemetrySourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetrySourceTypes",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetTelemetryType(val *string) {
	if err := j.validateSetTelemetryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetryType",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) PutDestinationConfiguration(value *ObservabilityadminOrganizationTelemetryRuleRuleDestinationConfiguration) {
	if err := o.validatePutDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDestinationConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetAllowFieldUpdates() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowFieldUpdates",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetAllRegions() {
	_jsii_.InvokeVoid(
		o,
		"resetAllRegions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetDestinationConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		o,
		"resetRegions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		o,
		"resetScope",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetSelectionCriteria() {
	_jsii_.InvokeVoid(
		o,
		"resetSelectionCriteria",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ResetTelemetrySourceTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetTelemetrySourceTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminOrganizationTelemetryRuleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

