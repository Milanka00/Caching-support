vcl 4.1;

import directors;

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
}
