package sipope

// Representa recursos paginados, ex: Risps, Unidades, etc.
type Paginated[T any] struct {
	CurrentPage  int     `json:"current_page"`
	Total        int     `json:"total"`
	Path         string  `json:"path"`
	PrevPageURL  *string `json:"prev_page_url"`
	NextPageURL  *string `json:"next_page_url"`
	FirstPageURL string  `json:"first_page_url"`
	LastPageURL  string  `json:"last_page_url"`
	Data         []T     `json:"data"`
	LastPage     int     `json:"last_page"`
	PerPage      int     `json:"per_page"`
	From         int     `json:"from"`
	To           int     `json:"to"`
}

func (p *Paginated[T]) GetNextPageURL() string {
	if p.NextPageURL == nil {
		return ""
	}
	return *p.NextPageURL
}

func (p *Paginated[T]) HasNextPageURL() bool {
	return p.NextPageURL != nil
}

func (p *Paginated[T]) GetPrevPageURL() string {
	if p.PrevPageURL == nil {
		return ""
	}
	return *p.PrevPageURL
}

func (p *Paginated[T]) HasPrevPageURL() bool {
	return p.PrevPageURL != nil
}

type Link struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

func (l *Link) HasURL() bool {
	return l.URL != nil
}

func (l *Link) GetURL() string {
	if l.URL == nil {
		return ""
	}
	return *l.URL
}
