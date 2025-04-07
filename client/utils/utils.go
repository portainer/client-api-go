package utils

import (
	"strconv"

	"github.com/portainer/client-api-go/v2/pkg/models"
)

// BuildAccessPolicies builds a map of access policies from a map of role names.
// The map key is the ID of the user or team, and the value is the role name.
// Invalid role names are ignored.
func BuildAccessPolicies[T models.PortainerTeamAccessPolicies | models.PortainerUserAccessPolicies](accesses map[int64]string) T {
	policies := make(T)
	for id, role := range accesses {
		roleID := getRoleFromName(role)
		if roleID != 0 {
			policies[strconv.Itoa(int(id))] = models.PortainerAccessPolicy{
				RoleID: roleID,
			}
		}
	}
	return policies
}

func getRoleFromName(roleName string) int64 {
	switch roleName {
	case "environment_administrator":
		return 1
	case "helpdesk_user":
		return 2
	case "standard_user":
		return 3
	case "readonly_user":
		return 4
	case "operator_user":
		return 5
	default:
		return 0
	}
}
