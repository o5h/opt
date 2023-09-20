package opt

type O[T any] struct{ V *T }
type Else struct {
	enable bool
}

func (e *Else) Else(fn func()) {
	if e.enable {
		fn()
	}
}

func Of[T any](t *T) O[T] { return O[T]{V: t} }

func At[T any](slice []*T, i int) O[T] {
	if i >= len(slice) {
		return O[T]{V: nil}
	}
	return O[T]{V: slice[i]}
}

func (o O[T]) Ok() bool { return o.V != nil }

func (o O[T]) Or(v T) T {
	if o.V != nil {
		return *o.V
	}
	return v
}

func (o O[T]) IfOk(fn func(*T)) *Else {
	if o.V != nil {
		fn(o.V)
		return &Else{enable: false}
	}
	return &Else{enable: false}
}

func (o O[T]) IfNil(fn func(*T)) *Else {
	if o.V == nil {
		fn(o.V)
		return &Else{enable: false}
	}
	return &Else{enable: true}
}
