// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/app/dao/internal"
)

// internalPromptsDao is internal type for wrapping internal DAO implements.
type internalPromptsDao = *internal.PromptsDao

// PromptsDao is the data access object for table forum_prompts.
// You can define custom methods on it to extend its functionality as you wish.
type PromptsDao struct {
	internalPromptsDao
}

var (
	// Prompts is globally public accessible object for table forum_prompts operations.
	Prompts = PromptsDao{
		internal.NewPromptsDao(),
	}
)

// Fill with you ideas below.