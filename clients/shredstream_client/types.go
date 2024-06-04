package shredstream_client

import (
	"github.com/FlowGoCrazy/jito-go/pkg"
	"github.com/FlowGoCrazy/jito-go/proto"
	"github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/grpc"
)

type client struct {
	GrpcConn *grpc.ClientConn
	RpcConn  *rpc.Client

	ShredstreamClient proto.ShredstreamClient

	Auth *pkg.AuthenticationService
}
