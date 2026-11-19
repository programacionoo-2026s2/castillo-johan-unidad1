
| 03-tarea4 | Conversor de unidades de almacenamiento utilizando paquetes internos en Go |
# 03\-tarea4

## Descripción y estructura

Esta aplicación permite convertir valores entre unidades de almacenamiento utilizando Go. El proyecto fue reorganizado para separar la lógica de conversión en un paquete propio llamado almacenamiento, ubicado dentro de la carpeta internal.

Estructura:

03\-tarea4/

├── go.mod

├── main.go

├── README.md

└── internal/

    └── almacenamiento/

        └── almacenamiento.go

- main.go: recibe los argumentos, valida los datos y muestra los resultados.
- almacenamiento.go: contiene las constantes, funciones de conversión y documentación del paquete.
- go.mod: define el módulo tarea4.

## Cómo ejecutar

Ejemplo:

go run . 2048 KB MB

Salida:

2048.00 KB = 2.00 MB

Códigos de salida:

| Código | Significado |
| --- | --- |
| 0 | Ejecución correcta |
| 1 | Cantidad incorrecta de argumentos |
| 2 | Valor numérico inválido o negativo |
| 3 | Unidad no válida |

## Pruebas realizadas

| Comando | Resultado |
| --- | --- |
| go run . 2048 KB MB | 2048.00 KB = 2.00 MB |
| go run . 1 gb mb | 1.00 GB = 1024.00 MB |
| go run . 1500 B KB | 1500.00 B = 1.46 KB |
| go run . 0.5 MB KB | 0.50 MB = 512.00 KB |
| go run . abc KB MB | Valor inválido (código 2) |
| go run . \-5 KB MB | Valor inválido (código 2) |
| go run . 5 TB MB | Unidad inválida (código 3) |
| go run . 5 KB | Error de uso (código 1) |

## Experimento del compilador

Se modificó temporalmente el archivo main.go para intentar llamar a la función:

resultado := almacenamiento.redondear(valor)

Al ejecutar:

go build ./...

el compilador mostró el siguiente error:

\# tarea4

.\\main.go:33:2: declared and not used: bytesOrigen

.\\main.go:39:2: declared and not used: bytesDestino

.\\main.go:44:29: name redondear not exported by package almacenamiento

La línea más importante es:

name redondear not exported by package almacenamiento

Este error ocurre porque redondear comienza con minúscula. En Go, los identificadores que empiezan con minúscula son privados y solo pueden utilizarse dentro de su propio paquete. Por esta razón, main.go no puede acceder directamente a redondear y debe utilizar la función pública Convertir.

## Preguntas

### a) ¿Por qué Bytes y Convertir empiezan con mayúscula y redondear con minúscula? ¿Qué ganas al dejar redondear oculta?

Las funciones Bytes y Convertir empiezan con mayúscula porque son funciones exportadas. Esto permite que otros paquetes del proyecto puedan utilizarlas mediante la importación del paquete almacenamiento.

La función redondear comienza con minúscula porque es una función no exportada. Solo puede utilizarse dentro del paquete almacenamiento. Al mantenerla oculta se evita que otros archivos dependan de detalles internos de implementación y se mantiene una mejor organización del código.

### b) ¿Qué mensaje dio el compilador al usar almacenamiento.redondear desde main.go y qué regla de Go lo explica?

Al ejecutar:

go build ./...

el compilador mostró:

.\\main.go:44:29: name redondear not exported by package almacenamiento

Este mensaje aparece porque Go solo permite acceder desde otros paquetes a identificadores que comienzan con letra mayúscula. Como redondear comienza con minúscula, es una función privada del paquete almacenamiento y no puede utilizarse desde main.go.

### c) ¿Para qué sirve que el paquete esté dentro de internal/ y qué pasaría si otro módulo intentara importarlo?

La carpeta internal sirve para restringir el acceso a los paquetes que contiene. De esta forma, el paquete almacenamiento solo puede ser utilizado por código perteneciente al mismo módulo.

Si otro proyecto intentara importar:

import "tarea4/internal/almacenamiento"

el compilador de Go impediría la compilación mostrando un error, ya que los paquetes ubicados dentro de internal no pueden ser utilizados desde módulos externos.

