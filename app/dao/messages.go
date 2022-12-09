// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/app/dao/internal"
)

// internalMessagesDao is internal type for wrapping internal DAO implements.
type internalMessagesDao = *internal.MessagesDao

// MessagesDao is the data access object for table forum_messages.
// You can define custom methods on it to extend its functionality as you wish.
type MessagesDao struct {
	internalMessagesDao
}

var (
	// Messages is globally public accessible object for table forum_messages operations.
	Messages = MessagesDao{
		internal.NewMessagesDao(),
	}
)

// Fill with you ideas below.