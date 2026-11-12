Tarea 2 - Permisos con operadores de bits en Go
Objetivo
Representar permisos de usuario utilizando operaciones bit a bit similares al sistema de permisos de Linux (chmod).

Estrategia de desarrollo
Para resolver el problema se utilizó un tipo personalizado llamado Permiso.

type Permiso uint8
Se definieron constantes para representar lectura, escritura y ejecución.

const (
    Leer     Permiso = 4
    Escribir Permiso = 2
    Ejecutar Permiso = 1
)
La numeración corresponde a la utilizada por el comando chmod en sistemas Linux:

Permiso	Valor
r	4
w	2
x	1
Función describir
La función recibe un conjunto de permisos y utiliza el operador & para verificar si cada bit se encuentra activo.

Ejemplo:

if p&Leer != 0
Si el permiso existe se imprime la letra correspondiente; caso contrario se imprime un guion.

Resultado:

rwx
r-x
r--
Permisos iniciales
Usuario:

rwx
Grupo:

r-x
Otros:

r--
Representación conjunta:

rwxr-xr--
Valores numéricos
Entidad	Valor
Usuario	7
Grupo	5
Otros	4
Representación chmod:

754
Modificaciones de permisos
Agregar escritura a otros
Se utilizó:

otros |= Escribir
Resultado:

rw-
Quitar ejecución al grupo
Se utilizó:

grupo &^= Ejecutar
Resultado:

r--
Resultado final
rwxr--rw-
Conclusión
El uso de operadores bit a bit permite manipular permisos de forma eficiente y es la misma idea utilizada por sistemas tipo Unix para representar permisos mediante números como 754, 755 o 777.