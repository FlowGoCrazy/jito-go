package relayer_client

import (
	"context"
	jito_go "github.com/FlowGoCrazy/jito-go"
	"github.com/FlowGoCrazy/jito-go/proto"
	"github.com/gagliardetto/solana-go"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	godotenv.Load(filepath.Join(filepath.Dir(filename), "..", "..", "..", "jito-go", ".env"))
	os.Exit(m.Run())
}

func Test_RelayerClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	privKey, ok := os.LookupEnv("PRIVATE_KEY")
	if !assert.True(t, ok, "getting PRIVATE_KEY from .env") {
		t.FailNow()
	}

	client, err := New(
		ctx,
		jito_go.Amsterdam.BlockEngineURL,
		solana.MustPrivateKeyFromBase58(privKey),
		nil,
	)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	defer client.GrpcConn.Close()

	t.Run("GetTpuConfig", func(t *testing.T) {
		resp, err := client.GetTpuConfigs()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		if assert.NotEqual(t, "", resp.Tpu.Ip) {
		}
	})

	t.Run("SubscribePacket", func(t *testing.T) {
		sub, err := client.NewPacketsSubscription()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		var recv *proto.SubscribePacketsResponse
		recv, err = sub.Recv()
		assert.NoError(t, err)

		recv.Header.String()
	})
}
