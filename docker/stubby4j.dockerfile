FROM azagniotov/stubby4j:7.5.1-jre11

WORKDIR /home

COPY test .

RUN ls -lrt /

RUN ls -lrt /home

RUN ls -lrt /home

CMD  java -jar /home/stubby4j.jar -da -ds -d /home/data/main.yaml -l 0.0.0.0 -s 8882