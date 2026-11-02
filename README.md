![ci](https://github.com/fjbredenkamp/k-mini/actions/workflows/ci.yml/badge.svg)
# k-mini

A tiny container runtime and orchestrator in Go, built from scratch to learn
how Kubernetes works underneath: processes, namespaces, cgroups, and the
reconciliation loop.

## Roadmap

- [x] **v0.0** Project setup, CI
- [ ] **v0.1** Process supervisor: /proc reader, health probes, restart loop, HTTP status
- [ ] **v0.2** Containers from scratch: namespaces, pivot_root, cgroups v2, networking
- [ ] **v0.3** Control plane: declarative YAML spec, SQLite state, reconciler, CLI
- [ ] **Next** A real Kubernetes operator, built on the same idea, running in my homelab

## Build

    make          # lint, test, build
    make run