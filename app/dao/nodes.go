// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/app/dao/internal"
)

// internalNodesDao is internal type for wrapping internal DAO implements.
type internalNodesDao = *internal.NodesDao

// NodesDao is the data access object for table forum_nodes.
// You can define custom methods on it to extend its functionality as you wish.
type NodesDao struct {
	internalNodesDao
}

var (
	// Nodes is globally public accessible object for table forum_nodes operations.
	Nodes = NodesDao{
		internal.NewNodesDao(),
	}
)

// Fill with you ideas below.
