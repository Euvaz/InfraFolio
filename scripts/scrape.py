import sqlite3
import httplib2
import json

# Initilize Connection and Cursor
connection = sqlite3.connect("locked.sqlite")
cursor = connection.cursor()

# Initialize HTTP interface
http = httplib2.Http()

GITHUB_USER = "https://api.github.com/users/Euvaz"
GITHUB_REPO_URL = GITHUB_USER + "/repos"

# Get repository information
gh_repo_payload = http.request(GITHUB_REPO_URL, method="GET")[1]
gh_repo_content = json.loads(gh_repo_payload.decode())

gh_repo_owners = []
gh_repo_names = []
for repo in gh_repo_content:
    gh_repo_owners.append(repo["owner"]["login"])
    gh_repo_names.append(repo["name"])

gh_repo_complete = list(map(lambda x, y:(x,y), gh_repo_owners, gh_repo_names))
for repo in gh_repo_complete:
    cursor.execute(f"""INSERT INTO github_repositories (owner, name)
                   VALUES ("{repo[0]}", "{repo[1]}")""")

# Commit changes
connection.commit()

# Close Connection
connection.close()
