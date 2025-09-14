package main 

type Sorter struct {
	sortStrategy SorterStrategy 
}

func NewSorter(sortStrategy SorterStrategy) *Sorter {
	return &Sorter{
		sortStrategy: sortStrategy,
	}
}

func (s *Sorter) SetSortStrategy(sortStrategy SorterStrategy) {
	s.sortStrategy = sortStrategy 
}

func (s *Sorter) Sort(arr []int ) {
	s.sortStrategy.sort(arr) 
}
