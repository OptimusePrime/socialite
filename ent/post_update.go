// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"socialite/ent/favourite"
	"socialite/ent/like"
	"socialite/ent/post"
	"socialite/ent/predicate"
	"socialite/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCaption sets the "caption" field.
func (pu *PostUpdate) SetCaption(s string) *PostUpdate {
	pu.mutation.SetCaption(s)
	return pu
}

// SetImages sets the "images" field.
func (pu *PostUpdate) SetImages(s []string) *PostUpdate {
	pu.mutation.SetImages(s)
	return pu
}

// AppendImages appends s to the "images" field.
func (pu *PostUpdate) AppendImages(s []string) *PostUpdate {
	pu.mutation.AppendImages(s)
	return pu
}

// SetLocation sets the "location" field.
func (pu *PostUpdate) SetLocation(s string) *PostUpdate {
	pu.mutation.SetLocation(s)
	return pu
}

// SetPosterID sets the "poster" edge to the User entity by ID.
func (pu *PostUpdate) SetPosterID(id uuid.UUID) *PostUpdate {
	pu.mutation.SetPosterID(id)
	return pu
}

// SetNillablePosterID sets the "poster" edge to the User entity by ID if the given value is not nil.
func (pu *PostUpdate) SetNillablePosterID(id *uuid.UUID) *PostUpdate {
	if id != nil {
		pu = pu.SetPosterID(*id)
	}
	return pu
}

// SetPoster sets the "poster" edge to the User entity.
func (pu *PostUpdate) SetPoster(u *User) *PostUpdate {
	return pu.SetPosterID(u.ID)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pu *PostUpdate) AddLikeIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.AddLikeIDs(ids...)
	return pu
}

// AddLikes adds the "likes" edges to the Like entity.
func (pu *PostUpdate) AddLikes(l ...*Like) *PostUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.AddLikeIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (pu *PostUpdate) AddFavouriteIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.AddFavouriteIDs(ids...)
	return pu
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (pu *PostUpdate) AddFavourites(f ...*Favourite) *PostUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.AddFavouriteIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearPoster clears the "poster" edge to the User entity.
func (pu *PostUpdate) ClearPoster() *PostUpdate {
	pu.mutation.ClearPoster()
	return pu
}

// ClearLikes clears all "likes" edges to the Like entity.
func (pu *PostUpdate) ClearLikes() *PostUpdate {
	pu.mutation.ClearLikes()
	return pu
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (pu *PostUpdate) RemoveLikeIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.RemoveLikeIDs(ids...)
	return pu
}

// RemoveLikes removes "likes" edges to Like entities.
func (pu *PostUpdate) RemoveLikes(l ...*Like) *PostUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.RemoveLikeIDs(ids...)
}

// ClearFavourites clears all "favourites" edges to the Favourite entity.
func (pu *PostUpdate) ClearFavourites() *PostUpdate {
	pu.mutation.ClearFavourites()
	return pu
}

// RemoveFavouriteIDs removes the "favourites" edge to Favourite entities by IDs.
func (pu *PostUpdate) RemoveFavouriteIDs(ids ...uuid.UUID) *PostUpdate {
	pu.mutation.RemoveFavouriteIDs(ids...)
	return pu
}

// RemoveFavourites removes "favourites" edges to Favourite entities.
func (pu *PostUpdate) RemoveFavourites(f ...*Favourite) *PostUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.RemoveFavouriteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks[int, PostMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Caption(); ok {
		_spec.SetField(post.FieldCaption, field.TypeString, value)
	}
	if value, ok := pu.mutation.Images(); ok {
		_spec.SetField(post.FieldImages, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedImages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldImages, value)
		})
	}
	if value, ok := pu.mutation.Location(); ok {
		_spec.SetField(post.FieldLocation, field.TypeString, value)
	}
	if pu.mutation.PosterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PosterTable,
			Columns: []string{post.PosterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PosterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PosterTable,
			Columns: []string{post.PosterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLikesIDs(); len(nodes) > 0 && !pu.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.FavouritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedFavouritesIDs(); len(nodes) > 0 && !pu.mutation.FavouritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCaption sets the "caption" field.
func (puo *PostUpdateOne) SetCaption(s string) *PostUpdateOne {
	puo.mutation.SetCaption(s)
	return puo
}

// SetImages sets the "images" field.
func (puo *PostUpdateOne) SetImages(s []string) *PostUpdateOne {
	puo.mutation.SetImages(s)
	return puo
}

// AppendImages appends s to the "images" field.
func (puo *PostUpdateOne) AppendImages(s []string) *PostUpdateOne {
	puo.mutation.AppendImages(s)
	return puo
}

// SetLocation sets the "location" field.
func (puo *PostUpdateOne) SetLocation(s string) *PostUpdateOne {
	puo.mutation.SetLocation(s)
	return puo
}

// SetPosterID sets the "poster" edge to the User entity by ID.
func (puo *PostUpdateOne) SetPosterID(id uuid.UUID) *PostUpdateOne {
	puo.mutation.SetPosterID(id)
	return puo
}

// SetNillablePosterID sets the "poster" edge to the User entity by ID if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePosterID(id *uuid.UUID) *PostUpdateOne {
	if id != nil {
		puo = puo.SetPosterID(*id)
	}
	return puo
}

// SetPoster sets the "poster" edge to the User entity.
func (puo *PostUpdateOne) SetPoster(u *User) *PostUpdateOne {
	return puo.SetPosterID(u.ID)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (puo *PostUpdateOne) AddLikeIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.AddLikeIDs(ids...)
	return puo
}

// AddLikes adds the "likes" edges to the Like entity.
func (puo *PostUpdateOne) AddLikes(l ...*Like) *PostUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.AddLikeIDs(ids...)
}

// AddFavouriteIDs adds the "favourites" edge to the Favourite entity by IDs.
func (puo *PostUpdateOne) AddFavouriteIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.AddFavouriteIDs(ids...)
	return puo
}

// AddFavourites adds the "favourites" edges to the Favourite entity.
func (puo *PostUpdateOne) AddFavourites(f ...*Favourite) *PostUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.AddFavouriteIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearPoster clears the "poster" edge to the User entity.
func (puo *PostUpdateOne) ClearPoster() *PostUpdateOne {
	puo.mutation.ClearPoster()
	return puo
}

// ClearLikes clears all "likes" edges to the Like entity.
func (puo *PostUpdateOne) ClearLikes() *PostUpdateOne {
	puo.mutation.ClearLikes()
	return puo
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (puo *PostUpdateOne) RemoveLikeIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.RemoveLikeIDs(ids...)
	return puo
}

// RemoveLikes removes "likes" edges to Like entities.
func (puo *PostUpdateOne) RemoveLikes(l ...*Like) *PostUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.RemoveLikeIDs(ids...)
}

// ClearFavourites clears all "favourites" edges to the Favourite entity.
func (puo *PostUpdateOne) ClearFavourites() *PostUpdateOne {
	puo.mutation.ClearFavourites()
	return puo
}

// RemoveFavouriteIDs removes the "favourites" edge to Favourite entities by IDs.
func (puo *PostUpdateOne) RemoveFavouriteIDs(ids ...uuid.UUID) *PostUpdateOne {
	puo.mutation.RemoveFavouriteIDs(ids...)
	return puo
}

// RemoveFavourites removes "favourites" edges to Favourite entities.
func (puo *PostUpdateOne) RemoveFavourites(f ...*Favourite) *PostUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.RemoveFavouriteIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks[*Post, PostMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Caption(); ok {
		_spec.SetField(post.FieldCaption, field.TypeString, value)
	}
	if value, ok := puo.mutation.Images(); ok {
		_spec.SetField(post.FieldImages, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedImages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, post.FieldImages, value)
		})
	}
	if value, ok := puo.mutation.Location(); ok {
		_spec.SetField(post.FieldLocation, field.TypeString, value)
	}
	if puo.mutation.PosterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PosterTable,
			Columns: []string{post.PosterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PosterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.PosterTable,
			Columns: []string{post.PosterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLikesIDs(); len(nodes) > 0 && !puo.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.FavouritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedFavouritesIDs(); len(nodes) > 0 && !puo.mutation.FavouritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FavouritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.FavouritesTable,
			Columns: []string{post.FavouritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
