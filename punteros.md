# Manual de Punteros en Go

## ¿Qué son los punteros?

Un **puntero** es una variable que almacena la **dirección de memoria** de otra variable, no el valor directamente.

```go
var x int = 42
var p *int = &x  // p apunta a la dirección de x
```

## Operadores principales

### `&` - Operador de dirección (address-of)
Obtiene la dirección de memoria de una variable.

```go
var numero int = 10
var direccion *int = &numero  // direccion = 0x1040a124 (ejemplo)
```

### `*` - Operador de desreferenciación
Tiene dos usos:

#### 1. Declaración de tipo puntero
```go
var p *int        // p es un puntero a int
var q *string     // q es un puntero a string
```

#### 2. Desreferenciación (acceso al valor)
```go
var x int = 42
var p *int = &x
fmt.Println(*p)   // 42 (accede al valor)
*p = 100          // modifica el valor original
```

## Paso por valor vs paso por referencia

### Paso por valor (copia)
```go
func modificar(x int) {
    x = 100  // Solo modifica la copia
}

func main() {
    numero := 42
    modificar(numero)
    fmt.Println(numero)  // 42 (sin cambios)
}
```

### Paso por referencia (puntero)
```go
func modificarRef(x *int) {
    *x = 100  // Modifica el valor original
}

func main() {
    numero := 42
    modificarRef(&numero)  // Pasa la dirección
    fmt.Println(numero)    // 100 (modificado)
}
```

## Ejemplo práctico: función swap

```go
func swap(x, y *int) {
    temp := *x  // temp = valor de x
    *x = *y     // x toma el valor de y
    *y = temp   // y toma el valor temp (original de x)
}

func main() {
    a, b := 10, 20
    fmt.Printf("Antes: a=%d, b=%d\n", a, b)  // Antes: a=10, b=20
    
    swap(&a, &b)  // Pasamos las direcciones
    fmt.Printf("Después: a=%d, b=%d\n", a, b)  // Después: a=20, b=10
}
```

## Punteros con structs

### Por valor (copia)
```go
type Persona struct {
    Nombre string
    Edad   int
}

func cumplirAnos(p Persona) {
    p.Edad++  // Solo modifica la copia
}

func main() {
    juan := Persona{Nombre: "Juan", Edad: 25}
    cumplirAnos(juan)
    fmt.Println(juan.Edad)  // 25 (sin cambios)
}
```

### Por referencia (puntero)
```go
func cumplirAnosRef(p *Persona) {
    p.Edad++  // Modifica el original
    // Go permite p.Edad en vez de (*p).Edad
}

func main() {
    juan := Persona{Nombre: "Juan", Edad: 25}
    cumplirAnosRef(&juan)
    fmt.Println(juan.Edad)  // 26 (modificado)
}
```

## Puntero nulo (nil)

```go
var p *int
fmt.Println(p)        // <nil>
fmt.Println(p == nil) // true

// ¡CUIDADO! Esto causa panic
// fmt.Println(*p)    // panic: runtime error
```

## Cuándo usar punteros

### ✅ Usa punteros cuando:
- Necesitas modificar el valor original
- Quieres evitar copiar structs grandes
- Quieres compartir datos entre funciones
- Trabajas con métodos que modifican el receptor

### ❌ Evita punteros cuando:
- Solo necesitas leer el valor
- Trabajas con tipos básicos pequeños (int, bool, etc.)
- Quieres inmutabilidad

## Reglas importantes

1. **Los punteros no se pueden hacer aritmética** (como en C)
2. **Go maneja la memoria automáticamente** (garbage collector)
3. **Un puntero a puntero es válido**: `var p **int`
4. **Los slices, maps y channels son referencias por naturaleza**

## Ejemplos de memoria

```go
func main() {
    x := 42
    p := &x
    
    fmt.Printf("Valor de x: %d\n", x)           // 42
    fmt.Printf("Dirección de x: %p\n", &x)      // 0x1040a124
    fmt.Printf("Valor de p: %p\n", p)           // 0x1040a124
    fmt.Printf("Valor apuntado por p: %d\n", *p) // 42
    fmt.Printf("Dirección de p: %p\n", &p)      // 0x1040a130
}
```

## Consejos prácticos

- **Usa nombres descriptivos**: `userPtr` mejor que `p`
- **Siempre verifica nil** antes de desreferenciar
- **Recuerda**: `&` obtiene dirección, `*` accede al valor
- **Para structs grandes**, usa punteros por eficiencia
- **Para structs pequeños**, valor está bien