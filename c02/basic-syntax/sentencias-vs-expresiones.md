# Sentencias vs Expresiones en Go

## Definiciones

### Expresión
Una **expresión** produce un valor que puede ser evaluado y usado en otros contextos.

```go
5 + 3                    // expresión → 8
x * 2                    // expresión → valor de x multiplicado por 2
i < 10                   // expresión → true o false
func() int { return 42 }() // expresión → 42
```

### Sentencia
Una **sentencia** ejecuta una acción pero no produce un valor utilizable.

```go
var x int                    // sentencia de declaración
x = 5                        // sentencia de asignación
if x > 0 { }                 // sentencia condicional
for i := 0; i < 10; i++ { }  // sentencia de bucle
return x                     // sentencia de retorno
```

## La "Escuela de C" en Go

Go hereda de C la **separación estricta** entre sentencias y expresiones. Esto significa:

### Estructuras de control son sentencias

```go
// ❌ Esto NO funciona en Go
x := if true { 1 } else { 2 }  // ERROR: if es sentencia, no expresión

// ✅ Forma correcta
var x int
if true {
    x = 1
} else {
    x = 2
}
```

### No hay operador ternario

```go
// ❌ No existe en Go (como sí existe en C/Java/JavaScript)
x := condition ? value1 : value2

// ✅ Debes usar if-else
var x int
if condition {
    x = value1
} else {
    x = value2
}
```

### Asignaciones no retornan valores

```go
// ❌ Esto NO funciona (común en C)
if x := getValue(); x != nil { }  // Sintaxis especial requerida

// ✅ Forma correcta en Go
if x := getValue(); x != nil { }  // Declaración seguida de condición
```

## Ejemplos Prácticos

### Bucles son sentencias

```go
// ❌ No se puede asignar un bucle
result := for i := 0; i < 10; i++ {
    sum += i
}

// ✅ Los bucles ejecutan, no producen valores
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

### Switch es sentencia

```go
// ❌ Switch no produce valores directamente
value := switch x {
case 1: "uno"
case 2: "dos"
}

// ✅ Usar funciones o asignaciones dentro
var value string
switch x {
case 1:
    value = "uno"
case 2:
    value = "dos"
}
```

## Comparación con otros lenguajes

### Lenguajes funcionales (Rust, Scala)
```rust
// En Rust - if es expresión
let x = if condition { 1 } else { 2 };

// En Scala - todo es expresión
val result = for (i <- 1 to 10) yield i * 2
```

### Go mantiene el enfoque imperativo
```go
// Go prefiere ser explícito
var result []int
for i := 1; i <= 10; i++ {
    result = append(result, i*2)
}
```

## Ventajas del enfoque de Go

1. **Claridad**: Es obvio qué es control de flujo y qué produce valores
2. **Consistencia**: Todas las estructuras de control funcionan igual
3. **Simplicidad**: Menos casos especiales que recordar
4. **Legibilidad**: El código se lee de forma secuencial

## Desventajas

1. **Verbosidad**: Más líneas de código para operaciones simples
2. **No hay operador ternario**: Para casos simples puede ser tedioso
3. **Menos funcional**: Dificulta algunos patrones funcionales

## Conclusión

Go deliberadamente mantiene la distinción C entre sentencias y expresiones para favorecer:
- **Simplicidad conceptual**
- **Código predecible**  
- **Facilidad de lectura**

Aunque esto significa sacrificar algo de concisión, Go prioriza la claridad y mantenibilidad del código.