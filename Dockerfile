FROM golang:1.13.4 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN make compile

# ENV HOST=0.0.0.0 PORT=3000
# EXPOSE ${PORT}

CMD [ "make", "run" ]