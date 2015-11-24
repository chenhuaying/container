package container

type Comparer interface {
	Less(x Comparer) bool
}
