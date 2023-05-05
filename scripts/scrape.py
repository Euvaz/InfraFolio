import sqlite3

# Initilize Connection and Cursor
connection = sqlite3.connect("../locked.sqlite")
cursor = connection.cursor()

# Read from github table
response = cursor.execute("SELECT * FROM github")
print(response.fetchmany())

# Read from github_repositories table
response = cursor.execute("SELECT * FROM github_repositories")
print(response.fetchmany())

# Read from github_stars table
response = cursor.execute("SELECT * FROM github_stars")
print(response.fetchmany())

# Close Connection
connection.close()
