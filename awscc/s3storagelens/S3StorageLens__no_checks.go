// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3storagelens

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3StorageLens) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validatePutStorageLensConfigurationParameters(value *S3StorageLensStorageLensConfiguration) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_S3StorageLens) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateS3StorageLens_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateS3StorageLens_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3StorageLens_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateS3StorageLens_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_S3StorageLens) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3StorageLens) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3StorageLens) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_S3StorageLens) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewS3StorageLensParameters(scope constructs.Construct, id *string, config *S3StorageLensConfig) error {
	return nil
}

