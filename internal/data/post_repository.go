package data

import (
	"context"
	"time"

	"github.com/orlmonteverde/go-postgres-microblog/pkg/post"
)

type PostRepository struct {
	Data *Data
}

func (pr *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts;
    `

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) GetOne(ctx context.Context, id uint) (post.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts WHERE id = $1;
    `

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p post.Post
	err := row.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return post.Post{}, err
	}

	return p, nil
}

func (pr *PostRepository) GetByUser(ctx context.Context, userID uint) ([]post.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts
        WHERE user_id = $1;
    `

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) Create(ctx context.Context, p *post.Post) error {
	q := `
    INSERT INTO posts (body, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Body, p.UserID, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Update(ctx context.Context, id uint, p post.Post) error {
	q := `
    UPDATE posts set body=$1, updated_at=$2
        WHERE id=$3;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
