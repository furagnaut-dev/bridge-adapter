package legacy

type Pager struct {
	Online bool
	ID int
	Text string
}

const (
	StatusOK = 0
	StatusOffline = 1
	StatusBadID = 2
)

func (p *Pager) Page(id int, text []byte) int {
	if !p.Online {
		return StatusOffline
	}
	if id <= 0 {
		return StatusBadID
	}

	p.ID = id
	p.Text = string(text)

	return StatusOK
}