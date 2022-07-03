package synonymsdatastore

import (
	"fmt"
	"sync"

	"github.com/emirpasic/gods/sets/hashset"
)

// case-insensitive in Get and Set not needed for the purpose of this task

type synonymsDataStore struct {
	data map[string]*hashset.Set
	mu   sync.Mutex
}

func NewSynonymsStore() *synonymsDataStore {
	return &synonymsDataStore{
		data: make(map[string]*hashset.Set),
	}
}

func (store *synonymsDataStore) Get(word string) []string {
	store.mu.Lock()
	defer store.mu.Unlock()

	var result = []string{}

	listPointer, exist := store.data[word]
	if exist {
		list := listPointer.Values()

		for _, val := range list {
			tempVal := fmt.Sprint(val)
			if tempVal != word {
				result = append(result, tempVal)
			}
		}

		return result
	}

	return result
}

func (store *synonymsDataStore) Set(word, synonym string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	wordPointerSet, wordExist := store.data[word]
	synonymPointerSet, synonymExist := store.data[synonym]

	if !wordExist && synonymExist {
		synonymPointerSet.Add(word)

		store.data[word] = synonymPointerSet
	}

	if wordExist && !synonymExist {
		wordPointerSet.Add(synonym)

		store.data[synonym] = wordPointerSet
	}

	if !wordExist && !synonymExist {
		newSet := hashset.New()
		newSet.Add(word)
		newSet.Add(synonym)

		store.data[word] = newSet
		store.data[synonym] = newSet
	}
}
