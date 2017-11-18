FROM scratch
COPY mqconnect /mqconnect
EXPOSE 80
ENTRYPOINT [ "/mqconnect" ]
