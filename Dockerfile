FROM golang:1.18.4-alpine

COPY rusprofile rusprofile

CMD ["./rusprofile"]
