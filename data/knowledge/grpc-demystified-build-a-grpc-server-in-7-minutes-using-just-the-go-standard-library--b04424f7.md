---
title: "grpc demystified: Build a gRPC server in 7 minutes - using just the Go standard library!"
notion_id: b04424f7-40c6-4c7e-b95f-b42442ef34a8
notion_url: https://app.notion.com/p/grpc-demystified-Build-a-gRPC-server-in-7-minutes-using-just-the-Go-standard-library-b04424f740c64c7eb95fb42442ef34a8
last_edited: 2023-01-13T17:38:00.000Z
source_url: https://github.com/akshayjshah/grpc-demystified/#
tags: ["Guide", "English", "Backend", "Go"]
---
# gRPC Demystified

Details

This repository contains the slides and code for a lightning talk I hope to give at Gophercon 2022. In the talk, we build a gRPC server — from scratch — using just the Go standard library.

The slides are available in [Keynote](https://github.com/akshayjshah/grpc-demystified/blob/main/grpc-demystified.key) or [PDF](https://github.com/akshayjshah/grpc-demystified/blob/main/grpc-demystified.pdf) format. There's also a five-minute [recording](https://github.com/akshayjshah/grpc-demystified/blob/main/grpc-demystified.mp4) of me practicing the talk.

The code includes a [REST handler](https://github.com/akshayjshah/grpc-demystified/blob/main/rest.go) and a from-scratch [gRPC handler](https://github.com/akshayjshah/grpc-demystified/blob/main/grpc.go), both implementing the same logic. There's also a client for each, along with a `grpc-go` client to show that our handler is speaking the wire protocol correctly. To start the HTTP server and make a request with each client, `go run .`.

If this talk appeals to you, the [Connect](https://connect.build/) RPC framework may be right up your alley.
