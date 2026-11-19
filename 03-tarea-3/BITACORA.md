## Tarea 3

Durante el desarrollo de la tarea tuve un problema al validar los argumentos ingresados por el usuario. Inicialmente el programa se detenía con errores cuando faltaban parámetros. Lo solucioné verificando que len(os.Args) sea igual a 4 antes de procesar los datos.

También tuve dificultades al convertir el valor ingresado por el usuario a tipo float64. El problema se resolvió utilizando strconv.ParseFloat y validando el error devuelto por la función.

