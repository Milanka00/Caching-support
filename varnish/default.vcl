vcl 4.0;

backend default {
    .host = "backend";   //actual backend ip
    .port = "8081";
}

sub vcl_recv {
    # Perform hash lookup for any request
    return (hash);
}

sub vcl_backend_response {
    # Check if there is a Cache-Control header in the backend response
    if (beresp.http.Cache-Control) {
        # If Cache-Control header is present, we don't override it
        return (deliver);
    }
    
    # If there's no Cache-Control header, set default TTL
    set beresp.ttl = 60s;
}

sub vcl_deliver {
    # Optionally, set default cache control headers for responses
    if (!resp.http.Cache-Control) {
        set resp.http.Cache-Control = "public, max-age=3600";
    }
}
