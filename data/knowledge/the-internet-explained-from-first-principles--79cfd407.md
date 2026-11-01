---
title: "The Internet explained from first principles"
notion_id: 79cfd407-645e-4217-acaa-0083e314d1e2
notion_url: https://app.notion.com/p/The-Internet-explained-from-first-principles-79cfd407645e4217acaa0083e314d1e2
last_edited: 2023-02-23T00:34:00.000Z
source_url: https://explained-from-first-principles.com/internet/
tags: ["English", "Network", "Web Development", "Article", "Book"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you are visiting this website for the first time, then please first read the [front page](https://explained-from-first-principles.com/), where I explain the intention of this blog and how to best make use of it. As far as your privacy is concerned, all data entered on this page is stored locally in your browser unless noted otherwise. While I researched the content on this page thoroughly, you take or omit actions based on it at your own risk. In no event shall I as the author be liable for any damages arising from information or advice on this website or on referenced websites.

## Communication protocol

### Communication diagram

A [communication protocol](https://en.wikipedia.org/wiki/Communication_protocol) specifies how two parties can exchange information for a specific purpose. In particular, it determines which messages are to be transmitted in what order. If the two parties are computers, a formal, well-defined protocol is easiest to implement. In order to illustrate what is going on, however, let’s first look at an informal protocol, also known as [etiquette](https://en.wikipedia.org/wiki/Etiquette), which we’re all too familiar with:

[Alice and Bob](https://en.wikipedia.org/wiki/Alice_and_Bob) engage in the human greeting protocol.

This is a [sequence diagram](https://en.wikipedia.org/wiki/Sequence_diagram). It highlights the temporal dimension of a protocol in which messages are exchanged sequentially.

### Communication parties

It also illustrates that communication is commonly initiated by one party, whereby the recipient responds to the requests of the initiator. Please note that this is only the case for one-to-one protocols, in which each message is intended for a single recipient.

### Communication channel

The above greeting protocol is used among humans to establish a [communication channel](https://en.wikipedia.org/wiki/Communication_channel) for a longer exchange. In technical jargon, such an exchange in preparation for the actual communication is called a [handshake](https://en.wikipedia.org/wiki/Handshaking). The greeting protocol checks the recipient’s availability and willingness to engage in a conversation. When talking to someone you have never spoken before, it also ensures that the recipient understands your language. I’ve chosen these two examples for their figurative value. Why we actually [greet](https://en.wikipedia.org/wiki/Greeting) each other is mainly for different reasons: To show our good intentions by making our presence known to each other, to signal sympathy and courtesy by asking the superficial question, and to indicate our relative social status to each other and to bystanders. Another benefit of asking such a question is that, even though it’s very shallow, it makes the responder more likely to do you a favor due to the psychological effect of [commitment and consistency](https://www.influenceatwork.com/principles-of-persuasion/#consistency).

## Handling of anomalies

### Protocol deviation

Since your communication partner can be erratic, a protocol needs to be able to handle deviations:

Bob gives an unexpected response (in red), from which Alice has to recover (in green).

### Data corruption

Sometimes, data becomes unintelligible in transit, for example due to a lot of [background noise](https://en.wikipedia.org/wiki/Background_noise):

Bob asks Alice (in green) to repeat what he couldn’t understand (in red).

In order to detect transmission errors, computers typically append a [checksum](https://en.wikipedia.org/wiki/Checksum) to each message, which the recipient then verifies. The need to [retransmit](https://en.wikipedia.org/wiki/Retransmission_(data_networks)) messages can be reduced by adding redundancy to messages so that the recipient can [detect and correct small errors](https://en.wikipedia.org/wiki/Error_detection_and_correction) on their own. A simple and very inefficient way of doing this is to repeat the content within each message several times.

### Connection loss

It can also happen that a party loses their connection permanently, for example by moving too far away for the signal to reach the recipient. Since a conversation requires some attention from the communication partner, abandoning a conversation unilaterally without notifying the other party can be misused to block them from talking to someone else for some time. In order to avoid binding resources for a prolonged period of time and thereby potentially falling victim to a so-called [denial-of-service attack](https://en.wikipedia.org/wiki/Denial-of-service_attack), computers drop connections after a configurable duration of inactivity:

Bob terminates the connection after his [timeout period](https://en.wikipedia.org/wiki/Timeout_(computing)).

### Network latency

Other times, your communication partner is simply slow, which needs to be accommodated to some degree:

Bob has a high [network latency](https://en.wikipedia.org/wiki/Latency_(engineering)#Packet-switched_networks) for his [upstream](https://en.wikipedia.org/wiki/Upstream_(networking)) messages (in blue).

### Out-of-order delivery

The following rarely occurs between humans but as soon as messages are passed over various hops, such as forwarding notes among pupils in a classroom, they can arrive [out of order](https://en.wikipedia.org/wiki/Out-of-order_delivery):

Bob’s second message (in blue) arrives after his third message (in green).

The solution for this is to enumerate all messages, to reorder them on arrival, and to ask the other party to retransmit any missing messages, as we saw [above](https://explained-from-first-principles.com/internet/#data-corruption).

### Lack of interoperability

Besides defining the [syntax](https://en.wikipedia.org/wiki/Syntax) (the format), the [semantics](https://en.wikipedia.org/wiki/Semantics) (the meaning), and the order of the messages, a protocol should also specify how to handle anomalies like the above. Ambiguity in a standard and willful deviation therefrom result in incompatibilities between different implementations. In combination with a lack of established standards in many areas, which often leads to uncoordinated efforts by various parties, incompatibilities are quite common in computer systems, unfortunately. This causes a lot of frustration for users and programmers, who have to find workarounds for the encountered limitations, but this cannot be avoided in a free market of ideas and products.

## Network topologies

### Communication network

In practice, there are almost always more than two parties who want to communicate with each other. Together with the connections between them, they form a [communication network](https://en.wikipedia.org/wiki/Telecommunications_network). For the scope of this article, we’re only interested in symmetric networks, where everyone who can receive can also send. This is not the case for analog radio and television networks, where signals are broadcasted unidirectionally from the sender to the receivers. In the case of our symmetric networks, two entities are part of the same network if they can communicate with each other. If they cannot reach each other, they belong to separate networks.

### Nodes and links

[Nodes](https://en.wikipedia.org/wiki/Node_(networking)) are the entities that communicate with each other over communication [links](https://en.wikipedia.org/wiki/Data_link). We can visualize this as follows:

Two nodes (in green) are connected by a link (in yellow).

The terminology is borrowed from [graph theory](https://en.wikipedia.org/wiki/Graph_theory), where nodes are also called vertices and links are also called edges. The technical term for the structure of a network is [topology](https://en.wikipedia.org/wiki/Network_topology). Different arrangements of nodes and links lead to different characteristics of the resulting network.

### Fully connected network

A network is said to be fully connected if every node has a direct link to every other node:

A fully connected network with five nodes and ten links.

In graph theory, such a layout is known as a [complete graph](https://en.wikipedia.org/wiki/Complete_graph). Fully connected networks scale badly as the number of links grows quadratically with the number of nodes. You might have encountered the formula for the number of links before: n × (n – 1) / 2, with n being the number of nodes in the network. As a consequence, this topology is impractical for larger networks.

### Star network

The number of links can be reduced considerably by introducing a central node, which forwards the communication between the other nodes. In such a star-shaped network, the number of links scales linearly with the number of nodes. In other words, if you double the number of nodes, you also double the number of links. In a fully connected network, you would have quadrupled the number of links. For now, we call the newly introduced node a [router](https://en.wikipedia.org/wiki/Router_(computing)). As we will see [later on](https://explained-from-first-principles.com/internet/#hubs-switches-and-routers), such a relaying node is called differently depending on how it operates. Nodes that do not forward the communication of others form the communication endpoints of the network.

A star network with five nodes, five links, and one router (in blue).

While a star network scales optimally, it is by definition totally centralized. If the nodes belong to more than one organization, this topology is not desirable as the central party exerts complete control over the network. Depending on its market power, such a party can increase the price for its service and censor any communication it doesn’t like. Additionally, the central node becomes a [single point of failure](https://en.wikipedia.org/wiki/Single_point_of_failure): If it fails for whatever reason, then the whole network stops working. Since this lowers the [availability](https://en.wikipedia.org/wiki/Availability) of the network, the star topology should not just be avoided for political but also for technical reasons.

### Mesh network

We can avoid these drawbacks by increasing the number of nodes which forward the communication between the endpoints:

A mesh network with six nodes, three routers, and ten links.

In this graph, any of the three routers can go down, and communication is still possible between the nodes that are connected not only to the unavailable router. There are also five links that can break one at a time while leaving all nodes indirectly connected with each other. Such a partially connected network allows for a flexible tradeoff between [redundancy](https://en.wikipedia.org/wiki/Redundancy_(engineering)) and [scalability](https://en.wikipedia.org/wiki/Scalability). It is therefore usually the preferred network topology. Furthermore, the node marked with an asterisk is connected to two routers in order to increase its availability. Because of higher costs, this is usually only done for [critical systems](https://en.wikipedia.org/wiki/Critical_system), which provide crucial services.

## Signal routing

### Network addresses

Unlike in a fully connected network, where each node can simply pick the right link to reach the desired node, a network with relay nodes requires that nodes can address each other. Even if a router relays each signal on all of its links to other nodes, which would make it a [hub](https://explained-from-first-principles.com/internet/#hubs-switches-and-routers) instead of a router, the nodes still need a way to figure out whether they were the intended recipient of a message. This problem can be solved by assigning a unique identifier to each node in the network and by extending each transmitted message with the identifier of the intended recipient. Such an identifier is called a [network address](https://en.wikipedia.org/wiki/Network_address). Routers can learn on which link to forward the communication for which node. This works best when the addresses aren’t assigned randomly but rather reflect the – due to its physical nature often geographical – structure of the network:

Nodes with addresses according to the router they’re connected to. For the sake of simplicity, I no longer draw the arrow tips on links.

We’re all familiar with hierarchical addresses such as [postal codes](https://en.wikipedia.org/wiki/Postal_code), which are known as [ZIP Codes](https://en.wikipedia.org/wiki/ZIP_Code) in the United States, and [telephone numbers](https://en.wikipedia.org/wiki/Telephone_number) with their [country calling codes](https://en.wikipedia.org/wiki/List_of_country_calling_codes). Strictly speaking, the address denotes the network link of a node and not the node itself. This can be seen in the node on the right, which is known as B2 to router B and as C1 to router C. In other words, if a node belongs to several so-called [subnetworks](https://en.wikipedia.org/wiki/Subnetwork), such as B and C in this example, it also has several addresses.

### Routing tables

The process of selecting a path between two nodes across a network is called [routing](https://en.wikipedia.org/wiki/Routing). Routers are the nodes which perform the routing. They maintain a [routing table](https://en.wikipedia.org/wiki/Routing_table) so they know on which link to forward the communication for each node:

| Destination | Link | Cost |
| --- | --- | --- |
| A1 | 1 | 4 |
| A2 | 2 | 2 |
| B? | 3 | 5 |
| B? | 4 | 8 |
| C? | 3 | 9 |
| C? | 4 | 6 |

The routing table for router A. It contains all the destinations to be reached. The links are numbered according to the above graphic.

This table tells router A, for example, to forward all communications for node A2 on link 2. It doesn’t matter on which link router A receives such communications. The router also keeps track of how costly each route is. The cost can either be in terms of [network delay](https://en.wikipedia.org/wiki/Network_delay) or the economic cost of the transmission, based on what providers charge each other. In this example, router A forwards all communications for nodes starting with C on link 4 because the associated cost is lower than the cost for link 3 via router B.

### Routing protocols

Routers and the physical links between them can fail at any time, for example because a network cable is demolished by nearby construction work. On the other hand, new nodes and connections are added to communication networks all the time. Therefore, the routing tables of routers need to be updated continuously. Instead of updating them manually, routers communicate changes with each other using a [routing protocol](https://en.wikipedia.org/wiki/Routing_protocol). For example, as soon as router A detects that it’s no longer getting a response from router C, it updates its routing table to route all communication to C via B:

The link between the routers A and C failed.

| Destination | Link | Cost |
| --- | --- | --- |
| A1 | 1 | 4 |
| A2 | 2 | 2 |
| B? | 3 | 5 |
| C? | 3 | 9 |

The updated routing table of router A with the routes over link 4 removed. With only one route left, router A forwards all communications for C on link 3.

## Signal relaying

A signal can be [relayed through a network](https://en.wikipedia.org/wiki/Switched_communication_network) either with [circuit switching](https://explained-from-first-principles.com/internet/#circuit-switching) or with [packet switching](https://explained-from-first-principles.com/internet/#packet-switching).

### Circuit switching

In a [circuit-switched network](https://en.wikipedia.org/wiki/Circuit_switching), a dedicated [communications channel](https://en.wikipedia.org/wiki/Communications_channel) is established between the two parties for the duration of the [communication session](https://en.wikipedia.org/wiki/Communication_session):

A circuit-switched network with a communication channel (in orange).

The best-known example of a circuit-switched network is the early [telephone network](https://en.wikipedia.org/wiki/Telephone_network). In order to make a call, a [switchboard operator](https://en.wikipedia.org/wiki/Switchboard_operator) needed to connect the wires of the two telephones in order to create a closed circuit. This has the advantage that the delay of the signal remains constant throughout the call and that the communication is guaranteed to arrive in the same order as it was sent. On the other hand, establishing a dedicated circuit for each communication session can be inefficient as others cannot utilize the claimed capacity even when it’s temporarily unused, for example when no one is speaking.

### Packet switching

In a [packet-switched network](https://en.wikipedia.org/wiki/Packet_switching), the data to transfer is split into chunks. These chunks are called [packets](https://en.wikipedia.org/wiki/Network_packet) and consist of a [header](https://en.wikipedia.org/wiki/Header_(computing)) and a [payload](https://en.wikipedia.org/wiki/Payload_(computing)). The header contains information for the delivery of the packet, such as the [network address](https://explained-from-first-principles.com/internet/#network-addresses) of the sender and the recipient. Each router has a queue for incoming packets and then forwards each packet according to its [routing table](https://explained-from-first-principles.com/internet/#routing-tables) or, more precisely, its [forwarding table](https://explained-from-first-principles.com/internet/#forwarding-tables). Apart from these tables, packet-switching routers do not keep any state. In particular, no channels are opened or closed on the routing level.

A packet (in orange) travels through the network from the sender to the recipient.

Since each packet is routed individually, they can take different routes from the sender to the recipient and arrive out of order due to varying delays.

The response from the recipient takes a different route through the network.

Since no router has a complete view of the whole network, it could happen that packets get stuck in an infinite loop:

A packet travels in a circle because of an error in one of the routing tables.

In order to avoid wasting network resources, the header of a packet also contains a counter, which is decreased by one every time it passes a router. If this counter reaches zero before the packet arrives at its destination, then the router discards the packet rather than forwarding it. Such a counter limits the lifespan of a packet by limiting the number of hops it can take and is thus known as its [time-to-live (TTL)](https://en.wikipedia.org/wiki/Time_to_live) value. There are also other reasons why a packet can get lost in the network. The queue of a router might simply be full, which means that additional packets can no longer be stored and must therefore be dropped. Because packets are similar to cars on the road network, some terms are borrowed from the transportation industry. While the capacity of a packet-switched network can be utilized better than the capacity of a circuit-switched network, too much [traffic](https://en.wikipedia.org/wiki/Network_traffic) on the network leads to [congestion](https://en.wikipedia.org/wiki/Network_congestion).

## Internet layers

The [Internet](https://en.wikipedia.org/wiki/Internet) is a global network of computer networks. Its name simply means “[between](https://en.wiktionary.org/wiki/inter-) networks”. It is a [packet-switched](https://explained-from-first-principles.com/internet/#packet-switching) [mesh network](https://explained-from-first-principles.com/internet/#mesh-network) with only [best-effort delivery](https://en.wikipedia.org/wiki/Best-effort_delivery). This means that the Internet provides no guarantees about whether and in what time a packet is delivered. [Internet service providers (ISP)](https://en.wikipedia.org/wiki/Internet_service_provider) provide access to the Internet for businesses and private individuals. They maintain proprietary computer networks for their customers and are themselves interconnected through [international backbones](https://en.wikipedia.org/wiki/Internet_backbone). The big achievement of the Internet is making individual networks interoperable through the [Internet Protocol (IP)](https://en.wikipedia.org/wiki/Internet_Protocol).

The Internet operates in [layers](https://en.wikipedia.org/wiki/Internet_protocol_suite). Each layer provides certain functionalities, which can be fulfilled by different [protocols](https://explained-from-first-principles.com/internet/#communication-protocol). Such a modularization makes it possible to replace the protocol on one layer without affecting the protocols on the other layers. Because the layers above build on the layers below, they are always listed in the following order but then discussed in the opposite order:

| Name | Purpose | Endpoints | Identifier | Example |
| --- | --- | --- | --- | --- |
| [Application layer](https://explained-from-first-principles.com/internet/#application-layer) | Application logic | Application-specific resource | Application-specific | HTTP |
| [Security layer](https://explained-from-first-principles.com/internet/#security-layer) | Encryption and authentication | One or both of the parties | X.509 subject name | TLS |
| [Transport layer](https://explained-from-first-principles.com/internet/#transport-layer) | Typically reliable data transfer | Operating system processes | Port number | TCP |
| [Network layer](https://explained-from-first-principles.com/internet/#network-layer) | Packet routing across the Internet | Internet-connected machines | IP address | IP |
| [Link layer](https://explained-from-first-principles.com/internet/#link-layer) | Handling of the physical medium | Network interface controllers | MAC address | Wi-Fi |

The layers of the Internet. They differ in their purpose, the endpoints that communicate with each other, and how those endpoints are identified.

We will discuss each layer separately in the following subsections. For now, you can treat the above table as an overview and summary. Before we dive into the lowest layer, we first need to understand what “building on the layer below” means. [Digital data](https://en.wikipedia.org/wiki/Digital_data) can be copied perfectly from one memory location to another. The implementation of a specific protocol receives a chunk of data, known as the payload, from the layer above and wraps it with the information required to fulfill its purpose in the so-called header. The payload and the header then become the payload for the layer below, where another protocol specifies a new header to be added. Each of these wrappings is undone by the respective protocol on the recipient side. This can be visualized as follows:

A piece of data flows down through the layers on the sender side and up again on the recipient side.

While this graphic is useful to wrap your head around these concepts, it can be misleading in two ways. Firstly, the payload can be transformed by a specific protocol as long as the original payload can be reconstructed by the recipient. Examples for this are encryption and redundant encoding for automatic [error detection and correction](https://en.wikipedia.org/wiki/Error_detection_and_correction). Secondly, a protocol can split a payload into smaller chunks and transfer them separately. It can even ask the sender to retransmit a certain chunk. As long as all the chunks are recombined on the recipient side, the protocol above can be ignorant about such a process. As we’ve seen in the sections above, a lot of things can go wrong in computer networks. In the following subsections, we’ll have a closer look on how protocols compensate for the deficiencies of the underlying network. Before we do so, we should talk about standardization first.

### Link layer

Protocols on the [link layer](https://en.wikipedia.org/wiki/Link_layer) take care of delivering a packet over a direct link between two nodes. Examples of such protocols are [Ethernet](https://en.wikipedia.org/wiki/Ethernet) and [Wi-Fi](https://en.wikipedia.org/wiki/Wi-Fi). Link layer protocols are designed to handle the intricacies of the underlying physical medium and signal. This can be an electric signal over a copper wire, light over an optical fiber or an electromagnetic wave through space. The node on the other end of the link, typically a router, removes the header of the link layer, determines on the network layer on which link to forward the packet, and then wraps the packet according to the protocol spoken on that link. Link layer protocols typically detect [bit errors](https://en.wikipedia.org/wiki/Bit_error) caused by noise, interference, distortion, and faulty synchronization. If several devices want to send a packet over the same medium at the same time, the signals collide, and the packets must be retransmitted after a randomly chosen [backoff period](https://en.wikipedia.org/wiki/Exponential_backoff).

### Network layer

The purpose of the [network layer](https://en.wikipedia.org/wiki/Internet_layer) is to [route](https://explained-from-first-principles.com/internet/#signal-routing) packets between endpoints. It is the layer that ensures interoperability between separate networks on the Internet. As a consequence, there is only one protocol which matters on this layer: the [Internet Protocol (IP)](https://en.wikipedia.org/wiki/Internet_Protocol). If you want to use the Internet, you have to use this protocol. As we’ve seen earlier, [packet switching](https://explained-from-first-principles.com/internet/#packet-switching) provides only unreliable communication. It is left to the [transport layer](https://explained-from-first-principles.com/internet/#transport-layer) to compensate for this.

The first major version of the Internet Protocol is [version 4 (IPv4)](https://en.wikipedia.org/wiki/IPv4), which has been in use since 1982 and is still the dominant protocol on the Internet. It uses 32-bit numbers to address endpoints and routers, which are written as four numbers between 0 and 255 separated by a dot. These [IP addresses](https://en.wikipedia.org/wiki/IP_address) reflect the hierarchical structure of the Internet, which is important for efficient routing. They are assigned by the [Internet Assigned Numbers Authority (IANA)](https://en.wikipedia.org/wiki/Internet_Assigned_Numbers_Authority), which belongs to the American [Internet Corporation for Assigned Names and Numbers (ICANN)](https://en.wikipedia.org/wiki/ICANN), and by five [Regional Internet Registries (RIR)](https://en.wikipedia.org/wiki/Regional_Internet_registry). If you’re interested, you can check out the [current IPv4 address allocation](https://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.xhtml). There are just under 4.3 billion IPv4 addresses, which are quite unevenly [distributed among countries](https://en.wikipedia.org/wiki/List_of_countries_by_IPv4_address_allocation). Given the limited address space, we’re running out of IPv4 addresses. In order to deal with the [IPv4 address exhaustion](https://en.wikipedia.org/wiki/IPv4_address_exhaustion), the [Internet Protocol version 6 (IPv6)](https://en.wikipedia.org/wiki/IPv6) has been developed. IPv6 uses 128-bit addresses, which are represented as eight groups of four hexadecimal digits with the groups separated by colons. As IPv6 isn’t interoperable with IPv4, the transition has been [slow but steady](https://en.wikipedia.org/wiki/IPv6_deployment).

### Transport layer

### Operating systems

Before we can discuss the [transport layer](https://en.wikipedia.org/wiki/Transport_layer), we first need to talk about [operating systems (OS)](https://en.wikipedia.org/wiki/Operating_system). The job of an operating system is to manage the [hardware](https://en.wikipedia.org/wiki/Computer_hardware) of a computer. Its hardware includes [processors](https://en.wikipedia.org/wiki/Processor_(computing)), such as the [central processing unit (CPU)](https://en.wikipedia.org/wiki/Central_processing_unit) and the [graphics processing unit (GPU)](https://en.wikipedia.org/wiki/Graphics_processing_unit), [memory](https://en.wikipedia.org/wiki/Computer_memory), such as [volatile memory](https://en.wikipedia.org/wiki/Volatile_memory) and [non-volatile memory](https://en.wikipedia.org/wiki/Non-volatile_memory) like your [solid-state drive (SSD)](https://en.wikipedia.org/wiki/Solid-state_drive), [input/output (I/O) devices](https://en.wikipedia.org/wiki/Input/output), such as a [keyboard](https://en.wikipedia.org/wiki/Computer_keyboard) and a [mouse](https://en.wikipedia.org/wiki/Computer_mouse) for input, a [monitor](https://en.wikipedia.org/wiki/Computer_monitor) and [speakers](https://en.wikipedia.org/wiki/Computer_speakers) for output, as well as a [network interface controller (NIC)](https://en.wikipedia.org/wiki/Network_interface_controller) to communicate with other devices on the same network.

An operating system fulfills three different purposes:

- **Abstraction**: It simplifies and standardizes the access to the hardware, making it easier for engineers to develop software for several [computing platforms](https://en.wikipedia.org/wiki/Computing_platform).
- **Duplication**: It provides the same resources to several programs running on the same computer, thereby giving each program the illusion that it has the hardware just for itself.
- **Protection**: It enforces restrictions on the behavior of programs. For example, it can deny access to the webcam or certain parts of the [file system](https://en.wikipedia.org/wiki/File_system) unless the user has granted the necessary permissions.

### Port numbers

When a program is being executed, it is called a [process](https://en.wikipedia.org/wiki/Process_(computing)). This distinction is important because the same program can be executed several times in parallel, which results in several processes until they terminate. Since more than one process might want to use the network connection at the same time, the operating system needs a way to keep the traffic of different processes apart. The label used for this purpose is a 16-bit integer known as [port number](https://en.wikipedia.org/wiki/Port_(computer_networking)). When a process sends a request to another device, the operating system chooses an arbitrary but still unused port number and encodes it as the source port in the transport layer wrapping of the outgoing packet. The recipient then has to include the same port number as the destination port in its response. When the operating system of the requester receives this response, it knows which process to forward the incoming packet to because it kept track of which port numbers it used for which process.

But how does the operating system of the recipient know what to do with the incoming packet? The answer is registration and convention. A process can ask the operating system to receive all incoming packets which have a certain destination port. If no other process has claimed this port before, the operating system grants this port to the process. A port can be bound to at most one process. If it is already taken, then the operating system returns an error. Ports are distributed on a first-come, first-served basis. To claim port numbers below 1024, processes need a special privilege, though. Which port to claim as a receiving process is handled by convention. Each [application layer](https://explained-from-first-principles.com/internet/#application-layer) protocol defines one or several default ports to receive traffic on. Wikipedia has an extensive [list of established port numbers](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers).

An application process registers the port 25 at the operating system and then receives a packet on this port.

### Client-server model

A [server](https://en.wikipedia.org/wiki/Server_(computing)) is just a process registered with the operating system to handle incoming traffic on a certain port. It does this to provide a certain service, which is then requested by so-called [clients](https://en.wikipedia.org/wiki/Client_(computing)). This is called the [client-server model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model), which contrasts with a [peer-to-peer architecture](https://en.wikipedia.org/wiki/Peer-to-peer), where each node equally provides and consumes the service. The communication is always initiated by the client. If the server makes a request itself, it becomes the client in that interaction. A server is typically accessed via a network like the Internet but it can also run on the same machine as its client. In such a case, the client accesses the server via a so-called [loopback](https://en.wikipedia.org/wiki/Loopback), which is a virtual network interface where the destination is the same as the source. The current computer is often referred to as [localhost](https://en.wikipedia.org/wiki/Localhost). There is also a dedicated IP address for this purpose: `127.0.0.1` in the case of IPv4 and `::1` in the case of IPv6.

### Security layer

All the communication we have seen so far is neither authenticated nor encrypted. This means that any router can read and alter the messages that pass through it. Since the network determines the route of the packets rather than you as a sender, you have no control over which companies and nations are involved in delivering them. The lack of confidentiality is especially problematic when using the Wi-Fi in a public space, such as a restaurant or an airport, because your device simply connects to the [wireless access point](https://en.wikipedia.org/wiki/Wireless_access_point) of a [given network](https://en.wikipedia.org/wiki/Service_set_(802.11_network)#SSID) with the best signal. Since your device has no way to authenticate the network, anyone can impersonate the network and then inspect and modify your traffic by setting up a fake access point. This is known as an [evil twin attack](https://en.wikipedia.org/wiki/Evil_twin_%28wireless_networks%29), which also affects [mobile phone networks](https://www.youtube.com/watch?v=fQSu9cBaojc). As a general principle, you should never trust the [network layer](https://explained-from-first-principles.com/internet/#network-layer).

### Application layer

Everything we’ve covered so far serves a single purpose: to accomplish things we humans are interested in. This is done with protocols on the [application layer](https://en.wikipedia.org/wiki/Application_layer). Examples of application layer protocols are the [HyperText Transfer Protocol (HTTP)](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) as the foundation of the [World Wide Web (WWW)](https://en.wikipedia.org/wiki/World_Wide_Web), the [Simple Mail Transfer Protocol (SMTP)](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol) for delivering [email](https://en.wikipedia.org/wiki/Email), the [Internet Message Access Protocol (IMAP)](https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol) for retrieving email, and the [File Transfer Protocol (FTP)](https://en.wikipedia.org/wiki/File_Transfer_Protocol) for, as you can probably guess, transferring files. What all of these protocols have in common is that they all use a [text-based format](https://explained-from-first-principles.com/internet/#text-based-protocols), that they all run over [TCP](https://explained-from-first-principles.com/internet/#transmission-control-protocol), and that they all have a secure variant running over [TLS](https://explained-from-first-principles.com/internet/#transport-layer-security), namely [HTTPS](https://en.wikipedia.org/wiki/HTTPS), [SMTPS](https://en.wikipedia.org/wiki/SMTPS), [IMAPS](https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol#Security), and [FTPS](https://en.wikipedia.org/wiki/FTPS). This is the beauty of modularization: Application layer protocols can reuse the same protocols below, while the protocols below don’t need to know anything about the protocols above.

## Internet history

There are many nice articles about the [history of the Internet](https://en.wikipedia.org/wiki/History_of_the_Internet) and there’s no point in replicating their content here. Instead, I would like to give you a timeline of important milestones in the history of [telecommunication](https://en.wikipedia.org/wiki/History_of_telecommunication) and [computing](https://en.wikipedia.org/wiki/History_of_computing_hardware):

| Year | Description |
| --- | --- |
| 1816 | First working [electrical telegraph](https://en.wikipedia.org/wiki/Electrical_telegraph) built by the English inventor [Francis Ronalds](https://en.wikipedia.org/wiki/Francis_Ronalds). |
| 1865 | Adoption of the [Morse code](https://en.wikipedia.org/wiki/Morse_code), which originated in 1837, as an international standard. |
| 1876 | [Alexander Graham Bell](https://en.wikipedia.org/wiki/Alexander_Graham_Bell) receives the first patent for a [telephone](https://en.wikipedia.org/wiki/Telephone) in the United States. |
| 1941 | Invention of the [Z3](https://en.wikipedia.org/wiki/Z3_(computer)), the first programmable computer, by [Konrad Zuse](https://en.wikipedia.org/wiki/Konrad_Zuse) in Germany. |
| 1945 | Invention of the [ENIAC](https://en.wikipedia.org/wiki/ENIAC), the first computer with [conditional branching](https://en.wikipedia.org/wiki/Conditional_(computer_programming)), in the US. |
| 1954 | Invention of [time-sharing](https://en.wikipedia.org/wiki/Time-sharing) (share expensive computing resources among several users). |
|  | Increased interest in remote access for users because computers were huge and rare. |
| 1965 | Invention of [packet switching](https://explained-from-first-principles.com/internet/#packet-switching) at the [National Physical Laboratory (NPL)](https://en.wikipedia.org/wiki/National_Physical_Laboratory_(United_Kingdom)) in the UK. |
| 1969 | The [US Department of Defense](https://en.wikipedia.org/wiki/United_States_Department_of_Defense) initiates and funds the development of the [ARPANET](https://en.wikipedia.org/wiki/ARPANET). |
|  | Similar networks are built in London ([NPL](https://en.wikipedia.org/wiki/NPL_network)), Michigan ([MERIT](https://en.wikipedia.org/wiki/Merit_Network)), and France ([CYCLADES](https://en.wikipedia.org/wiki/CYCLADES)). |
| 1972 | [Jon Postel](https://en.wikipedia.org/wiki/Jon_Postel) establishes himself as [the czar of socket numbers](https://datatracker.ietf.org/doc/html/rfc349), which leads to the [IANA](https://en.wikipedia.org/wiki/Internet_Assigned_Numbers_Authority). |
| 1973 | [Bob Kahn](https://en.wikipedia.org/wiki/Bob_Kahn) and [Vint Cerf](https://en.wikipedia.org/wiki/Vint_Cerf) publish research on [internetworking](https://en.wikipedia.org/wiki/Internetworking) leading to IP and TCP. |
| 1978 | Public discovery of the [first public-key cryptosystem](https://en.wikipedia.org/wiki/RSA_(cryptosystem)) for encryption and signing,which was [already discovered](https://en.wikipedia.org/wiki/Public-key_cryptography#Classified_discovery) in 1973 at the British intelligence agency [GCHQ](https://en.wikipedia.org/wiki/Government_Communications_Headquarters). |
| 1981 | Initial release of the text-based [MS-DOS](https://en.wikipedia.org/wiki/MS%2DDOS) by [Microsoft](https://en.wikipedia.org/wiki/Microsoft), licensed by [IBM](https://en.wikipedia.org/wiki/IBM) for its [PC](https://en.wikipedia.org/wiki/IBM_PC_compatible). |
| 1982 | The US Department of Defense makes IP the [only approved protocol on ARPANET](https://en.wikipedia.org/wiki/Internet_protocol_suite#Adoption). |
| 1982 | First definition of the [Simple Mail Transfer Protocol (SMTP)](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol) for email in [RFC 821](https://datatracker.ietf.org/doc/html/rfc821). |
| 1983 | Creation of the [Domain Name System (DNS)](https://en.wikipedia.org/wiki/Domain_Name_System) as specified in [RFC 882](https://datatracker.ietf.org/doc/html/rfc882) and [RFC 883](https://datatracker.ietf.org/doc/html/rfc883). |
| 1984 | Version 1 of the [Post Office Protocol (POP)](https://en.wikipedia.org/wiki/Post_Office_Protocol) to fetch emails from a mailbox ([RFC 918](https://datatracker.ietf.org/doc/html/rfc918)). |
| 1985 | [First commercial registration](https://en.wikipedia.org/wiki/Domain_name#Domain_name_registration) of a domain name in the `.com` [top-level domain](https://en.wikipedia.org/wiki/Top-level_domain). |
| 1986 | Design of the [Internet Message Access Protocol (IMAP)](https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol), documented in [RFC 1064](https://datatracker.ietf.org/doc/html/rfc1064). |
| 1990 | Invention of the [World Wide Web](https://en.wikipedia.org/wiki/World_Wide_Web) by [Tim Berners-Lee](https://en.wikipedia.org/wiki/Sir_Timothy_John_Berners-Lee) at [CERN](https://en.wikipedia.org/wiki/CERN) in Switzerland,which includes the [HyperText Transfer Protocol (HTTP)](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol), the [HyperText Markup](https://en.wikipedia.org/wiki/Hypertext_Markup_Language)[Language (HTML)](https://en.wikipedia.org/wiki/Hypertext_Markup_Language), the [Uniform Resource Locator (URL)](https://en.wikipedia.org/wiki/Uniform_resource_locator), a [web server](https://en.wikipedia.org/wiki/Web_server), and a [browser](https://en.wikipedia.org/wiki/Web_browser). |
| 1993 | Specification of the [Dynamic Host Configuration Protocol (DHCP)](https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol) in [RFC 1541](https://datatracker.ietf.org/doc/html/rfc1541). |
| 1995 | Release of the [Secure Sockets Layer (SSL)](https://en.wikipedia.org/wiki/Transport_Layer_Security) by [Netscape](https://en.wikipedia.org/wiki/Netscape), renamed to TLS in 1999. |
| 1995 | Standardization of [IPv6](https://en.wikipedia.org/wiki/IPv6) by the IETF in [RFC 1883](https://datatracker.ietf.org/doc/html/rfc1883), obsoleted by [RFC 2460](https://datatracker.ietf.org/doc/html/rfc2460) in 1998. |
| 1998 | [Google](https://en.wikipedia.org/wiki/Google) is founded by [Larry Page](https://en.wikipedia.org/wiki/Larry_Page) and [Sergey Brin](https://en.wikipedia.org/wiki/Sergey_Brin) at [Stanford University](https://en.wikipedia.org/wiki/Stanford_University) in California. |
| 2005 | Specification of [DNSSEC](https://en.wikipedia.org/wiki/Domain_Name_System_Security_Extensions) in [RFC 4033](https://datatracker.ietf.org/doc/html/rfc4033), [4034](https://datatracker.ietf.org/doc/html/rfc4034) & [4035](https://datatracker.ietf.org/doc/html/rfc4035) after [earlier attempts](https://datatracker.ietf.org/doc/html/rfc2535) in 1995. |
| 2007 | [Apple](https://en.wikipedia.org/wiki/Apple_Inc.) launches the [iPhone](https://en.wikipedia.org/wiki/IPhone) with the [iOS](https://en.wikipedia.org/wiki/IOS) operating system one year before [Android](https://en.wikipedia.org/wiki/Android_(operating_system)). |
| 2010 | Deployment of DNSSEC [in the root zone](https://en.wikipedia.org/wiki/Domain_Name_System_Security_Extensions#Deployment_at_the_DNS_root), eliminating intermediary [trust anchors](https://en.wikipedia.org/wiki/Trust_anchor). |
| 2018 | The [UN](https://en.wikipedia.org/wiki/International_Telecommunication_Union) estimates that [more than half](https://www.itu.int/en/ITU-D/Statistics/Documents/facts/FactsFigures2019.pdf) of the global population uses the Internet. |
