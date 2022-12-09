package service

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

var (
	adminId = uint(1)
)

func TestAdministratorService_GetAdministratorSummary(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		admin, err := Administrator.GetAdministratorSummary(ctx, adminId)
		t.AssertNil(err)
		fmt.Printf("TestAdministratorService_GetAdministratorSummary rs is %#v", admin)
	})
}
