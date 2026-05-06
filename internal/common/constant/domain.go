package constant

const (
	// AggregateTypeUser 用户聚合类型。
	AggregateTypeUser = "user"

	// FieldName 数据库 name 字段。
	FieldName = "name"

	// FieldEmail 数据库 email 字段。
	FieldEmail = "email"

	// FieldAvatar 数据库 avatar 字段。
	FieldAvatar = "avatar"

	// FieldPermission 数据库 permission 字段。
	FieldPermission = "permission"

	// FieldLastLogin 数据库 last_login 字段。
	FieldLastLogin = "last_login"

	// OAuthProviderGithub Github OAuth2 平台名。
	OAuthProviderGithub = "github"

	// OAuthProviderGoogle Google OAuth2 平台名。
	OAuthProviderGoogle = "google"

	// DefaultUserNamePrefix OAuth2 用户名非法时的默认前缀。
	DefaultUserNamePrefix = "ArisUser"
)

// UserRepoFieldsFull 返回用户仓储重建聚合所需字段。
func UserRepoFieldsFull() []string {
	return []string{
		"id",
		"name",
		"email",
		"avatar",
		"permission",
		"last_login",
		"created_at",
		"github_bind_id",
		"google_bind_id",
	}
}
