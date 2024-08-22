/*
Copyright © 2024 Kai Blin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"context"
	"database/sql"
	"time"
)

const TIMEOUT = time.Second * 5

type Link struct {
	Short string
	Url   string
}

type DbModel interface {
	Init() error
	Add(short, url string) error
	Update(short, url string) error
	Remove(short string) error
	List() ([]Link, error)
	Get(short string) (string, error)
}

type ShortisModel struct {
	db *sql.DB
}

func NewShortisModel(db *sql.DB) ShortisModel {
	return ShortisModel{db: db}
}

func (m *ShortisModel) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	_, err := m.db.ExecContext(ctx, `
CREATE TABLE shortcuts (
	short TEXT NOT NULL PRIMARY KEY,
	url TEXT NOT NULL
);
	`)
	return err
}

func (m *ShortisModel) Add(short, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	_, err := m.db.ExecContext(ctx, `
INSERT INTO shortcuts (short, url) VALUES (?, ?)
		`, short, url)

	return err
}

func (m *ShortisModel) Update(short, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	_, err := m.db.ExecContext(ctx, `
UPDATE shortcuts SET url=? WHERE short=?
		`, short, url)

	return err
}

func (m *ShortisModel) Remove(short string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	_, err := m.db.ExecContext(ctx, `
DELETE FROM shortcuts WHERE short = ?
		`, short)

	return err
}

func (m *ShortisModel) List() ([]Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	rows, err := m.db.QueryContext(ctx, `
SELECT short, url FROM shortcuts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]Link, 0)
	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.Short, &link.Url); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func (m *ShortisModel) Get(short string) (string, error) {
	var url string

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	row := m.db.QueryRowContext(ctx, `
SELECT url FROM shortcuts WHERE short=?
	`, short)

	err := row.Scan(&url)
	return url, err
}
