25/11/2026 - Tarea 5
Actividad realizada
Desarrollé un programa en Go para analizar textos y obtener la cantidad de bytes, runas, palabras, letras, dígitos, mayúsculas y runas que ocupan más de un byte. También implementé la validación del límite de 42 runas para subtítulos accesibles y el manejo de códigos de salida para errores.

Problema encontrado
Al intentar crear el módulo de Go y ejecutar el programa, PowerShell mostraba el error:

"go: El término 'go' no se reconoce como nombre de un cmdlet"

Esto impedía utilizar comandos como go mod init, go run y go build.

Solución aplicada
Instalé Go correctamente en Windows y reinicié la terminal. Después verifiqué la instalación con:

go version