# default.vcl

vcl 4.0;

backend default {
    .host = "backend";   //actual backend ip
    .port = "8081";
}

sub vcl_recv {
       if (req.url ~ "^/albums") {
        return (hash);
    }
}

sub vcl_backend_response {
    if (beresp.status == 200 && bereq.url ~ "^/albums") {
        set beresp.ttl = 60s;
    }
    if (bereq.url ~ "^/albums") {
        if (beresp.status != 200) {
            set beresp.http.x-error = "Backend returned " + beresp.status;
            return (abandon);
        }
        if (beresp.ttl <= 0s ||
            beresp.http.Set-Cookie ||
            beresp.http.Vary == "*") {
            set beresp.ttl = 15s;
            return (deliver);
        }
    }
}

sub vcl_deliver {
    if (req.url ~ "^/albums") {
        set resp.http.Cache-Control = "public, max-age=3600"; 
    }
}