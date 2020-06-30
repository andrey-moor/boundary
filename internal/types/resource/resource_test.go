package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Resource(t *testing.T) {

	tests := []struct {
		typeString string
		want       Type
	}{
		{
			typeString: "unknown",
			want:       Unknown,
		},
		{
			typeString: "scope",
			want:       Scope,
		},
		{
			typeString: "user",
			want:       User,
		},
		{
			typeString: "static-group",
			want:       StaticGroup,
		},
		{
			typeString: "role",
			want:       Role,
		},
		{
			typeString: "organization",
			want:       Organization,
		},
		{
			typeString: "static-group-member",
			want:       StaticGroupMember,
		},
		{
			typeString: "static-group-user-member",
			want:       StaticGroupUserMember,
		},
		{
			typeString: "assigned-role",
			want:       AssignedRole,
		},
		{
			typeString: "assigned-user-role",
			want:       AssignedUserRole,
		},
		{
			typeString: "assigned-static-group-role",
			want:       AssignedStaticGroupRole,
		},
		{
			typeString: "role-grant",
			want:       RoleGrant,
		},
		{
			typeString: "auth-method",
			want:       AuthMethod,
		},
		{
			typeString: "project",
			want:       Project,
		},
		{
			typeString: "*",
			want:       All,
		},
		{
			typeString: "host-catalog",
			want:       HostCatalog,
		},
		{
			typeString: "host-set",
			want:       HostSet,
		},
		{
			typeString: "host",
			want:       Host,
		},
		{
			typeString: "target",
			want:       Target,
		},
	}
	for _, tt := range tests {
		t.Run(tt.typeString, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringToResourceType(tt.typeString), "unexpected type for %s", tt.typeString)
			assert.Equalf(t, tt.typeString, tt.want.String(), "unexpected string for %s", tt.typeString)
		})
	}
}