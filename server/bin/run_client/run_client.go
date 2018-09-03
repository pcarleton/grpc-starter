package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nmrshll/oauth2-noserver"
	pb "github.com/pcarleton/grpc-starter/proto/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"os"
)

var address string
var cert string
var insecure bool

func init() {
	flag.StringVar(&address, "address", "", "Address and port to connect to")
	flag.StringVar(&cert, "cert", "certs/localhost.crt", "Path to the cert file to use for TLS")
	flag.BoolVar(&insecure, "insecure", false, "Run without certs")
}

type googleConfig struct {
	Web struct {
		ClientID                string   `json:"client_id"`
		ProjectID               string   `json:"project_id"`
		AuthUri                 string   `json:"auth_uri"`
		TokenUri                string   `json:"token_uri"`
		AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
		JavascriptOrigins       []string `json:"javascript_origins"`
	} `json:"web"`
}

func main() {
	flag.Parse()
	//Set up a connection to the server.
	secretsPath := os.Getenv("CLIENT_SECRETS_PATH")
	if secretsPath == "" {
		log.Fatalf("No client secrets specified under CLIENT_SECRETS_PATH")
	}

	b, err := ioutil.ReadFile(secretsPath)

	if err != nil {
		panic(err)
	}

	var configData googleConfig

	json.Unmarshal(b, &configData)

	fmt.Printf("%+v", configData)
	conf := &oauth2.Config{
		ClientID:     configData.Web.ClientID,     // also known as slient key sometimes
		ClientSecret: configData.Web.ClientSecret, // also known as secret key
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  configData.Web.AuthUri,
			TokenURL: configData.Web.TokenUri,
		},
	}
	fmt.Printf("%+v", conf)

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		panic(err)
	}
	var conn *grpc.ClientConn

	if insecure {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			panic(err)
		}

		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	}

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewApiClient(conn)

	md := metadata.Pairs("token", client.Token.Extra("id_token").(string))
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//ctx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()
	r, err := c.GetHealth(ctx, &pb.GetHealthRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: %s", r.Status)
}
