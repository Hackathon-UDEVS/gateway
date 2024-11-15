package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

// Enforcer is the Casbin enforcer
var Enforcer *casbin.Enforcer

// InitCasbin initializes the Casbin enforcer with RBAC model and policy
func InitCasbin() error {
	// Define RBAC model
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`

	m, err := model.NewModelFromString(text)
	if err != nil {
		return err
	}

	// Create enforcer
	enforcer, err := casbin.NewEnforcer(m)
	if err != nil {
		return err
	}

	// Add policies for admin role
	enforcer.AddPolicy("admin", "/api/*", "*")
	
	// Add policies for user role
	enforcer.AddPolicy("user", "/api/v1/test", "GET")
	enforcer.AddPolicy("user", "/api/v1/function1", "GET")
	enforcer.AddPolicy("user", "/api/v1/function2", "GET")

	// Add role assignments
	enforcer.AddGroupingPolicy("admin", "user") // Admin inherits user permissions

	Enforcer = enforcer
	return nil
}