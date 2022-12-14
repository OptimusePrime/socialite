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
		UsersTable,
	}
)

func init() {
	FollowsTable.ForeignKeys[0].RefTable = UsersTable
	FollowsTable.ForeignKeys[1].RefTable = UsersTable
}
