FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/helloworld-k8s-operator \
    USER_UID=1001 \
    USER_NAME=helloworld-k8s-operator

# install operator binary
COPY build/_output/bin/helloworld-k8s-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
