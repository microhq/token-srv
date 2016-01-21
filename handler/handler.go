package handler

import (
	"crypto/rand"
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/micro/token-srv/db"
	"github.com/micro/token-srv/proto/record"
	uuid "github.com/streadway/simpleuuid"
	"golang.org/x/net/context"
)

var (
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func random(i int) string {
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
	return ""
}

type Record struct{}

func (s *Record) Generate(ctx context.Context, req *record.GenerateRequest, rsp *record.GenerateResponse) error {
	name := random(16)
	id, err := uuid.NewTime(time.Now())
	if err != nil {
		return errors.InternalServerError("go.micro.srv.token.Generate", err.Error())
	}
	namespace := "default"
	if len(req.Namespace) > 0 {
		namespace = req.Namespace
	}
	tk := &record.Token{
		Id:        id.String(),
		Namespace: namespace,
		Name:      name,
		Created:   time.Now().Unix(),
		Updated:   time.Now().Unix(),
	}
	if err := db.Create(tk); err != nil {
		return errors.InternalServerError("go.micro.srv.token.Generate", err.Error())
	}
	rsp.Token = tk
	return nil
}

func (s *Record) Create(ctx context.Context, req *record.CreateRequest, rsp *record.CreateResponse) error {
	if err := db.Create(req.Token); err != nil {
		return errors.InternalServerError("go.micro.srv.token.Create", err.Error())
	}
	return nil
}

func (s *Record) Read(ctx context.Context, req *record.ReadRequest, rsp *record.ReadResponse) error {
	token, err := db.Read(req.Id)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.token.Read", err.Error())
	}
	rsp.Token = token
	return nil
}

func (s *Record) Update(ctx context.Context, req *record.UpdateRequest, rsp *record.UpdateResponse) error {
	if err := db.Update(req.Token); err != nil {
		return errors.InternalServerError("go.micro.srv.token.Update", err.Error())
	}
	return nil
}

func (s *Record) Delete(ctx context.Context, req *record.DeleteRequest, rsp *record.DeleteResponse) error {
	if err := db.Delete(req.Id); err != nil {
		return errors.InternalServerError("go.micro.srv.token.Delete", err.Error())
	}
	return nil
}

func (s *Record) Search(ctx context.Context, req *record.SearchRequest, rsp *record.SearchResponse) error {
	tokens, err := db.Search(req.Namespace, req.Name, req.Limit, req.Offset)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.token.Search", err.Error())
	}
	rsp.Tokens = tokens
	return nil
}
