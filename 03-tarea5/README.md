Tarea 5 - Analizador de texto
Descripción
Esta aplicación de consola desarrollada en Go analiza un texto recibido como argumento por línea de comandos.

El programa calcula:

Cantidad de bytes.
Cantidad de runas.
Cantidad de palabras.
Cantidad de letras.
Cantidad de dígitos.
Cantidad de letras mayúsculas.
Cantidad de runas que ocupan más de un byte.
Además verifica si el texto supera el límite de 42 runas definido para subtítulos accesibles.

Cómo ejecutar
Compilar el programa:

go build -o analizar .
Ejecutar el programa:

.\analizar "Hola mundo"
También puede ejecutarse con:

go run . "Hola mundo"
Es importante colocar el texto entre comillas cuando contiene espacios.

Códigos de salida
Código	Significado
0	Ejecución correcta
1	Error de uso
2	Texto vacío
3	El texto supera el límite de 42 runas
Pruebas realizadas
Comando	Resultado
analizar "Hola mundo"	Correcto
analizar "Canción de cuna 2026"	Correcto
analizar "ÁÉÍ ñandú"	Correcto
analizar " espacios múltiples "	Correcto
analizar " "	Finaliza con código 2
analizar "Este subtítulo es demasiado largo para leerse bien"	Finaliza con código 3
analizar	Finaliza con código 1
analizar hola mundo	Finaliza con código 1
Preguntas
a) En "Canción de cuna 2026" el programa muestra 21 bytes y 20 runas. ¿Por qué no coinciden y cuál usarías para el límite de subtítulos?
Los bytes y las runas no coinciden porque algunos caracteres Unicode ocupan más de un byte cuando se almacenan en UTF-8. En este caso, la letra "ó" utiliza dos bytes. Por esa razón el texto tiene 21 bytes pero solamente 20 runas. Para controlar el límite de subtítulos utilizaría las runas, porque representan los caracteres que realmente ve y lee la persona usuaria.

b) En for i, r := range texto, ¿qué valen i y r en cada vuelta? Si texto es ñu, ¿qué valores toma i?
La variable i representa la posición en bytes dentro de la cadena de texto y la variable r representa la runa actual que se está recorriendo.

Si el texto es:

ñu
los valores de i serán:

0
2
Esto ocurre porque la letra ñ ocupa dos bytes en UTF-8, por lo que la siguiente runa comienza en la posición 2.

c) ¿Por qué las palabras se cuentan con la variable enPalabra y no contando espacios? Usa el caso 4 como ejemplo.
Las palabras se cuentan utilizando la variable enPalabra porque puede haber varios espacios consecutivos, espacios al inicio o espacios al final del texto. Si solamente se contaran espacios, el resultado podría ser incorrecto.

Por ejemplo, en el caso:

 espacios   múltiples
existen espacios al inicio y varios espacios entre las palabras, pero únicamente hay dos palabras. La variable enPalabra permite detectar exactamente cuándo comienza una palabra nueva y evitar contar espacios adicionales como si fueran palabras.