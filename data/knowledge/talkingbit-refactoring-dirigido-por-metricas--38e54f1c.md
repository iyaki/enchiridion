---
title: "TalkingBit: Refactoring dirigido por métricas"
notion_id: 38e54f1c-7d23-81e0-a6a1-f242e64fe356
notion_url: https://app.notion.com/p/TalkingBit-Refactoring-dirigido-por-m-tricas-38e54f1c7d2381e0a6a1f242e64fe356
last_edited: 2026-06-29T03:39:00.000Z
source_url: https://franiglesias.github.io/metric-driven-refactoring/
tags: ["Article", "The Talking Bit - Fran Iglesias", "Español", "Software Architecture", "Quality", "Metrics", "Refactoring"]
---






## 





### 















```

```



```

```





### 





## 



```

```



```

```

### 



| Complejidad Ciclomática | Complejidad |
| --- | --- |
| 1-10 | Baja |
| 11-20 | Media |
| 21-50 | Alta |
| 50 o más | Muy alta |





### 





```

```



```

```



```

```



```

```



- 
- 



## 











1. 
2. 
3. 



```

```



```

```



```

```



```

```





## 







- 
- 



```

```

| Operador | Apariciones |
| --- | --- |
| `if` | 3 |
| `else` | 1 |
| `return` | 4 |
| `>` | 3 |
| `*` | 4 |
| `/` | 2 |
| `.` | 4 |
| `()` | 1 |
| `{}` | 4 |



| Operando | Apariciones |
| --- | --- |
| `param` | 7 |
| `this` | 4 |
| `aProperty` | 4 |
| `10` | 1 |
| `20` | 1 |
| `15` | 1 |
| `3` | 1 |
| `2` | 1 |
| `4` | 1 |



| Métrica | Fórmula | Resultado |
| --- | --- | --- |
| Vocabulario | n = 9 + 9 | **18** |
| Longitud | N = 26 + 21 | **47** |
| Volumen | V = 47 × log₂(18) | **196,3** |
| Dificultad | D = (9/2) × (21/9) | **10,5** |
| Esfuerzo | E = 10,5 × 196,3 | **2.061** |
| Tiempo | T = 2.061 / 18 | **114 seg** |
| Bugs estimados | B = 196,3 / 3000 | **0,065** |





- 
- 
- 





```

```



| Métrica | Valor |
| --- | --- |
| n1 | 9 |
| n2 | 9 |
| N1 | 26 |
| N2 | 34 |
| Vocabulary (n) | 18 |
| Length (N) | 60 |
| Volume (V) | ~250 |



```

```





| Operador | Apariciones |
| --- | --- |
| `const` | 1 |
| `=` | 1 |
| `{}` | 2 |
| `.` | 2 |
| `()` | 2 |
| `=>` | 1 |
| `>` | 1 |
| `??` | 1 |
| `return` | 1 |
| `*` | 2 |





| Operando | Apariciones |
| --- | --- |
| `factor` | 3 |
| `this` | 2 |
| `MULTIPLIERS` | 1 |
| `find` | 1 |
| `threshold` | 2 |
| `param` | 2 |
| `4` | 1 |
| `aProperty` | 1 |





| Métrica | Fórmula | Resultado |
| --- | --- | --- |
| Vocabulario | n = 10 + 8 | **18** |
| Longitud | N = 14 + 13 | **27** |
| Volumen | V = 27 × log₂(18) | **112,8** |
| Dificultad | D = (10/2) × (13/8) | **8,1** |
| Esfuerzo | E = 8,1 × 112,8 | **913,7** |
| Tiempo | T = 913,7 / 18 | **50,8 seg** |
| Bugs estimados | B = 112,8 / 3000 | **0,038** |



| Métrica | Original | Refactor (tabla) | Δ |
| --- | --- | --- | --- |
| Volumen (V) | 196,3 | 112,8 | −42% |
| Dificultad (D) | 10,5 | 8,1 | −23% |
| Esfuerzo (E) | 2.061 | 913,7 | −56% |
| Tiempo (seg) | 114 | 50,8 | −56% |
| Bugs estimados | 0,065 | 0,038 | −42% |

## 



- 
- 
- 









## 





- 
- 
- 
- 
- 
- 



```

```

```

```







```

```







### 















- 
- 
- 
- 
- 
- 





### 









- 
- 





### 







- 
- 









|  | db | emailClient | paymentGateway | inventoryService | logger | cache |
| --- | --- | --- | --- | --- | --- | --- |
| processOrder | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| cancelOrder | ✓ | ✓ | — | — | ✓ | ✓ |
| refundOrder | ✓ | ✓ | ✓ | — | ✓ | — |
| methods per attr | 3 | 3 | 2 | 1 | 3 | 2 |



- 
- 
- 



- 
- 





### 





```

```









- 
- 
- 



```

```



## 











- 
- 
- 









- 
- 
- 

## 





| Métrica | Qué mide | Nivel | Señal de refactor | Umbral orientativo |
| --- | --- | --- | --- | --- |
| **Complejidad Ciclomática** | Caminos de ejecución independientes | Función | Extraer métodos, simplificar lógica | > 10 |
| **Complejidad Cognitiva** | Esfuerzo mental de lectura | Función | Reducir anidamiento, aplanar estructura | > 15 |
| **Volumen (Halstead)** | Cantidad de información del programa | Función | Simplificar expresiones, renombrar | > 1000 |
| **Dificultad (Halstead)** | Propensión a errores | Función | Reducir variedad de operadores y operandos | > 10 |
| **Esfuerzo (Halstead)** | Carga mental de implementación | Función | Comparativo entre versiones | — |
| **Bugs estimados (Halstead)** | Errores esperados estadísticamente | Función | Comparativo entre versiones | — |
| **LOC** | Tamaño físico | Función/Clase | Extraer métodos o clases | > 200 (función) |
| **WMC** | Lógica total de la clase | Clase | Dividir clase, mover responsabilidades | > 50 |
| **DIT** | Profundidad de herencia | Jerarquía | Aplanar herencia, preferir composición | > 5 |
| **NOC** | Amplitud de herencia | Jerarquía | Revisar abstracción de clase base | > 10 |
| **CBO** | Acoplamiento con otras clases | Clase | Reducir dependencias, introducir interfaces | > 14 |
| **RFC** | Alcance de ejecución ante un mensaje | Clase | Reducir responsabilidades, extraer servicios | > 50 |
| **LCOM** | Cohesión interna de la clase | Clase | Dividir clase | > 5 |
| **LCOM-HS** | Cohesión interna de la clase | Clase | Dividir clase | tiende a 1 |
| **MI** | Mantenibilidad agregada | Módulo | Priorizar qué refactorizar primero | < 20 (VS) |

## 



| Herramienta | TS | PHP | Java | Go | C# | Nivel |
| --- | --- | --- | --- | --- | --- | --- |
| SonarQube | ✓ | ✓ | ✓ | ✓ | ✓ | Enterprise |
| ESLint | ✓ | – | – | – | – | Ligero |
| PhpMetrics | – | ✓ | – | – | – | OO PHP |
| CK | – | – | ✓ | – | – | Académico |
| PMD | – | – | ✓ | – | – | Java clásico |
| golangci-lint | – | – | – | ✓ | – | Go estándar |
| NDepend | – | – | – | – | ✓ | Muy profundo |
| VS Metrics | – | – | – | – | ✓ | Integrado |



### 







- 
- 
- 
- 
- 
- 
- 
- 

## 









1. 
2. 
3. 
4. 
5. 
6. 
7. 
8. 
9. 
