// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datasynctask

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncTaskExcludesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskExcludesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncTaskExcludesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskExcludesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskExcludesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskExcludesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncTaskExcludesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncTaskExcludesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

