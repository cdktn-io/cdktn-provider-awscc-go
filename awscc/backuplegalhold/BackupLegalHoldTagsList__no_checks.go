// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backuplegalhold

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupLegalHoldTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupLegalHoldTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupLegalHoldTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupLegalHoldTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupLegalHoldTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupLegalHoldTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupLegalHoldTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupLegalHoldTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

