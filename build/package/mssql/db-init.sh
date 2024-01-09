sleep 5s
echo "running set up script"
/opt/mssql-tools/bin/sqlcmd -S localhost -U $SA_USER -P $SA_PASSWORD -d master -i ./db-init.sql
echo "migrated"
pkill sqlservr 