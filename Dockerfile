FROM ubuntu

WORKDIR /

COPY ./main ./main

ENTRYPOINT [ "./main" ]