// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueregistry

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueRegistry) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueRegistry) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateGlueRegistry_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGlueRegistry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlueRegistry_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGlueRegistry_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueRegistry) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewGlueRegistryParameters(scope constructs.Construct, id *string, config *GlueRegistryConfig) error {
	return nil
}

