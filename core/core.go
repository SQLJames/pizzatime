package core

import "context"

// Action is something that core can do
type Action interface {
	Run(ctx context.Context) error
	Config() Config
}