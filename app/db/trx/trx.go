package trx

import (
	"context"
	"github.com/jinzhu/gorm"
	"sync"
)

const prefix = "app.db.trx"

type Trx struct {
	db   *gorm.DB
	done context.CancelFunc
	mu   sync.Mutex
}

// Rollback
func (t *Trx) Rollback() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.db.Rollback()
	t.done()
	return t.db.Error
}

// Commit
func (t *Trx) Commit() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.db.Commit()
	t.done()
	return t.db.Error
}

// Manager
type Manager struct {
	db *gorm.DB
}

// Begin
func (t *Manager) Begin(ctx context.Context) (trx *Trx, c context.Context) {
	trx = &Trx{}
	trx.db = t.db.Begin()
	c = context.WithValue(ctx, prefix, trx)
	c, trx.done = context.WithCancel(c)
	return trx, c
}

// Inject
func Inject(ctx context.Context, db *gorm.DB) *gorm.DB {
	if t, ok := ctx.Value(prefix).(*Trx); ok {
		return t.db
	}
	return db
}

// NewTrxManager
func NewTrxManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}
