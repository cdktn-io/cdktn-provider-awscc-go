// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsdbsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbSnapshotTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDbSnapshotTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbSnapshotTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbSnapshotTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDbSnapshotTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbSnapshotTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbSnapshotTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbSnapshotTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

