package cache

import (
	"context"

	pb "github.com/andersnormal/outlaw/proto"
	"github.com/andersnormal/outlaw/provider"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

var _ autocert.Cache = (*cache)(nil)

func New(p provider.Provider) Cache { // interface

	// create logger
	l := log.WithFields(log.Fields{
		"provider": p.Name(),
	})

	return &cache{p, l}
}

func (c *cache) GetDomain(ctx context.Context, domain string) (*pb.Domain, error) {
	var (
		err  error
		d    *pb.Domain
		done = make(chan struct{})
		l    = c.l.WithFields(log.Fields{"host": domain})
	)

	go func() {
		defer close(done)
		d, err = c.p.GetDomain(ctx, &pb.Domain{Name: domain})
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
	}

	if err != nil {
		l.Error(err)
		return nil, err
	}

	return d, nil
}

func (c *cache) AllowHostPolicy(ctx context.Context, host string) error {
	var (
		err  error
		done = make(chan struct{})
		l    = c.l.WithFields(log.Fields{"host": host})
	)

	go func() {
		defer close(done)
		err = c.p.AllowHostPolicy(ctx, host)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	if err != nil {
		l.Error(err)
		return err
	}

	return nil
}

func (c *cache) Get(ctx context.Context, key string) ([]byte, error) {
	var (
		data []byte
		err  error
		done = make(chan struct{})
	)

	go func() {
		defer close(done)
		data, err = c.p.Get(ctx, key)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
	}

	if err != nil {
		return nil, autocert.ErrCacheMiss
	}

	return data, err
}

func (c *cache) Put(ctx context.Context, key string, data []byte) error {
	var (
		err  error
		done = make(chan struct{})
	)

	go func() {
		defer close(done)
		err = c.p.Put(ctx, key, data)
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return err
}

func (c *cache) Delete(ctx context.Context, key string) error {
	var (
		err  error
		done = make(chan struct{})
	)

	go func() {
		defer close(done)
		err = c.p.Delete(ctx, key)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return err
}

func (c *cache) provider() provider.Provider {
	return c.p
}
