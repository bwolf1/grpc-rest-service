# TODO.md

Project Task List

### Todo

- [ ] Add a REVERSE bool env var that when true and passed to the Echo RPC will echo the string reversed  
  - [ ] This is how we use env vars in the code <https://go.dev/play/p/G0i481VEcBp>  
- [ ] Add Zap logging  
- [ ] Add Prometheus and Grafana  
  - [ ] [This article](https://adamtheautomator.com/prometheus-kubernetes/) might help  
- [ ] Add README instructions for bringing up the service with prometheus and grafana  
- [ ] Add a database  
  - [ ] Add logic for inserting the request payload into the database and using the record ID in the response body  
  - [ ] [This article](https://www.sohamkamani.com/golang/sql-database/) might be helpful  
- [ ] Add log storage in gcs  
- [ ] Update to the latest versions of everything  

### In Progress

- [ ] Add Docker  
- [ ] Add Kubernetes  
- [ ] Add Pulumi  
- [ ] Add README instructions for bringing up everything in Kubernetes  

### Done ✓

- [x] Add GitHub actions to build and run tests  
- [x] Add README instructions for bringing up the gRPC service  
- [x] Add README instructions for bringing up the gRPC client  
- [x] Add README instructions for bringing up the REST gateway  
- [x] Add the gRPC service  
- [x] Add the gRPC client  
- [x] Add the REST gateway  

