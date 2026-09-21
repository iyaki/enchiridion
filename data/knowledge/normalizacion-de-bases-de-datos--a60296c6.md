---
title: "Normalización de bases de datos"
notion_id: a60296c6-94f1-4c6d-832e-961ab21e3684
notion_url: https://app.notion.com/p/Normalizaci-n-de-bases-de-datos-a60296c694f14c6d832e961ab21e3684
last_edited: 2022-12-31T22:08:00.000Z
source_url: https://franiglesias.github.io/db-normalization/
tags: ["Article", "The Talking Bit - Fran Iglesias", "Español", "System Design / Software Architecture", "Databases"]
---














## 



1. 
2. 
3. 
4. 
5. 











| id | name | surname | email1 | email2 | email3 |
| --- | --- | --- | --- | --- | --- |
| 1 | Pepa | López | pepa.lopez@example.com | `null` | `null` |
| 2 | Jaime | Rodríguez | j.r@example.com | jaime2123@example.com | `null` |









| id | name | surname | email |
| --- | --- | --- | --- |
| 1 | Pepa | López | pepa.lopez@example.com |
| 2 | Jaime | Rodríguez | j.r@example.com, jaime2123@example.com |







| id | name | surname |
| --- | --- | --- |
| 1 | Pepa | López |
| 2 | Jaime | Rodríguez |



| person_id | email |
| --- | --- |
| 1 | pepa.lopez@example.com |
| 2 | jaime2123@example.com |
| 2 | jaime2123@example.com |



## 















| id | name | surname |
| --- | --- | --- |
| 1 | Pepa | López |
| 2 | Jaime | Rodríguez |







| id | name | surname |
| --- | --- | --- |
| 1 | Pepa | López |
| 2 | Jaime | Rodríguez |
| 3 | Pepa | López |
| 4 | Jaime | Martínez |







| section | product | name | store |
| --- | --- | --- | --- |
| fruits | 001 | oranges | main st |
| fruits | 002 | apples | main st |
| dairy | 001 | greek yoghourt | river st |
| bakery | 001 | bread | river st |
| bakery | 002 | donut | river st |











| section | product | name |
| --- | --- | --- |
| fruits | 001 | oranges |
| fruits | 002 | apples |
| dairy | 001 | greek yoghourt |
| bakery | 001 | bread |
| bakery | 002 | donut |



| section | store |
| --- | --- |
| fruits | main st |
| dairy | river st |
| bakery | river st |



## 









| id | name | team_id | team_name |
| --- | --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 | Accounts |
| 2 | Michael Caine | 2 | Technology |
| 3 | Mary Shelley | 2 | Technology |
| 4 | Jane Austen | 1 | Sales |







| id | name | team_id | team_name |
| --- | --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 | Accounts |
| 2 | Michael Caine | **2** | **Technology** |
| 3 | Mary Shelley | **2** | **Systems** |
| 4 | Jane Austen | 1 | Sales |





| id | name | team_id |
| --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 |
| 2 | Michael Caine | 2 |
| 3 | Mary Shelley | 2 |
| 4 | Jane Austen | 1 |



| id | team_name |
| --- | --- |
| 1 | Sales |
| 2 | Techonology |
| 3 | Accounts |

## 









| id | name | team_id | supervisor |
| --- | --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 | Juana López |
| 2 | Michael Caine | 2 | Inma González |
| 3 | Mary Shelley | 2 | Javier Pons |
| 4 | Jane Austen | 1 | Ángela Martínez |





| id | team_name |
| --- | --- |
| 1 | Sales |
| 2 | Techonology |
| 3 | Accounts |



| id | name | team_id |
| --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 |
| 2 | Michael Caine | 2 |
| 3 | Mary Shelley | 2 |
| 4 | Jane Austen | 1 |



| employee_id | supervisor |
| --- | --- |
| 1 | Juana López |
| 2 | Inma González |
| 3 | Javier Pons |
| 4 | Ángela Martínez |

## 







| id | name | team_id |
| --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 |
| 2 | Michael Caine | 2 |
| 3 | Mary Shelley | 2 |
| 4 | Jane Austen | 1 |



| id | team_name |
| --- | --- |
| 1 | Sales |
| 2 | Techonology |
| 3 | Accounts |







| id | name | team_id |
| --- | --- | --- |
| 1 | Ebenizer Scrooge | 3 |
| 2 | Michael Caine | 2 |
| 3 | Mary Shelley | 2 |
| 4 | Jane Austen | 1 |
| 3 | Mary Shelley | 3 |
| 4 | Jane Austen | 2 |





| id | name |
| --- | --- |
| 1 | Ebenizer Scrooge |
| 2 | Michael Caine |
| 3 | Mary Shelley |
| 4 | Jane Austen |



| id | team_name |
| --- | --- |
| 1 | Sales |
| 2 | Technology |
| 3 | Accounts |



| employee_id | team_id |
| --- | --- |
| 1 | 3 |
| 2 | 2 |
| 3 | 2 |
| 4 | 1 |
| **3** | **1** |

## 



- 
- 









| workshop | vendor | service | assurance |
| --- | --- | --- | --- |
| López | Renault | Mechanics | Generalli |
| López | Seat | Mechanics | Generalli |
| López | Volvo | Mechanics | Generalli |
| López | Volvo | Mechanics | Axa |
| López | Seat | Mechanics | Axa |
| Tuercas | Renault | Painting | Generalli |
| Tuercas | Volvo | Painting | Mutua M |
| Tuercas | Volkswagen | Electronic | Mutua M |
| Tuercas | Renault | Electronic | Axa |







| workshop | vendor |
| --- | --- |
| López | Renault |
| López | Volvo |
| López | Seat |
| Tuercas | Volvo |
| Tuercas | Volkswagen |
| Tuercas | Renault |



| workshop | service |
| --- | --- |
| López | Mechanics |
| Tuercas | Painting |
| Tuercas | Electronic |



| workshop | assurance |
| --- | --- |
| López | Generalli |
| López | Axa |
| Tuercas | Generalli |
| Tuercas | Mutua M |
| Tuercas | Axa |



| service | assurance |
| --- | --- |
| Mechanics | Generalli |
| Mechanics | Axa |
| Painting | Generalli |
| Painting | Mutua M |
| Electronic | Mutua M |
| Electronic | Axa |



| vendor | assurance |
| --- | --- |
| Renault | Generalli |
| Renault | Axa |
| Seat | Generalli |
| Seat | Axa |
| Volvo | Generalli |
| Volvo | Axa |
| Volvo | Mutua M |
| Volkswagen | Mutua M |


