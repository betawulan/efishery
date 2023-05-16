README.md


**Migration**

Command to run:

        goose -dir migration/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" up

For rollback:

        goose -dir migration/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" down
