FROM mcr.microsoft.com/mssql/server:2017-CU8-ubuntu
COPY . .
ENV ACCEPT_EULA Y
ENV SA_PASSWORD yourStrong(!)Password
RUN ./wait-for-it.sh 127.0.0.1:1433 --timeout=25 -- echo "mssql is up"
RUN ./db-init.sh & /opt/mssql/bin/sqlservr
