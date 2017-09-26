package service

var defaultResultsPerPage = 10

//Pageable ...
type Pageable struct {
	Limt           int
	Offset         int
	ResultsPerPage int
}

//SetResultsPerPage ...
func (p *Pageable) SetResultsPerPage(num int) *Pageable {

	p.ResultsPerPage = num
	return p
}
