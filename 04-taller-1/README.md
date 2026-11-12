Taller 1 - Análisis de texto en Go
Objetivo
Analizar una cadena de texto utilizando las librerías estándar de Go para contar bytes, runas, palabras, mayúsculas, dígitos y caracteres con tilde o ñ.

Código utilizado
texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."
Actividades realizadas
1. Conteo de bytes y runas
Se utilizó:

len(texto)
para contar bytes y:

utf8.RuneCountInString(texto)
para contar runas.

Resultado
Bytes: 62
Runas: 58
2. Conteo de palabras
Se utilizó:

strings.Fields(texto)
para separar las palabras usando espacios.

Resultado:

Palabras: 9
3. Conteo de mayúsculas y dígitos
Se recorrió la cadena mediante:

for _, r := range texto
y se emplearon las funciones:

unicode.IsUpper(r)
unicode.IsDigit(r)
Resultado:

Mayúsculas: 3
Dígitos: 4
4. Conteo de letras con tilde o ñ
Se utilizó:

strings.ContainsRune("áéíóúÁÉÍÓÚñÑ", r)
para identificar caracteres acentuados o la letra ñ.

Resultado:

Con tilde/ñ: 3
¿Por qué bytes y runas no coinciden?
Los bytes representan la cantidad de espacio ocupada en memoria, mientras que las runas representan caracteres Unicode.

En UTF-8 algunas letras especiales ocupan más de un byte.

Ejemplos presentes en el texto:

ó
á
¡
ñ
Cada uno de estos caracteres utiliza más de un byte para almacenarse.

Por eso:

Bytes: 62
Runas: 58
Hay más bytes que runas.

Conclusión
Go permite trabajar correctamente con texto Unicode mediante runas. El uso de paquetes como unicode, utf8 y strings facilita el análisis de cadenas que contienen caracteres especiales del idioma español.