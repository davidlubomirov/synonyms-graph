package synonymsdatastore

import (
	"reflect"
	"sync"
	"testing"

	"github.com/emirpasic/gods/sets/hashset"
)

func Test_synonymsDataStore_Get(t *testing.T) {
	type fields struct {
		data map[string]*hashset.Set
		mu   sync.Mutex
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "test with empty data store",
			fields: fields{
				data: make(map[string]*hashset.Set),
			},
			args: args{
				word: "A",
			},
			want: []string{},
		},
		{
			name: "test with one existing word synonym in the data store",
			fields: fields{
				data: func() map[string]*hashset.Set {
					dataStore := make(map[string]*hashset.Set)

					allSets := hashset.New()
					allSets.Add("A")
					allSets.Add("B")

					dataStore["A"] = allSets
					dataStore["B"] = allSets

					return dataStore
				}(),
			},
			args: args{
				word: "A",
			},
			want: []string{"B"},
		},
		{
			name: "test with with two existing word synonyms in the data store",
			fields: fields{
				data: func() map[string]*hashset.Set {
					dataStore := make(map[string]*hashset.Set)

					allSets := hashset.New()
					allSets.Add("A")
					allSets.Add("B")
					allSets.Add("C")

					dataStore["A"] = allSets
					dataStore["B"] = allSets
					dataStore["C"] = allSets

					return dataStore
				}(),
			},
			args: args{
				word: "A",
			},
			want: []string{"B", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &synonymsDataStore{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			// DeepEqual not sufficient !
			if got := store.Get(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("synonymsDataStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_synonymsDataStore_Set(t *testing.T) {
	type fields struct {
		data map[string]*hashset.Set
		mu   sync.Mutex
	}
	type args struct {
		word    string
		synonym string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "add of new word but with existing synonym in the data store",
			fields: fields{
				data: func() map[string]*hashset.Set {
					dataStore := make(map[string]*hashset.Set)

					allSets := hashset.New()
					allSets.Add("A")
					allSets.Add("B")
					allSets.Add("C")

					dataStore["A"] = allSets
					dataStore["B"] = allSets
					dataStore["C"] = allSets

					return dataStore
				}(),
			},
			args: args{
				word:    "D",
				synonym: "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &synonymsDataStore{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			store.Set(tt.args.word, tt.args.synonym)

			if !reflect.DeepEqual(store.data[tt.args.word], store.data[tt.args.synonym]) {
				t.Errorf("synonymsDataStore.Set() - invalid data")
			}
		})
	}
}
