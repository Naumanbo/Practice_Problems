package main

import (
	"fmt"
	"sort"
	"strings"
)

// Tests: Maps, nested maps, sorting map results, real-world usage patterns
//
// Build a simple in-memory record store using maps.
//
// 1. Define a Record struct: ID int, Tags []string, Score int
//
// 2. Implement a Store struct with these methods:
//    - Add(record Record)                    - add a record
//    - Get(id int) (Record, bool)            - retrieve by ID
//    - Delete(id int) bool                   - delete by ID, return false if not found
//    - AllByTag(tag string) []Record         - return all records with given tag, sorted by Score descending
//    - TopN(n int) []Record                  - return top n records by Score (sorted descending)
//    - TagCounts() map[string]int            - count how many records have each tag
//    - AverageScore() float64                - average score across all records (0.0 if empty)
//
// The tricky parts:
// - Store must maintain BOTH an ID index AND a tag index for O(1) lookups
// - Delete must clean up BOTH indexes
// - Think about what data structure the tag index should be

type Record struct {
	ID    int
	Tags  []string
	Score int
}

type Store struct {
	// You need at least two maps here
	// Think: map[int]Record for ID lookup, map[string]??? for tag lookup
	records    map[int]Record
	tags       map[string][]int
	numRecords int
}

// NewStore creates an initialized Store
func NewStore() *Store {
	return &Store{records: make(map[int]Record), tags: make(map[string][]int), numRecords: 0}
}

// TODO: Implement Add
func (s *Store) Add(record Record) {
	s.records[record.ID] = record
	for _, tag := range record.Tags {
		s.tags[tag] = append(s.tags[tag], record.ID)
	}
	s.numRecords++
}

// TODO: Implement Get
func (s *Store) Get(id int) (Record, bool) {
	record, ok := s.records[id]
	if ok == true {
		return record, true
	}
	return Record{}, false

}

// TODO: Implement Delete
func (s *Store) Delete(id int) bool {
	record, ok := s.records[id]
	if ok == true {
		delete(s.records, id)
		// TODO: handle deleting the id from all tags that have the id {HOW??}
		recordTags := record.Tags
		for _, tag := range recordTags { // remove from all tags containing id
			for i, v := range s.tags[tag] {
				if v == id {
					newTags := []int{}
					newTags = append(newTags, s.tags[tag][:i]...)
					newTags = append(newTags, s.tags[tag][i+1:]...)
					s.tags[tag] = newTags // remove index of the id from tags map
				}
			}
		}
		s.numRecords--
		return true
	}

	return false

}

// TODO: Implement AllByTag (sorted by Score descending)
func (s *Store) AllByTag(tag string) []Record {
	tagList, ok := s.tags[tag]
	if ok && len(tagList) > 0 {
		recordList := make([]Record, 0)
		for _, id := range tagList {
			recordList = append(recordList, s.records[id])
		}

		sort.Slice(recordList, func(i, j int) bool { //in place sorting a slice using a comparator/less function
			return recordList[i].Score > recordList[j].Score
		})

		return recordList

	}
	return make([]Record, 0)
}

// TODO: Implement TopN (sorted by Score descending)
func (s *Store) TopN(n int) []Record {
	topNSlice := make([]Record, 0)

	for recordID := range s.records {
		topNSlice = append(topNSlice, s.records[recordID])
	}

	sort.Slice(topNSlice, func(i, j int) bool {
		return topNSlice[i].Score > topNSlice[j].Score
	})

	if n > len(topNSlice) {
		return topNSlice
	} else {
		return topNSlice[:n]
	}
}

// TODO: Implement TagCounts
func (s *Store) TagCounts() map[string]int {
	tagCounts := make(map[string]int, 0)
	for tag := range s.tags {
		tagCounts[tag] = len(s.tags[tag])
	}

	return tagCounts
}

// TODO: Implement AverageScore
func (s *Store) AverageScore() float64 {
	var avgScore float64 = 0.0
	aggregateScore := 0
	for key := range s.records {
		aggregateScore += s.records[key].Score
	}
	if s.numRecords != 0 {
		avgScore = float64(aggregateScore) / float64(s.numRecords)
	}
	return avgScore
}

func main() {
	store := NewStore()

	// Add records
	store.Add(Record{ID: 1, Tags: []string{"go", "backend"}, Score: 95})
	store.Add(Record{ID: 2, Tags: []string{"go", "cli"}, Score: 80})
	store.Add(Record{ID: 3, Tags: []string{"python", "ml"}, Score: 90})
	store.Add(Record{ID: 4, Tags: []string{"go", "backend", "api"}, Score: 85})
	store.Add(Record{ID: 5, Tags: []string{"python", "backend"}, Score: 70})

	// Test Get
	fmt.Println("=== Get ===")
	if r, ok := store.Get(1); ok {
		fmt.Printf("ID 1: %+v\n", r)
	}
	if _, ok := store.Get(99); !ok {
		fmt.Println("ID 99: not found (correct)")
	}

	// Test AllByTag
	fmt.Println("\n=== AllByTag ===")
	goRecords := store.AllByTag("go")
	fmt.Println("Tag 'go' (sorted by score desc):")
	for _, r := range goRecords {
		fmt.Printf("  ID=%d Score=%d\n", r.ID, r.Score)
	}
	// Expected order: ID=1(95), ID=4(85), ID=2(80)

	backendRecords := store.AllByTag("backend")
	fmt.Println("Tag 'backend':")
	for _, r := range backendRecords {
		fmt.Printf("  ID=%d Score=%d\n", r.ID, r.Score)
	}
	// Expected order: ID=1(95), ID=4(85), ID=5(70)

	// Test TopN
	fmt.Println("\n=== TopN ===")
	top3 := store.TopN(3)
	fmt.Println("Top 3:")
	for _, r := range top3 {
		fmt.Printf("  ID=%d Score=%d\n", r.ID, r.Score)
	}
	// Expected: ID=1(95), ID=3(90), ID=4(85)

	// Test TagCounts
	fmt.Println("\n=== TagCounts ===")
	counts := store.TagCounts()
	fmt.Println(counts) // map[api:1 backend:3 cli:1 go:3 ml:1 python:2]

	// Test AverageScore
	fmt.Println("\n=== AverageScore ===")
	fmt.Printf("Average: %.1f\n", store.AverageScore()) // 84.0

	// Test Delete
	fmt.Println("\n=== Delete ===")
	store.Delete(2)
	goAfterDelete := store.AllByTag("go")
	fmt.Println("Tag 'go' after deleting ID 2:")
	for _, r := range goAfterDelete {
		fmt.Printf("  ID=%d Score=%d\n", r.ID, r.Score)
	}
	// Expected: ID=1(95), ID=4(85) — ID 2 gone
	fmt.Println("Tag 'cli' after deleting ID 2:", store.AllByTag("cli")) // []

	// Run test cases
	allPassed := true

	// Empty store
	empty := NewStore()
	if empty.AverageScore() != 0 {
		fmt.Println("FAIL: AverageScore on empty store")
		allPassed = false
	}
	if len(empty.TopN(5)) != 0 {
		fmt.Println("FAIL: TopN on empty store")
		allPassed = false
	}
	if len(empty.AllByTag("anything")) != 0 {
		fmt.Println("FAIL: AllByTag on empty store")
		allPassed = false
	}

	// Add and Get
	s := NewStore()
	s.Add(Record{ID: 1, Tags: []string{"a"}, Score: 50})
	if r, ok := s.Get(1); !ok || r.Score != 50 {
		fmt.Println("FAIL: Add then Get")
		allPassed = false
	}

	// Delete nonexistent
	if s.Delete(999) {
		fmt.Println("FAIL: Delete nonexistent should return false")
		allPassed = false
	}

	// Delete cleans up tag index
	s2 := NewStore()
	s2.Add(Record{ID: 1, Tags: []string{"x", "y"}, Score: 10})
	s2.Delete(1)
	if len(s2.AllByTag("x")) != 0 {
		fmt.Println("FAIL: Delete should clean tag index")
		allPassed = false
	}
	if _, ok := s2.Get(1); ok {
		fmt.Println("FAIL: Get after Delete should return false")
		allPassed = false
	}

	// TopN with n > number of records
	s3 := NewStore()
	s3.Add(Record{ID: 1, Tags: []string{}, Score: 10})
	s3.Add(Record{ID: 2, Tags: []string{}, Score: 20})
	top := s3.TopN(10)
	if len(top) != 2 || top[0].Score != 20 {
		fmt.Println("FAIL: TopN n > count")
		allPassed = false
	}

	// AllByTag is sorted descending by score
	s4 := NewStore()
	s4.Add(Record{ID: 1, Tags: []string{"t"}, Score: 10})
	s4.Add(Record{ID: 2, Tags: []string{"t"}, Score: 30})
	s4.Add(Record{ID: 3, Tags: []string{"t"}, Score: 20})
	tagged := s4.AllByTag("t")
	if len(tagged) != 3 || tagged[0].Score != 30 || tagged[1].Score != 20 || tagged[2].Score != 10 {
		fmt.Println("FAIL: AllByTag sort order")
		allPassed = false
	}

	// TagCounts accuracy
	s5 := NewStore()
	s5.Add(Record{ID: 1, Tags: []string{"a", "b"}, Score: 1})
	s5.Add(Record{ID: 2, Tags: []string{"b", "c"}, Score: 2})
	tc := s5.TagCounts()
	if tc["a"] != 1 || tc["b"] != 2 || tc["c"] != 1 {
		fmt.Println("FAIL: TagCounts")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: sort and strings packages imported for your use
var _ = sort.Slice
var _ = strings.Join
