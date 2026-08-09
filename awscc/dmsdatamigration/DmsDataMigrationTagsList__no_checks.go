// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dmsdatamigration

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DmsDataMigrationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DmsDataMigrationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DmsDataMigrationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DmsDataMigrationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDmsDataMigrationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

