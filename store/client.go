package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Jasun0925/go-oauth2-mysql"
	"github.com/Jasun0925/go-oauth2-mysql/models"
	"sync"
)

// NewClientStore create client store
func NewClientStore() *ClientStore {
	return &ClientStore{
		data: make(map[string]oauth2.ClientInfo),
	}
}

// NewClientMySqlStore TODO 改动1，mysql存储 改在初始化存储时
func NewClientMySqlStore(dbs *sql.DB, table string) *ClientStore {
	return &ClientStore{
		db:        dbs,
		tableName: table,
		data:      make(map[string]oauth2.ClientInfo),
	}
}

// ClientStore client information store
type ClientStore struct {
	db        *sql.DB
	tableName string

	sync.RWMutex
	data map[string]oauth2.ClientInfo
}

// TODO，改动2，新添加了两个数据，一个是传入的mysql连接，一个是表名
// GetByID according to the ID for the client information
func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	// 如果不是使用mysql，就使用默认存储
	if cs.db == nil {
		// 通过查询clientStore里的data的id，返回clientInfo
		if c, ok := cs.data[id]; ok {
			return c, nil
		}
		return nil, errors.New("not found")
	}
	// mysql里查询数据
	clientSql := fmt.Sprintf("select * from %v where id = %v", cs.tableName, id)
	row, err := cs.db.Query(clientSql)
	if err != nil {
		return nil, err
	}
	// 解析数据，没有使用sqlx
	var client models.Client
	for row.Next() {
		err = row.Scan(&client.ID, &client.Secret, &client.Domain, &client.UserID)
		if err != nil {
			return nil, err
		}
	}
	if client.ID == "" {
		return nil, errors.New("not found client")
	}

	// 返回数据
	return &client, nil
}

// Set set client information
func (cs *ClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()

	cs.data[id] = cli
	return
}
