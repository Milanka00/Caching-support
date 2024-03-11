vcl 4.1;

import directors;
import std;

backend backend1 {
    .host = "backend";
    .port = "8081";
}

backend backend2 {
    .host = "second_backend";
    .port = "8082";
}

sub vcl_init {
    new vdir = directors.round_robin();
    vdir.add_backend(backend1);
    vdir.add_backend(backend2);
     
}

sub vcl_recv {
    set req.backend_hint = vdir.backend();
    std.log("Host header received: " + req.http.Host);
    
    
}

sub vcl_backend_response {
# Don't cache 404 responses
if ( beresp.status == 404 ) {
set beresp.uncacheable = true;
return (deliver);
}
}

sub vcl_deliver {
    if (obj.hits > 0) {
        set resp.http.X-Cache-Host = req.http.Host;
        set resp.http.X-Cache-Info = "Cached under host: " + req.http.Host + "; Request URI: " + req.url;
    }
}
