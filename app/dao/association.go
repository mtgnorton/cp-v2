// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/app/dao/internal"
)

// internalAssociationDao is internal type for wrapping internal DAO implements.
type internalAssociationDao = *internal.AssociationDao

// AssociationDao is the data access object for table forum_association.
// You can define custom methods on it to extend its functionality as you wish.
type AssociationDao struct {
	internalAssociationDao
}

var (
	// Association is globally public accessible object for table forum_association operations.
	Association = AssociationDao{
		internal.NewAssociationDao(),
	}
)

// Fill with you ideas below.