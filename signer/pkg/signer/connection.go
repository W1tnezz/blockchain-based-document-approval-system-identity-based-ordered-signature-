package signer

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ConnectionManager struct {
	sync.RWMutex
	BatchVerifier *BatchVerifier
	account        common.Address
	connections    map[common.Address]*grpc.ClientConn
}

func NewConnectionManager(BatchVerifier *BatchVerifier, account common.Address) *ConnectionManager {
	return &ConnectionManager{
		BatchVerifier: BatchVerifier,
		account:        account,
		connections:    make(map[common.Address]*grpc.ClientConn),
	}
}

// 创建连接,使用数组保存,对每一个节点使用数组保存连接状态
func (m *ConnectionManager) NewConnection(node  OracleContractOracleNode) (*grpc.ClientConn, error) {
	m.Lock()
	// defer是延迟执行语句,当Windows界面的语句都执行完,才会执行defer,并且多个defer之间使用逆序执行顺序
	defer m.Unlock()
	if conn, ok := m.connections[node.Addr]; ok {
		return conn, nil
	}
	conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial %s: %v", node.IpAddr, err)
	}
	m.connections[node.Addr] = conn

	log.Infof("New connection to %s with index %d and IP address %s", node.Addr, node.Index, node.IpAddr)

	return conn, nil
}

func (m *ConnectionManager) FindByAddress(address common.Address) (*grpc.ClientConn, error) {
	m.RLock()
	defer m.RUnlock()
	if conn, ok := m.connections[address]; ok {
		return conn, nil
	}
	return nil, fmt.Errorf("connection to node with address %s not found", address)
}

func (m *ConnectionManager) Close() {
	for _, conn := range m.connections {
		_ = conn.Close()
	}
}
