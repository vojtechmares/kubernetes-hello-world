FROM scratch

ENTRYPOINT [ "/hello" ]
EXPOSE 8080

COPY hello /
