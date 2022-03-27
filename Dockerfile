FROM golang:latest

RUN git clone https://github.com/alakbarz/alak.bar.git
RUN cd alak.bar && go run *.go

EXPOSE 4000