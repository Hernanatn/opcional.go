<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# opcional

```go
import "github.com/Hernanatn/opcional.go"
```

## Index

- [type Opcional](<#Opcional>)
  - [func CrearOpcional\[T any\]\(valor T\) Opcional\[T\]](<#CrearOpcional>)
  - [func CrearOpcionalVacio\[T any\]\(\) Opcional\[T\]](<#CrearOpcionalVacio>)
  - [func NuevoOpcional\[T any\]\(valor T\) \*Opcional\[T\]](<#NuevoOpcional>)
  - [func \(o Opcional\[T\]\) Consumir\(\) \(Valor T, tieneValor bool\)](<#Opcional[T].Consumir>)
  - [func \(o Opcional\[T\]\) TieneValor\(\) bool](<#Opcional[T].TieneValor>)
  - [func \(o Opcional\[T\]\) ValorO\(porDefecto T\) T](<#Opcional[T].ValorO>)
  - [func \(o Opcional\[T\]\) ValorOCero\(\) T](<#Opcional[T].ValorOCero>)
  - [func \(o Opcional\[T\]\) ValorOFunc\(f func\(argumentos \[\]any\) T, argumentos \[\]any\) T](<#Opcional[T].ValorOFunc>)


<a name="Opcional"></a>
## type Opcional



```go
type Opcional[T any] struct {
	valor T
	vacio bool
}
```

<a name="CrearOpcional"></a>
### func CrearOpcional

```go
func CrearOpcional[T any](valor T) Opcional[T]
```



<a name="CrearOpcionalVacio"></a>
### func CrearOpcionalVacio

```go
func CrearOpcionalVacio[T any]() Opcional[T]
```



<a name="NuevoOpcional"></a>
### func NuevoOpcional

```go
func NuevoOpcional[T any](valor T) *Opcional[T]
```



<a name="Opcional[T].Consumir"></a>
### func \(Opcional\[T\]\) Consumir

```go
func (o Opcional[T]) Consumir() (Valor T, tieneValor bool)
```



<a name="Opcional[T].TieneValor"></a>
### func \(Opcional\[T\]\) TieneValor

```go
func (o Opcional[T]) TieneValor() bool
```



<a name="Opcional[T].ValorO"></a>
### func \(Opcional\[T\]\) ValorO

```go
func (o Opcional[T]) ValorO(porDefecto T) T
```



<a name="Opcional[T].ValorOCero"></a>
### func \(Opcional\[T\]\) ValorOCero

```go
func (o Opcional[T]) ValorOCero() T
```



<a name="Opcional[T].ValorOFunc"></a>
### func \(Opcional\[T\]\) ValorOFunc

```go
func (o Opcional[T]) ValorOFunc(f func(argumentos []any) T, argumentos []any) T
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
