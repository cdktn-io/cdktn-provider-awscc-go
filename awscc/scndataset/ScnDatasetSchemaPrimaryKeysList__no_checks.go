// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scndataset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScnDatasetSchemaPrimaryKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScnDatasetSchemaPrimaryKeysListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

