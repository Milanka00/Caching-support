# default.vcl

vcl 4.1;

backend default {
    .host = "backend";
    .port = "8081";
}

sub vcl_recv {
    if (req.url == "/albums") {
        return (hash);
    }
}

sub vcl_backend_response {
    if (bereq.url == "/albums") {
        set beresp.ttl = 1m;  # Set cache TTL to 1 min
    }
}
