package cluster

import (
	"fmt"
	"net"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"pitaya/config"
	"pitaya/helpers"
	"pitaya/metrics"
	protosmocks "pitaya/protos/mocks"
)

func TestNewGRPCServer(t *testing.T) {
	t.Parallel()
	sv := getServer()
	gs, err := NewGRPCServer(*config.NewDefaultGRPCServerConfig(), sv, []metrics.Reporter{})
	assert.NoError(t, err)
	assert.NotNil(t, gs)
}

func TestGRPCServerInit(t *testing.T) {
	t.Parallel()
	c := config.NewDefaultGRPCServerConfig()
	c.Port = helpers.GetFreePort(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPitayaServer := protosmocks.NewMockPitayaServer(ctrl)

	sv := getServer()
	gs, err := NewGRPCServer(*c, sv, []metrics.Reporter{})
	gs.SetPitayaServer(mockPitayaServer)
	err = gs.Init()
	assert.NoError(t, err)
	assert.NotNil(t, gs)

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", c.Port))
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NotNil(t, gs.grpcSv)
}
