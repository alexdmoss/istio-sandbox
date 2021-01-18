# istio-sandbox

A set of apps and config for me to experiment with Istio features, outside of the safety of the BookInfo sample app that ships with Istio itself.

Some inspiration for ideas from [istiobyexample.dev](https://www.istiobyexample.dev/) as well as the [Istio docs](https://istio.io/latest/docs/tasks/) themselves.

---

## Usage

Each app ships with a Makefile. For development you may also wish to `go get github.com/pilu/fresh` for your `$GOPATH` and then you can benefit from hot reloads by just running `fresh`.

---

## The Plan

- [ ] Basic frontend app
- [ ] Integrate the above with basic backend app
- [ ] CI pipeline to deploy both to separate namespaces
- [ ] Validate Envoy injection
- [ ] Istio features to explore:
  - [ ] Observability
    - [ ] Topology / service graphing
    - [ ] In/out metrics
    - [ ] Access logging
    - [ ] Distributed trace collection
  - [ ] Operability
    - [ ] Route by header
    - [ ] Canarying
    - [ ] Request timeouts
    - [ ] Circuit breaking
    - [ ] Fault injection
    - [ ] Retries
    - [ ] Traffic mirroring
  - [ ] Security
    - [ ] mTLS
    - [ ] validate JWTs
    - [ ] Auth policies
    - [ ] Third party traffic monitoring

---

## Golang Fun

I'm not great with Golang, so a few objectives for me while building these apps too:

- [ ] Move tests out of root - is Go too opinionated for this?
- [ ] Watch Tests equivalent
- [ ] Proper logging
- [ ] Test coverage
- [ ] Clean close of app
