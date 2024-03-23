# Server Monitor

Server Monitor is a lightweight Go application designed to monitor different URLs and provide reports on their status. If a URL fails to respond correctly, the application sends a notification to the specified Slack channel, mentioning relevant users.

## Features

- **Last Checked Report**: Stay informed with up-to-date reports on server status.
- **Slack Notification**: Receive immediate notifications via Slack in case of server issues. (Optional)
- **User Mentioning**: Seamlessly mention specific users in Slack messages for targeted communication.
- **Custom Time Tickers**: Configure different time tickers for monitoring multiple servers according to your needs.
- **Secret Variables Management**: Securely manage sensitive information such as API keys and credentials using a `.env` file.
- **Custom Timeout Value**: Set a custom timeout value to control how long the application waits for server responses.

With Server Monitor, you can proactively monitor your servers, ensuring smooth operations and quick response to any issues that may arise.
