package opcional

type Opcional[T any] struct {
	valor T
	vacio bool
}

func (o Opcional[T]) TieneValor() bool {
	return !o.vacio
}

func (o Opcional[T]) Consumir() (Valor T, tieneValor bool) {
	return o.valor, o.TieneValor()
}

func (o Opcional[T]) ValorOCero() T {
	if o.vacio {
		var cero T
		return cero
	}
	return o.valor
}

func (o Opcional[T]) ValorO(porDefecto T) T {
	if o.vacio {
		return porDefecto
	}
	return o.valor
}

func (o Opcional[T]) ValorOFunc(f func(argumentos []any) T, argumentos []any) T {
	if o.vacio {
		return f(argumentos)
	}
	return o.valor
}

func CrearOpcional[T any](valor T) Opcional[T] {
	return Opcional[T]{
		valor,
		false,
	}
}

func CrearOpcionalVacio[T any]() Opcional[T] {
	var cero T
	return Opcional[T]{
		cero,
		false,
	}
}

func NuevoOpcional[T any](valor T) *Opcional[T] {
	return &Opcional[T]{
		valor,
		false,
	}
}
