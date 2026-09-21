---
title: "Delegación de dominios de nic.ar"
notion_id: 94762364-c4e7-4985-a7f1-ff2d96917834
notion_url: https://app.notion.com/p/Delegaci-n-de-dominios-de-nic-ar-94762364c4e74985a7f1ff2d96917834
last_edited: 2026-09-21T17:20:00.000Z
source_url: https://gist.github.com/franleplant/6a419687c402f898c883
tags: ["Español", "Web Development", "SysAdmin", "Network", "Article"]
---
[https://gist.github.com/franleplant/6a419687c402f898c883](https://gist.github.com/franleplant/6a419687c402f898c883)

IMPORTANTE mi recomendacion es evitar nic.ar en lo posible ya que es tenido diversos problemas con su servicio, la renovacion es incomoda, me han disputado y quitado dominios que he tenido y pagado por años, tenes que sacar clave fiscal para crear un dominio. Si no tienen problemas teniendo un dominio que no sea .com.ar entonces usen Godaddy u otra cosa.

Dicho esto, si no te es posible escaparle al .com.ar segui leyendo.

Si te sirvio esta informacion considera aportar para apoyar mi trabajo via Bitcoin 1MBakrHeKFxks4SbhY7MdnhDEQnkDnPoJK Desde ya muchas gracias

## Como delegar?

Mi recomendacion es usar un DNS hosted como Route 53 de AWS. DNS hosted significa que es un servidor de DNS, conectado a la base de datos distribuida que es DNS, hostead / administrado por otro, en este caso AWS. Hay otras alternativas solo que yo solo use Route 53. Busquen Hosted DNS en internet para ver alternativas. Route 53 es realmente muy barato, algo asi como 1 dolar por mes mas 0.2 dolares por hosted zone (dominio)

La delegacion es el proceso por el cual vos asignas un registro NS a tu dominio en nic.ar, esto quiere decir, indicar a la base de datos global DNS que si quiere saber comopara acceder a tudominio.com.ar tiene que pedirle el ip a NS.

Como DNS es una base de datos distribuida gerarquica este cambio tipicamente toma tiempo. En el caso de la entrada que vamos a crear en Route 53 (CNAME, A, etc), como es de AWS, la actualizacion va a ser rapida, unas cuantas horas cuanto mucho. En el caso del cambio de nic.ar, yo he tenido que esperar 3 o 4 dias, pero lo esperable seria 24hs (otra razon mas para evitar este servicio de #@#%#!).

Entonces

- en Route 53, clickear en Created hosted zone (hosted zone = dns para un dominio)
- ingresar nombre de dominiop tudominio.com.ar y dejar todas las opciones por defecto
- en este punto vamos a ver los registros que AWS crean por defecto para nuestro dominio: NS y SOA
- el registro NS (name server) es el que vamos a agregar en la parte de delegacion de nic.ar, recomiendo agregar uno por uno todos los NS que nos da amazon (se usan para contingencias y balancear cargas). Una vez que pasa esto esperar.
- En este punto tenemos que cuando alguien accede a tudominio.com.ar se va a generar una consulta DNS a nic.ar y este va a redireccionar a Amazon, pero todavia Amazon no sabe a donde redireccionar, para eso, el siguiente paso
- Para propiamenete resolver la consulta DNS tudominio.com.ar, tenemos dos alternativas, que se resuelva a un IP de donde esta hosteada tu pagina (i.e. tu propio servidor), para eso agregamos un registro A con el IP del servidor; o si queremos redirigir a otro dominio (i.e. un servidor de amazon o las paginas de github) tenemos que agregar una entrada CNAME con el dominio destion (i.e. aws-123.com).

Para mas informacion de como redirigir a Github leer la documentacion de github.

## Checkear el estado de la delegacion

Delegar DNS con nic.ar apesta y suele ser bastante lento. A continuacion detallo una buena forma para averiguar como viene el tramite:

(para alternativas, ver comentarios)

```plain text
dig +trace tudominio.com.ar
```

Al ejecutar esto vas a ver como se va intentando resolver el dominio.

Asi se resuelve un dominio no delegado:

```plain text
$ dig +trace asdasdjhasdkjaskj123o123.com.ar 23:50:47 ; <<>> DiG 9.8.3-P1 <<>> +trace asdasdjhasdkjaskj123o123.com.ar ;; global options: +cmd . 12266 IN NS m.root-servers.net. . 12266 IN NS k.root-servers.net. . 12266 IN NS j.root-servers.net. . 12266 IN NS c.root-servers.net. . 12266 IN NS i.root-servers.net. . 12266 IN NS e.root-servers.net. . 12266 IN NS b.root-servers.net. . 12266 IN NS g.root-servers.net. . 12266 IN NS h.root-servers.net. . 12266 IN NS d.root-servers.net. . 12266 IN NS f.root-servers.net. . 12266 IN NS l.root-servers.net. . 12266 IN NS a.root-servers.net. ;; Received 228 bytes from 8.8.8.8#53(8.8.8.8) in 66 ms ar. 172800 IN NS b.dns.ar. ar. 172800 IN NS c.dns.ar. ar. 172800 IN NS a.dns.ar. ar. 172800 IN NS ns2.switch.ch. ar. 172800 IN NS ar.cctld.authdns.ripe.net. ar. 172800 IN NS d.dns.ar. ;; Received 419 bytes from 192.36.148.17#53(192.36.148.17) in 401 ms com.ar. 3600 IN SOA c.dns.ar. noc.nic.gob.ar. 2015110709 43200 3600 1728000 86400 ;; Received 103 bytes from 200.16.99.17#53(200.16.99.17) in 102 ms
```

La parte notable es que el ultimo escalon es el SOA y eso apunta a nic.ar, lo cual da la pauta de que nic.ar todavia no propago la delegacion.
