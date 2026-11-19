# 04\-taller1

## Descripción

Este programa implementa el ejercicio clásico FizzBuzz utilizando un switch sin condición. Recorre los números del 1 al 30 e imprime:

- Fizz para los múltiplos de 3.
- Buzz para los múltiplos de 5.
- FizzBuzz para los múltiplos de 15.
- El número cuando no cumple ninguna condición.

Además cuenta cuántas veces aparece cada palabra y muestra los resultados al final.

## Cómo ejecutar

Ejecutar:

go run .

## Resultado

Salida obtenida:

1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz

Fizz=8 Buzz=4 FizzBuzz=2

### Experimento del orden

Al mover el caso:

case i%3 == 0

antes de:

case i%15 == 0

la salida cambia porque los múltiplos de 15 también son múltiplos de 3. Como Go ejecuta únicamente el primer caso verdadero, nunca llega a evaluar el caso de FizzBuzz.

## Preguntas

### a) ¿Qué salida y qué conteos obtuviste al poner el case de 3 antes que el de 15, y por qué FizzBuzz dejó de aparecer?

Cuando coloqué el caso de múltiplos de 3 antes que el de múltiplos de 15, los números 15 y 30 fueron tomados como Fizz. Como resultado, FizzBuzz dejó de aparecer. Esto ocurrió porque Go ejecuta el primer case verdadero que encuentra.

### b) Cuando i vale 15, ¿cuántos case son verdaderos y cuál ejecuta Go? ¿Por qué no hace falta escribir break al final de cada case?

Cuando i vale 15 son verdaderos los casos de múltiplos de 15, de 3 y de 5. Sin embargo, Go ejecuta únicamente el primer case verdadero que encuentra. No hace falta usar break porque Go sale automáticamente del switch después de ejecutar un case.

