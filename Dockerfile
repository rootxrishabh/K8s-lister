FROM ubuntu

WORKDIR /

COPY ./linuxmain ./linuxmain

ENTRYPOINT [ "./linuxmain" ]