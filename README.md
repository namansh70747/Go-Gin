# Go Gin Project Structure Explained

Below is a simple explanation of what each folder in this project is for. If you're new, read this to understand where to put your code!

```
cmd/         # The starting point of your app. main.go lives here. Run your app from here.
config/      # Handles settings and secrets (like database passwords). Loads environment variables.
controllers/ # Functions that handle requests (what happens when someone visits a URL).
models/      # Defines your data shapes (like Note, User, Task). Used for database too.
routes/      # Connects URLs to controllers. Says which function runs for each URL.
middleware/  # Code that runs before/after your main logic (like checking if user is logged in).
services/    # Special helpers for things like sending emails or uploading files.
repository/  # Talks to the database. Handles saving, updating, deleting data.
utils/       # Helper functions you use in many places (like password hashing).
static/      # For images, uploads, or other files you want to serve as-is.
templates/   # HTML templates if you want to send web pages (optional for APIs).
.env         # Stores secrets and settings (like passwords). Never share this file!
README.md    # This file! Explains your project and folder structure.
```

Feel free to ask questions or add more explanations as you learn!
