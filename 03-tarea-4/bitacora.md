## 18/11/2026 \- Tarea 4

Durante el desarrollo de la tarea tuve un error porque el archivo go.mod no existía en el proyecto. Lo solucioné ejecutando el comando go mod init tarea4.

También tuve un error de importación porque mezclé código del archivo main.go dentro de almacenamiento.go, generando un ciclo de importación. Lo corregí dejando la lógica principal únicamente en main.go y las funciones de conversión dentro del paquete almacenamiento.

