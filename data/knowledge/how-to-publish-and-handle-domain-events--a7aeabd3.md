---
title: "How to publish and handle Domain Events"
notion_id: a7aeabd3-7cc6-464a-886c-3cb286085bda
notion_url: https://app.notion.com/p/How-to-publish-and-handle-Domain-Events-a7aeabd37cc6464a886c3cb286085bda
last_edited: 2022-12-21T15:29:00.000Z
source_url: http://www.kamilgrzybek.com/design/how-to-publish-and-handle-domain-events/
tags: ["System Design / Software Architecture", "Article", "Kamil Grzybek", "English"]
---


## 





## 





### 



| 12345678910111213141516 | public class ShopApplicationService{private readonly IOrderRepository orderRepository;private readonly IShopRepository shopRepository;private readonly IEventsPublisher eventsPublisher;public void PlaceOrder(int shopId, int orderId){Shop shop = this.shopRepository.GetById(shopId);Order order = this.orderRepository.GetById(orderId);List<IDomainEvent> events = shop.PlaceOrder(order);eventsPublisher.Publish(events);}} |
| --- | --- |



## 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





### 

















<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



## 


