**How to set up:**

1. Create file .env (.env.example) and fill the value
2. Create table in database with migration
3. Open postman to hit API
    - API doc using swagger (open folder docs -> swagger.yaml)

**Migration**

Command to run:

        goose -dir migration/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" up

For rollback:

        goose -dir migration/ mysql "<username>:<password>@tcp(<url>:<port>)/<db-name>?parseTime=true" down


