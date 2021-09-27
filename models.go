package main

import "time"

type Articolo struct {
	ID                int
	DataPubblicazione time.Time
	TitoloIt          string
	TitoloEn          string
	TestoIt           string
	TestoEn           string
	AutoreId          int
}
