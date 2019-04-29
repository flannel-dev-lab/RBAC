package RBAC

/*import (
	"github.com/flannel-dev-lab/RBAC/vars"
	"testing"
)


func TestGrantPermission(t *testing.T) {
    // Create an object
    object, err := CreateObject("test-object", "test-object-description")

    _, err = GrantPermission(object, TestOperation, 1)

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = RemoveObject(object)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestRevokePermission(t *testing.T) {
    // Create an object
    object, err := CreateObject("test-object", "test-object-description")
    if err != nil {
        t.Errorf("%v", err)
    }
    
    // Grant a permission
    _, err = GrantPermission(object, TestOperation, 1)
    if err != nil {
        t.Errorf("%v", err)
    }

    // This is what we are actually testing for
    _, err = RevokePermission(object, TestOperation, 1)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = RemoveObject(object)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestRolePermissions(t *testing.T) {
    role := vars.Role{Id: 1, Name: "test-role", Description: "Reserved role for testing"}
    _, err := RolePermissions(role)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserPermissions(t *testing.T) {
    user := vars.User{Id: 1}

    _, err := UserPermissions(user)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestSessionPermissions(t *testing.T) {
    user := vars.User{Id: 1}
    session, sessionErr := CreateSession(user.Id, "123-123-123")
    if sessionErr != nil {
        t.Errorf("%v", sessionErr)
    }

    // Permissions - what we are actually testing
    _, err := SessionPermissions(session)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = DeleteSession(user, session.Name)
    if err != nil {
        t.Errorf("%v", err)
    }
}
*/