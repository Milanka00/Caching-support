# Caching-support
sample application of response caching with envoy, varnish and Go lang

Access the Backend Server within Docker Container:
  run curl http://localhost:8081/albums from inside the Docker container, the request is routed internally within the container

backend service cluster should be like
          endpoints:
          - lb_endpoints:
              - endpoint:
                  address:
                    socket_address:
                      address: <backend_container_ip>
                      port_value: 8081

access backend container through envoy : curl http://localhost:8080/albums