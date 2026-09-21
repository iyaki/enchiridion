---
title: "Cómo empezar a programar con Solidity, el lenguaje preferido para crear los cada vez más demandados 'smart contracts'"
notion_id: 8785f416-f573-498b-8449-91daafc0ac81
notion_url: https://app.notion.com/p/C-mo-empezar-a-programar-con-Solidity-el-lenguaje-preferido-para-crear-los-cada-vez-m-s-demandados--8785f416f573498b844991daafc0ac81
last_edited: 2023-01-13T01:32:00.000Z
source_url: https://www.genbeta.com/a-fondo/como-empezar-a-programar-solidity-lenguaje-preferido-para-crear-cada-vez-demandados-smart-contracts
tags: ["Article", "Xataka | Genbeta", "Español", "Crypto"]
---
**Los 'smart contracts' o 'contratos inteligentes'** fueron un concepto a debate (y un sueño, en el mundo de los negocios) varios años antes de que existiera la tecnología necesaria para respaldarlos. De hecho, el abogado y criptógrafo **Nick Szabo ya los definió en 1998**, una década antes de que Satoshi Nakamoto creara el Bitcoin…

…y, con él, **la blockchain que ahora nos permite disponer de estos contratos autoejecutables**, automatizando así relaciones contractuales que no requieren de la intervención de un intermediario de confianza. **Pero, ¿cómo se crean estos contratos?**

## Programador, te presento a Solidity. Solidity, aquí el programador

Pues bien: [Solidity](https://solidity-es.readthedocs.io/es/latest/) es **el lenguaje más usado para escribir contratos inteligentes para la cadena de bloques Ethereum**. Es, de hecho, un lenguaje enfocado específicamente a esta tarea, desarrollado desde 2014 por diversos colaboradores del Proyecto Ethereum (su creador, Gavin Wood, es también co-creador de dicha criptodivisa).

**Solidity fue creado con el objetivo de ejecutarse en la Ethereum Virtual Machine (EVM**) que funciona sobre la blockchain de Ethereum. Sin embargo, la similitud entre esta cadena de bloques, y otras similares (como Polygon o Binance) permite implementar Solidity en otras redes y que su funcionamiento siga siendo predecible.

Es un lenguaje de alto nivel y orientado a objetos **cuya sintaxis se basa en ECMAScript** (al igual que JavaScript), con la principal diferencia de implementar **un tipado fuerte** a la hora de declarar el tipo de variables y argumentos. El objetivo de esto es **garantizar el rigor del contrato**: el compilador analizará nuestro código en tiempo de ejecución para verificar que intentamos realizar la operación adecuada con el tipo de valor adecuado.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Primer vistazo a Remix IDE

Para tu primera vez programando en Solidity, usaremos [Remix IDE](https://remix.ethereum.org/), un entorno de desarrollo basado en navegador desde el que podremos escribir, compilar e implementar 'smart contracts', además de —por ejemplo— almacenar archivos persistentes. **Hemos elegido esta herramienta porque es gratis, sencilla de usar, y no requiere descargas ni registros**.

Una vez entremos en la web, el aspecto de Remix IDE será algo bastante parecido a la siguiente imagen:

Deberemos **desplegar la carpeta 'contracts' dentro del explorador de archivos** del espacio de trabajo por defecto, y abrir uno de los ficheros .sol. Como están numerados, abriremos el primero, _1_Storage.sol_, que presentará este aspecto:

El código muestra dos elementos principales:

- 

En primer lugar, el **'pragma', donde especificamos qué versiones de Solidity podrán usarse para compilar** nuestro smart contract, un aspecto muy relevante a tenor de la rápida evolución del lenguaje.

- 

En segundo lugar, **el 'contract', equivalente al concepto de clase ('class')** presente en la mayoría de los lenguajes de programación modernos. Dentro de la clase nos encontramos con dos funciones predefinidas.

Sabiendo eso, ahora podemos crear **nuestro propio fichero desde cero**, y proceder a compilarlo y ejecutarlo.

## Hola, Mundo

Hacemos clic en el icono de 'New file' (Nuevo archivo), justo encima del árbol de directorios, y creamos **un fichero 'HolaMundo.sol'** en la carpeta 'contracts'.

Copiamos el siguiente código:

> 

// Mi primer contrato inteligente

pragma solidity >=0.5.0 <0.7.0;

contract HelloWorld {

function get()public pure returns (string memory){

```plain text
   return 'Hola, parte contratante de la primera parte';

```

Una vez hecho esto, hacemos clic en el tercer icono de la barra lateral, en el compilador de Solidity, tal como se muestra a continuación y, **asegurándonos de seleccionar una versión compatible del compilador, compilamos** sin tocar ningún otro aspecto de la configuración por defecto:

Una vez finalizado con éxito este paso, hacemos clic en el siguiente icono de la barra lateral. Pulsamos en el botón **'Deploy'**, y, una vez implementado el smart contract, aparecerá un botón con la leyenda 'Get', correspondiente a la única función de nuestro contrato inteligente. Haciendo clic, **veremos cómo nos devuelve la cadena de texto especificada en el código**:

## ¿Y ahora?

Ya has creado y ejecutado tu primer contrato inteligente, enhorabuena. Ahora que tienes unas nociones básicas sobre Solidity, **ha llegado el momento de empezar a formarte por tu cuenta** sobre este lenguaje. Te recomendamos algunos recursos online:

- Un vídeocurso de 3:20 h. en español, '**Curso Solidity desde Cero**', disponible en YouTube.
- 

Si te atreves con el inglés, [este otro vídeo](https://www.youtube.com/watch?v=ipwxYa-F1uY) del canal **freeCodeCamp.org** puede ser también de ayuda.

- 

Un curso gratuito, ['Uso de Solidity'](https://docs.microsoft.com/es-es/learn/modules/blockchain-learning-solidity/), disponible gratuitamente **en español en la plataforma Microsoft Learn**. O, si quieres empezar por el principio, puedes hacer toda la ruta de aprendizaje de ['Introducción al desarrollo de cadenas de bloques'](https://docs.microsoft.com/es-es/learn/paths/ethereum-blockchain-development/), que incluye el contenido anterior.

- 

**La documentación oficial y actualizada de Solidity**, en inglés y [en PDF](https://buildmedia.readthedocs.org/media/pdf/solidity/develop/solidity.pdf), creada por el proyecto Ethereum.
