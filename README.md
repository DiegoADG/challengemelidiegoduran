# challengeMELI

Procesos, hilos y corrutinas
● Un caso en el que usarías procesos para resolver un problema y por qué.
los procesos son las ejecuciones de los programas que nosotros realizamos para desarrollo es un proceso y por lo cual si hablamos que problema podemos resolver con un proceso seria la ejecucion de programas.
● Un caso en el que usarías threads para resolver un problema y por qué.
los hilos son las ejecuciones que si se dan paralelamente, estos se utilizan para ejecuctar tareas que se pueden realizar de forma paralela y asi ahorrar tiempo de procesamiento, los hilos usan la memoria de un solo procesamiento y por lo cual tambien suelen ser pesadas.

● Un caso en el que usarías corrutinas para resolver un problema y por qué.
las corrutinas son las ejecuciones donde se aprovecha al maximo la cpu, ya que permite hacer ejecuciones para varias tareas, pero liberando el hilo para que se ejucuten todas y asi parecer que se estuvieran ejecuctando paralelamente, pero la ejecucion se realiza una tarea y luego la para y ejecuta otra y asi sucesivamente. estas corrutinas son muy optimas para los desarrollos que puede tener varios procesos que se pueden ejeucutar de forma "paralela" para asi poder ganar mucho mas tiempo en la respuesta de los desarrollos

Optimización de recursos del sistema operativo
Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de
ellos información en una API HTTP. ¿Cómo lo harías? Explicar.
Para poder realizar una consulta de semejante magnitud yo utilizaria corrutinas, ya que son mas livianas y permite generar una mayor cantidad de rutinas ya si poder realizar la gran cantidad de tareas "paralelas" y asi obtener una disminucion de tiempos en la respuesta del api, muy diferente a los hilos, que nos podria dejar sin memoria de procesamiento.

Análisis de complejidad
● Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con
complejidades O(n2), O(n3), O(2n) y O(n log n), respectivamente, ¿Cuál de los
algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.

Para estos algoritmos lo que favoreceria serian O(n log n), ya que es un algoritmo mas eficiente porque crece a un ritmo mucho mas lento a medida que aumenta el tamaño de entrada, pero aun asi su crecimiento es muy rapido considerando otro tipo de complejidades como la O(n) o O(log n)
las otras complejidades O(n2), O(n3), O(2n) las descartaria, esto por que tienen un crecimiento tan rapido que no son para nada eficientes y la recomendacion es no usarlas

● Asume que dispones de dos bases de datos para utilizar en diferentes
problemas a resolver. La primera llamada AlfaDB tiene una complejidad de O(1)
en consulta y O(n2) en escritura. La segunda llamada BetaDB que tiene una
complejidad de O(log n) tanto para consulta, como para escritura.¿Describe en
forma sucinta, qué casos de uso podrías atacar con cada una?

 Para la ambas bases de datos las podemos utilizar haciendo uso del patrón cqrs, utilizando la primera bd para lectura y la segunda para escritura y se puede utilizar algún mecanismo asincrónico para sincronizarla cercano al tiempo real y asi poder aprovechar los beneficios de cada db. 