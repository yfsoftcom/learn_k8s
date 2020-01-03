FROM 10.12.3.74:8881/library/air_maven:1.0

ADD ./run.sh /run.sh
RUN chmod +x /run.sh