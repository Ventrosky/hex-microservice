package shortener

type RedirecRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
