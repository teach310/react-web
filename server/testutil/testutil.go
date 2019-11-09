package testutil

import (
	"todo/infra/mysql"
)

// TestInit
func TestInit() {
	mysql.Connect()
	mysql.Migrate()
}
