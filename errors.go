package emb

import "fmt"

type MissingTop struct{}

func (e MissingTop) Error() string {
	return fmt.Sprintf("error: emb.Top has not been assigned by developer")
}

type MissingFS struct{}

func (e MissingFS) Error() string {
	return fmt.Sprintf("error: emb.FS has not been assigned by developer")
}
