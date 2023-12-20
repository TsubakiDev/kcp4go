# kcp4go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/TsubakiDev/kcp4go/blob/master/LICENSE)

**kcp4go** is a native [KCP](https://github.com/skywind3000/kcp) implementation for the Golang platform. KCP is a low-latency, high-reliability protocol designed for scenarios with real-time requirements, such as online gaming.

## Features

- High Performance: Implemented in native Golang for excellent performance.
- Low Latency: KCP protocol is designed to minimize network transmission latency.
- High Reliability: Enhances data transmission reliability through mechanisms like FEC (Forward Error Correction) and ARQ (Automatic Repeat reQuest).

## Installation

Install kcp4go using the `go get` command:

```bash
go get -u github.com/TsubakiDev/kcp4go
