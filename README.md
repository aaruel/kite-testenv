# Kontrol TestEnv

A full-stack environment following the microservice architecture using Kite as a service framework and Kontrol for cross-service communication and delegation.

---

Stack:
- Backend
  - Golang
    - [Kite](https://github.com/koding/kite)
    - [Kontrol](https://github.com/koding/kite)
    - [Chi](https://github.com/go-chi/chi)
- Frontend
  - Javascript
    - [CRA](https://github.com/facebook/create-react-app)


---

### How it works

The flow is pretty simple:

```
Kontrol <--> Microservices <--> API Gateway <--> Client
```

Each microservice registers with Kontrol so it can be found by other microservices through a query. The Gateway itself is a microservice so it can communicate with all existing microservices through Kontrol, as well as expose an outward facing API to coordinate requests from the client.

---

### The Test Environment

**Kontrol** is installed by `go get github.com/koding/kite/kontrol/kontrol`, needing SSL keys generated. See `run-kontrol.sh` for relevant environment variables.

**Microservices** are located in the `server` folder. In this case there's only one, but can be easily expanded using the same flow exemplified in the `main.go` file.

**API Gateway** is located in the `gateway` folder. This is split into two parts: a repo connecting to all the services making it easy to call methods by passing a service name, a method name, and arguments from anywhere; and a router which gives clients a centralized address to fetch data from the microservices.

**Client** is simply what's used to test connectivity to the gateway, using create-react-app and a test request to the only existing method, `square`. 