package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/pcarleton/grpc-starter/proto/api"
	server "github.com/pcarleton/grpc-starter/server"
	"flag"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

const (
	port = ":5001"
)

var cert string
var key string

func init() {
	flag.StringVar(&cert, "cert", "certs/localhost.crt", "Path to the cert file to use for TLS")
	flag.StringVar(&key, "key", "certs/localhost.key", "TLS cert private key")
}


func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if len(meta["token"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "no token sent")
	}

	token, err := auth.VerifyGoogleJwt(meta["token"][0])
	if err != nil {
		log.Printf("Error: %s", err)
		log.Printf("Token: %s", meta["token"][0])
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	log.Printf("%+v", token)

	if err = token.Claims.Valid(); err != nil {
		log.Printf("Error: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid claims")

	}

	ctype := reflect.TypeOf(token.Claims)
	log.Printf("Claims type: %v", ctype)
	log.Printf("Claims path: %v", ctype.PkgPath())

	for i := 0; i < ctype.NumMethod(); i++ {
		log.Printf("Method %i : %v", i, ctype.Method(i).PkgPath)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Printf("Invalid clamis: %+v", token.Claims)
		return nil, status.Errorf(codes.Unauthenticated, "uncastable claims")
	}
	email, ok := claims["email"].(string)
	name, ok := claims["name"].(string)

	log.Printf(email)
	log.Printf(name)

	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
    panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(AuthInterceptor))
	apiServer := server.NewServer()
	pb.RegisterCashCoachApiServer(s, apiServer)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
