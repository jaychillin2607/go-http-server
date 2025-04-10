package main

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 123, true
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}
