FROM cenoq/bbin:init
 
ADD api/core/app /bin/app
ENV GOROOT="/usr/local/go"
ENV PATH="/usr/local/go/bin:${PATH}"

ENTRYPOINT ["/bin/app"]