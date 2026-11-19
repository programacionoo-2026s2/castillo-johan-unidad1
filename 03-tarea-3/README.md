# 03-tarea3

## Descripción

Este programa permite convertir unidades de almacenamiento entre Bytes (B), Kilobytes (KB), Megabytes (MB) y Gigabytes (GB). El usuario ejecuta el programa desde la línea de comandos indicando un valor, una unidad de origen y una unidad de destino.

## Cómo ejecutar

Ejemplo:

```bash
go run . 2048 KB MB
```

Resultado:

```text
2048 KB = 2.00 MB
```

Códigos de salida:

| Código | Significado |
|----------|----------|
| 0 | Ejecución correcta |
| 1 | Número incorrecto de argumentos |
| 2 | Valor inválido o negativo |
| 3 | Unidad no válida |

## Pruebas realizadas

| Comando | Resultado |
|----------|----------|
| go run . 2048 KB MB | 2048 KB = 2.00 MB |
| go run . 1 gb mb | 1 GB = 1024.00 MB |
| go run . 1500 B KB | 1500 B = 1.46 KB |
| go run . 0.5 MB KB | 0.5 MB = 512.00 KB |
| go run . abc KB MB | Valor inválido |
| go run . -5 KB MB | Valor inválido |
| go run . 5 TB MB | Unidad inválida |
| go run . 5 KB | Error de uso |

## Preguntas

### a) ¿Por qué el valor está en os.Args[1] y no en os.Args[0]?

En mi programa utilizo `os.Args[1]` porque `os.Args[0]` contiene el nombre del programa ejecutado. Los argumentos ingresados por el usuario comienzan desde la posición 1.

### b) ¿Por qué los mensajes de error se escriben en stderr y no con fmt.Println?

Los mensajes de error se envían a `stderr` porque este canal está destinado a reportar errores y permite separarlos de la salida normal del programa.

### c) ¿Qué pasaría en tu programa si escribieras valor, _ := strconv.ParseFloat(...) e ignoraras el error al ejecutar convertir abc KB MB?

Si ignorara el error, el programa no detectaría que "abc" no es un número válido. El resultado podría ser incorrecto y el usuario no recibiría una explicación adecuada del problema.

### d) ¿Qué ventaja tiene declarar los códigos de salida con const en lugar de escribir 1, 2 y 3 dentro de os.Exit?

Las constantes hacen que el código sea más fácil de entender y mantener. Es más claro usar `ErrValor` que escribir simplemente `2`, porque el nombre indica el significado del código.