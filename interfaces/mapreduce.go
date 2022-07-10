package interfaces

import (
	mdl "github.com/sebmaspd/rnd/ieq/models"
)

// MapReduceSimple has simple map reduce functions
type MapReduceSimple interface {
	// DoMap takes a function and applies it to each of the elements in the list and returns a list of all results
	DoMap(IDs []string, fn func(string) float64) (results []mdl.IDValue, err error)
	// DoResduce that takes the mapped result list and boils it down to a single result..
	DoReduce(mappedResults []mdl.IDValue, fn func(float64, float64) float64) (result mdl.IDValue, err error)
}
