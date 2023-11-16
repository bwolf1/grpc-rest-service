# TODO

- Add Kubernetes  
- Add Pulumi  
- Add README instructions for bringing up everything with Pulumi  
- Make repo public  
- General README cleanup  
- Add a REVERSE bool env var that when true and passed to the Echo RPC will echo the string reversed  
  - This is how we use env vars in the code <https://go.dev/play/p/G0i481VEcBp>  
- Add a database  
  - Add logic for inserting the request payload into the database and using the record ID in the response body  
  - [This article](https://www.sohamkamani.com/golang/sql-database/) might be helpful  
- Add Prometheus and Grafana  
  - [This article](https://adamtheautomator.com/prometheus-kubernetes/) might help  
- Add README instructions for bringing up the service with prometheus and grafana  
- Add log storage in gcs  
- Update to the latest versions of everything
