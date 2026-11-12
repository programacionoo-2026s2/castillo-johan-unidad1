Tarea 1 - Ámbito y Shadowing en Go
Código analizado

package main

import "fmt"

func suma() int {
	total := 0

	for i := 1; i <= 3; i++ {
		total := total + i
		_ = total
	}

	return total
}

func main() {
	fmt.Println("total:", suma())
}
1. ¿Qué imprime?
Al ejecutar el programa se obtiene:

total: 0
2. ¿Por qué ocurre?
El resultado es 0 debido al ámbito (scope) de las variables.

La variable total se declara inicialmente fuera del ciclo:

total := 0
Sin embargo, dentro del for se utiliza nuevamente:

total := total + i
El operador := crea una nueva variable local dentro del ámbito del ciclo en lugar de modificar la variable total declarada fuera.

Esto significa que en cada iteración se crea una nueva variable llamada total, la cual solo existe dentro del bloque del for.

Por esta razón, la variable total externa nunca cambia su valor inicial de 0.

Al llegar a:

return total
se devuelve la variable externa, que sigue valiendo 0.

Este fenómeno se conoce como shadowing o ocultamiento de variables, porque una variable interna oculta temporalmente otra variable con el mismo nombre declarada en un ámbito superior.

3. Corríjalo cambiando un solo carácter
La corrección consiste en cambiar:

total := total + i
por:

total = total + i
Únicamente se elimina el carácter :.

Código corregido:

func suma() int {
	total := 0

	for i := 1; i <= 3; i++ {
		total = total + i
	}

	return total
}
Resultado:

total: 6
Porque:

1 + 2 + 3 = 6

Extra: go vet ./...
Al ejecutar:

go vet ./...
normalmente no detecta este problema.

go vet está diseñado para encontrar errores comunes y construcciones sospechosas, pero no reporta todos los casos de shadowing de variables.

Por esta razón es necesario utilizar herramientas adicionales.

Investigación: Analizador Shadow
¿Qué es?
Shadow es un analizador estático para Go que detecta variables declaradas en ámbitos internos que ocultan variables con el mismo nombre declaradas en ámbitos externos.

Su objetivo es evitar errores difíciles de identificar producidos por el shadowing.

Instalación
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
Uso
go vet -vettool=$(which shadow) ./...
En Windows:

go vet -vettool=$env:GOPATH\bin\shadow.exe ./...
¿Qué detecta en este programa?
Detecta la línea:

total := total + i
porque declara una nueva variable total que oculta la variable total definida fuera del ciclo.

El analizador genera una advertencia similar a:

declaration of "total" shadows declaration at line X
Ventajas del analizador Shadow
Detecta errores de ámbito difíciles de observar.
Ayuda a evitar resultados inesperados.
Complementa las verificaciones de go vet.
Mejora la calidad y mantenibilidad del código.
Conclusión
El programa imprime total: 0 debido al ámbito de las variables y al fenómeno de shadowing. La variable total interna oculta a la variable externa, impidiendo que esta última sea modificada. La corrección consiste en reemplazar := por =. Aunque go vet normalmente no detecta el problema, el analizador shadow sí permite identificar este tipo de errores.