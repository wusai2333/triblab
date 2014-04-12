package triblab

import (
	"trib"
)

import "tribsol"

func NewBinClient(backs []string) trib.BinStorage {
	return tribsol.NewBinClient(backs)
}

func ServeKeeper(kc *trib.KeeperConfig) error {
	return tribsol.ServeKeeper(kc)
}

func NewFront(s trib.BinStorage) trib.Server {
	return tribsol.NewFront(s)
}

