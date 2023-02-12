// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FollowsColumns holds the columns for the "follows" table.
	FollowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "follow_follower", Type: field.TypeUUID},
		{Name: "follow_followee", Type: field.TypeUUID},
	}
	// FollowsTable holds the schema information for the "follows" table.
	FollowsTable = &schema.Table{
		Name:       "follows",
		Columns:    FollowsColumns,
		PrimaryKey: []*schema.Column{FollowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "follows_users_follower",
				Columns:    []*schema.Column{FollowsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "follows_users_followee",
				Columns:    []*schema.Column{FollowsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LikesColumns holds the columns for the "likes" table.
	LikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "post_likes", Type: field.TypeUUID, Nullable: true},
		{Name: "user_likes", Type: field.TypeUUID, Nullable: true},
	}
	// LikesTable holds the schema information for the "likes" table.
	LikesTable = &schema.Table{
		Name:       "likes",
		Columns:    LikesColumns,
		PrimaryKey: []*schema.Column{LikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "likes_posts_likes",
				Columns:    []*schema.Column{LikesColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "likes_users_likes",
				Columns:    []*schema.Column{LikesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "caption", Type: field.TypeString},
		{Name: "images", Type: field.TypeJSON},
		{Name: "user_posts", Type: field.TypeUUID, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "old", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "birth_date", Type: field.TypeTime, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "biography", Type: field.TypeString, Nullable: true},
		{Name: "gender", Type: field.TypeString, Nullable: true},
		{Name: "pronouns", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FollowsTable,
		LikesTable,
		PostsTable,
		UsersTable,
	}
)

func init() {
	FollowsTable.ForeignKeys[0].RefTable = UsersTable
	FollowsTable.ForeignKeys[1].RefTable = UsersTable
	LikesTable.ForeignKeys[0].RefTable = PostsTable
	LikesTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
}
