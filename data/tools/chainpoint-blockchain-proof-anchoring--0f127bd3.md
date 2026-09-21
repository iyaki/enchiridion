---
title: "Chainpoint - Blockchain Proof & Anchoring"
notion_id: 0f127bd3-3b55-43fd-857f-c89f084b82c2
notion_url: https://app.notion.com/p/Chainpoint-Blockchain-Proof-Anchoring-0f127bd33b5543fd857fc89f084b82c2
last_edited: 2022-12-21T15:13:00.000Z
source_url: https://chainpoint.org/
tags: ["Tool", "Service", "English", "Databases", "Information Security", "Programming"]
---


## Chainpoint is an open standard for creating a timestamp proof of any data, file, or process.

## Anchor an unlimited amount of data to the Bitcoin blockchain. Verify the integrity and existence of data without relying on a trusted third-party.

...plus hundreds of others use Chainpoint.

## How does Chainpoint work?

Chainpoint links a hash of your data to a blockchain and returns a timestamp proof. A Chainpoint Node receives hashes which are aggregated together using a Merkle tree. The root of this tree is published in a Bitcoin transaction. The final Chainpoint proof defines a set of operations that cryptographically link your data to the Bitcoin blockchain.

## What is a Chainpoint Proof?

A Chainpoint proof contains the information to verify that the hash of some data is anchored to a blockchain. It proves the data existed at time it was anchored. Chainpoint proofs can be verified without reliance on a trusted third party.

## Chainpoint Node API

Chainpoint Nodes have an HTTP API. You can interact with Nodes using any programming language that supports HTTP calls. Thousands of Chainpoint Nodes operate as part of a global network.

Chainpoint Node API Tutorial

## Chainpoint JS Client

This Chainpoint javascript client can be used in both Browser and Node.js based Javascript applications using callback functions, Promises (using .then, .catch), or Promises (using async/await) functional styles.

Chainpoint JS Client

## Chainpoint Proof Versions

The third major version of the Chainpoint proof specification is currently in testing and scheduled for release soon. A JSON schema validator is available.

- Version 3.0
- Version 2.0
- Version 1.x

Name Description @context string, required the JSON-LD context for the proof type string, required the JSON-LD type definition hash string, required hash value between 40 and 128 hex characters. Must be even length. hash_id_node string, required a Version 1 UUID with embedded timestamp and 5 byte BLAKE2s hash of the input data. Timestamp represents Node server time (UTC) of hash submission. hash_submitted_node_at string, required Human readable ISO 8601 timestamp extracted from time embedded in the hash_id_node hash_id_core string, required a Version 1 UUID with embedded timestamp and 5 byte BLAKE2s hash of the input data. Timestamp represents Core server time (UTC) of hash submission. hash_submitted_core_at string, required Human readable ISO 8601 timestamp extracted from time embedded in the hash_id_core branches - an array of branch objects. Branches can be nested without limit and MUST be traversed only after executing 'ops'. (required only at root) label string, optional text string representing the branch name ops array, optional an array of operations objects. Operations are performed in sequence to arrive at an intermediate hash prior to entering a nested branch. branches array, optional nested array of branch objects. Each branch contains ops; labels and additional nested branches are optional. ops - an array of operation objects (required under every 'branches' object) lstring, optional left concatenate a value. If the value is a hexadecimal string, it will be read as a hexadecimal byte array, otherwise the string will be converted to its byte value assuming UTF-8 encoding. rstring, optional right concatenate a value. If the value is a hexadecimal string, it will be read as a hexadecimal byte array, otherwise the string will be converted to its byte value assuming UTF-8 encoding. op string, optional an operation to perform on the current value combined with a previous 'l' or 'r' operation. Current operations: 'sha-224', 'sha-256', 'sha-384', 'sha-512', 'sha3-224', 'sha3-256', 'sha3-384', 'sha3-512', or the special purpose 'sha-256-x2' which applies 'sha-256' twice. anchors - an array of anchor objects (required under every 'ops' object). type string, required one of 'cal' (Calendar), 'btc' (Bitcoin), 'tcal' (Testnet Calendar), 'tbtc' (Testnet Bitcoin), or 'eth' (Ethereum) anchor types anchor_id string, required an identifier used to look up embedded anchor data. e.g. a Bitcoin transaction or block ID. uris array, optional an array of special purpose string URI's, each of which can be used to lookup and retrieve the exact hash resource required to validate this anchor. The URI MUST return only a Hexadecimal hash value as a string. The URI MUST also contain the 'anchor_id' value to lookup the URI resource. This strict requirement is to allow automated clients to retrieve and validate intermediate hashes when verifying a proof. The body value returned by the URI MUST be of even length and match the regex [a-fA-F0-9].

### Example

```plain text
{ "@context": "https://w3id.org/chainpoint/v3", "type": "Chainpoint", "hash": "ffff27222fe366d0b8988b7312c6ba60ee422418d92b62cdcb71fe2991ee7391", "hash_id_node": "66a34bd0-f4e7-11e7-a52b-016a36a9d789", "hash_submitted_node_at": "2018-01-09T02:47:15Z", "hash_id_core": "66bd6380-f4e7-11e7-895d-0176dc2220aa", "hash_submitted_core_at": "2018-01-09T02:47:15Z", "branches": [ { "label": "cal_anchor_branch", "ops": [ { "l": "node_id:66a34bd0-f4e7-11e7-a52b-016a36a9d789" }, { "op": "sha-256" }, { "l": "core_id:66bd6380-f4e7-11e7-895d-0176dc2220aa" }, { "op": "sha-256" }, { "l": "nist:1515465960:1041862e0f3987dca3aab3a91767d2a2ebbf251451b740879adb0926f0ee325e608d5c311e3f64a002dc5266337efc34ebdbf0032c7a253a8fbb64c1b0fb625f" }, { "op": "sha-256" }, { "r": "725a969557e64600aa2bbe50e75fc12dd913620144660836441a97f6d36babf9" }, { "op": "sha-256" }, { "l": "f21aac3945aee46d0cd888faff3364cc7640f88c9bdfefb1072a4bb82c6702b6" }, { "op": "sha-256" }, { "r": "c59058f17b93b609f4b49366c8808099a715836b6c08b45a1dc6ac762820ae27" }, { "op": "sha-256" }, { "l": "985635:1515466042:1:https://a.chainpoint.org:cal:985635" }, { "r": "0e20cff025777bec277cd3a0599eaf5efbeb1ea7adf5ec5a39126a77fa57f837" }, { "op": "sha-256" }, { "anchors": [ { "type": "cal", "anchor_id": "985635", "uris": [ "https://a.chainpoint.org/calendar/985635/hash" ] } ] } ], "branches": [ { "label": "btc_anchor_branch", "ops": [ { "l": "0e20cff025777bec277cd3a0599eaf5efbeb1ea7adf5ec5a39126a77fa57f837" }, { "op": "sha-256" }, { "r": "9d7e8027c869d7446db8f2a5f371d967f5ba9d3a88f1703a1674f57963d3448d" }, { "op": "sha-256" }, { "l": "28c6aa4416d1b0aa474bc52fd32175ec7d15980772874617b5000aff043ac6cb" }, { "op": "sha-256" }, { "r": "4c297218f2015d4f84a6561ca06c1c28b2f6cca1500315ef6d4944ad6822b974" }, { "op": "sha-256" }, { "r": "f6a15401357e6e177583dbf5aa82b5ed5ae1043d1bda3faba88ca0fdb90e01c0" }, { "op": "sha-256" }, { "r": "ae9137386a03fdcdb9a1554a6e4fcd9697efed17caaa0221ce35e12bfc9fbf2d" }, { "op": "sha-256" }, { "l": "fa5643778470a9175644affe35e0177a13b2446d73182be0963d53b1d09214ab" }, { "op": "sha-256" }, { "l": "01000000013d9bfb8c553b3a7c9c030ea9b0f47c7e4c457e47a1ad2d9c751c8eb0e02fee70010000006a47304402201eac07288c3881f354564bb9da0d8267174cdc9e8c42ca82c2129a0416c806220220104e9932a89259472c84be7722f77324efa43a65ca79dd5bb8b6aab0ac9788000121032695ca0d3c0f7f8082a6ef66e7127e48d4eb99bef86be99432b897c485962fa8ffffffff020000000000000000226a20" }, { "r": "ca694202000000001976a9149f1f4038857beedd34cc5ba9f26ac7a20c04d51988ac00000000" }, { "op": "sha-256-x2" }, { "l": "aa7008cdf722a674cc3532727ee39e9ebc810fb047cc7f4edc302705fcee3985" }, { "op": "sha-256-x2" }, { "l": "f0fae6f1dc00b678596e230584430b95bad9c1439f03293250b5a9bfb993b500" }, { "op": "sha-256-x2" }, { "l": "a79b18abcde7db6554e95c14ed544231f59670318033fc6e2e28142341ef223a" }, { "op": "sha-256-x2" }, { "l": "12105db21e488b1d8eb44fbce8bc5e3fcb7becc35fe4d9d30696ef7baff853eb" }, { "op": "sha-256-x2" }, { "l": "0ce1848d74ea8705858e468e045e7891f2b5f9c8ed37eeaa00be51846460294e" }, { "op": "sha-256-x2" }, { "l": "52af6b21e7b370f680e984b8a1e34ffdb45770d3cf599357ce245bad8c820d50" }, { "op": "sha-256-x2" }, { "l": "bb5bd9669a3bc3202e460091185f8103863da4263f417e85479fc3bb40a882d1" }, { "op": "sha-256-x2" }, { "l": "25bb84e8a36904224182b28adb04956d1251d4312b4e975c4ee3ff74a50bce1d" }, { "op": "sha-256-x2" }, { "l": "a55b52dc8079febc3a8b673ee123829c176aca7dabb330299afdeac2bfea16d6" }, { "op": "sha-256-x2" }, { "r": "3bf18e7d4ffaab9988d14b1402fe9817ea6c50fa626dd78bcaba18a9b16184f1" }, { "op": "sha-256-x2" }, { "r": "af9ae1010333cf6e5ea124e5827a8bf0f40f68ab9a5bf283f93f744046b07a5d" }, { "op": "sha-256-x2" }, { "anchors": [ { "type": "btc", "anchor_id": "503275", "uris": [ "https://a.chainpoint.org/calendar/985814/data" ] } ] } ] } ] } ] }
```

Name Description @context the JSON-LD context for the receipt type receipt type definition specifying hash method and version targetHash hash value being anchored to the blockchain merkleRoot merkle tree root value that is anchored to the blockchain proof merkle proof establishing link from the targetHash to the merkleRoot anchors type anchor type definition specifying anchoring method sourceId identifier, such as a transaction id, used to locate anchored data

### Example

```plain text
{ "@context": "https://w3id.org/chainpoint/v2", "type": "ChainpointSHA256v2", "targetHash": "bdf8c9bdf076d6aff0292a1c9448691d2ae283f2ce41b045355e2c8cb8e85ef2", "merkleRoot": "51296468ea48ddbcc546abb85b935c73058fd8acdb0b953da6aa1ae966581a7a", "proof": [ { "left": "bdf8c9bdf076d6aff0292a1c9448691d2ae283f2ce41b045355e2c8cb8e85ef2" }, { "left": "cb0dbbedb5ec5363e39be9fc43f56f321e1572cfcf304d26fc67cb6ea2e49faf" }, { "right": "cb0dbbedb5ec5363e39be9fc43f56f321e1572cfcf304d26fc67cb6ea2e49faf" } ], "anchors": [ { "type": "BTCOpReturn", "sourceId": "f3be82fe1b5d8f18e009cb9a491781289d2e01678311fe2b2e4e84381aafadee" } ] }
```

### Receipt Types

Receipt type values indicate the hash type and version of the receipt. The following values are supported.

Type Description ChainpointSHA224v2 Chainpoint 2.0 receipt using SHA-224 ChainpointSHA256v2 Chainpoint 2.0 receipt using SHA-256 ChainpointSHA384v2 Chainpoint 2.0 receipt using SHA-384 ChainpointSHA512v2 Chainpoint 2.0 receipt using SHA-512 ChainpointSHA3-224v2 Chainpoint 2.0 receipt using SHA3-224 ChainpointSHA3-256v2 Chainpoint 2.0 receipt using SHA3-256 ChainpointSHA3-384v2 Chainpoint 2.0 receipt using SHA3-384 ChainpointSHA3-512v2 Chainpoint 2.0 receipt using SHA3-512

### Anchor Types

Anchor type values indicate the method and location of the anchored data. The following values are supported. Support for additional methods is planned.

Type Description BTCOpReturn Anchored to a Bitcoin transaction within an OP_RETURN output. ETHData Anchored to an Ethereum transaction within the data field.

Name Description Header chainpoint_version version of the Chainpoint standard hash_type hashing algorithm used to encrypt target data (sha­256) merkle_root root of the Merkle Tree that is published in the blockchain tx_id blockchain transaction id timestamp non­authoritative Unix timestamp of the target Target target_hash hash of the target that is being recorded in the blockchain target_proof Merkle proof used to prove target_hash is part of Merkle tree target_URI (optional) path to the target Extra custom (optional) array of user defined key value pairs

### Example

```plain text
{ "header": { "chainpoint_version": "1.1", "hash_type": "SHA-256", "merkle_root": "8dbd52c5ff89b70711b06c143520ef4eb295c51040757c7a4ab56303f0f6b68f", "tx_id": "77d93bce0f0ff76f1a52424f9c0aeb165990dc37a5e3aaf85031a2f6ab2967d1", "timestamp": 1438317621 }, "target": { "target_hash": "a64f10ca86a880115cc271232dc8577606dc6223cf009a4dd1290cd55b2d6a28", "target_uri ": "https://www.someurl.com/target/id", "target_proof": [ { "parent": "ee32ac6eb702c289ba2a932e6562b2c45070121cacb26cabdee433a6cf2086cd", "left": "a64f10ca86a880115cc271232dc8577606dc6223cf009a4dd1290cd55b2d6a28", "right": "f07934733450ba14972b8fbea4c39ac05b0b8e2d8a45ddde9dbe03bb70942b6a" }, { "parent": "8dbd52c5ff89b70711b06c143520ef4eb295c51040757c7a4ab56303f0f6b68f", "left": "ee32ac6eb702c289ba2a932e6562b2c45070121cacb26cabdee433a6cf2086cd", "right": "07084efda9cf5488ef9d4f9ad8bca8521581aaa32331b62aeb2fe3e132e62d1a" } ] }, "extra": [ { "custom_key_1": "value_1" }, { "custom_key_2": "value_2" } ] }
```

## Verifying Chainpoint Proofs

Verification confirms that the proof is well formatted, and all proof operations lead to the expected anchor hash on the blockchain. Verification of version 3.x proofs can be performed with the Chainpoint CLI. Verification of older versions can be performed using the chainpoint-validate Javascript package for Node.js.
