FROM golang
WORKDIR $HOME/terraform-inventory
COPY . .
RUN go install
ENTRYPOINT ["terraform-inventory"]
