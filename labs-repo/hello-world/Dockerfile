FROM quay.io/projectquay/golang:1.17

ADD src/hello-world.go hello-world.go

ENV MESSAGE "Welcome! You can change this message by replacing the MESSAGE environment variable."
ENV HOME /go

RUN chgrp -R 0 /go && chmod -R g+rwX /go

EXPOSE 8080

LABEL io.openshift.expose-services 8080/http

USER 1001

CMD go run hello-world.go
# CMD pwd












