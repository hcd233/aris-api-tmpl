package aggregate

import (
	"testing"
	"time"

	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/common/enum"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity/vo"
)

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 5, 6, 15, 0, 0, 0, time.UTC)
	user, err := RegisterUser(vo.UserName("alice"), vo.Email("alice@example.com"), vo.Avatar("avatar"), constant.OAuthProviderGithub, "github-id", now)
	if err != nil {
		t.Fatalf("RegisterUser() error = %v", err)
	}
	if user.Name().String() != "alice" {
		t.Fatalf("RegisterUser().Name() = %q, want %q", user.Name().String(), "alice")
	}
	if user.Permission() != enum.PermissionPending {
		t.Fatalf("RegisterUser().Permission() = %q, want %q", user.Permission(), enum.PermissionPending)
	}
	if user.GithubBindID() != "github-id" {
		t.Fatalf("RegisterUser().GithubBindID() = %q, want %q", user.GithubBindID(), "github-id")
	}
	if !user.LastLogin().Equal(now) {
		t.Fatalf("RegisterUser().LastLogin() = %v, want %v", user.LastLogin(), now)
	}
}

func TestUser_UpdateProfile_RejectsEmptyName(t *testing.T) {
	t.Parallel()

	user := RestoreUser(1, vo.UserName("alice"), vo.Email("alice@example.com"), vo.Avatar("avatar"), enum.PermissionUser, time.Time{}, time.Time{}, "", "")
	if err := user.UpdateProfile(vo.UserName(""), vo.Email("new@example.com"), vo.Avatar("new")); err == nil {
		t.Fatal("UpdateProfile() error = nil, want validation error")
	}
}
