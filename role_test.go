package mongodb

import "testing"

// TestRole_Create ...
func TestRole_Create(t *testing.T) {
	NewGenesis().Create()
	NewAdmin().Create()
	NewMonitor().Create()
	NewGod().Create()
	NewOrg().Create()
}

// TestRoleUser_Create ...
func TestRoleUser_Create(t *testing.T) {
	ru := NewRoleUser()
	ru.RoleID = ID("5c336fb32509a10d616b3b35")
	ru.UserID = ID("5c33711e06b5362b5f8dccbf")
	e := ru.CreateIfNotExist()
	t.Log(ru, e)
}
