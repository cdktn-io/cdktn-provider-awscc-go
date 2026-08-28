// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gluesession

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueSession) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validatePutCommandParameters(value *GlueSessionCommand) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validatePutConnectionsParameters(value *GlueSessionConnections) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueSession) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateGlueSession_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGlueSession_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlueSession_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGlueSession_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetDefaultArgumentsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetGlueVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetIdleTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetMaxCapacityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetNumberOfWorkersParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetRequestOriginParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetRoleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetSecurityConfigurationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetSessionIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_GlueSession) validateSetWorkerTypeParameters(val *string) error {
	return nil
}

func validateNewGlueSessionParameters(scope constructs.Construct, id *string, config *GlueSessionConfig) error {
	return nil
}

