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

# Get user information
gh_user_payload = http.request(GITHUB_USER, method="GET")[1]
gh_user_content = json.loads(gh_user_payload.decode())

cursor.execute("""INSERT INTO github (name)
               VALUES (?)""", (gh_user_content["login"],))

# Commit changes
connection.commit()

# Get repository information
gh_repo_payload = http.request(GITHUB_REPO_URL, method="GET")[1]
gh_repo_content = json.loads(gh_repo_payload.decode())

gh_repo_owners = []
gh_repo_names = []
gh_repo_preview = []
gh_repo_stargazers_count = []
for repo in gh_repo_content:
    gh_repo_owners.append(repo["owner"]["login"])
    gh_repo_names.append(repo["name"])
    response, content = http.request(f"https://github.com/{repo['owner']['login']}/{repo['name']}")
    start_index = content.find(b'<meta property="og:image" content="')
    if start_index != -1:
        end_index = content.find(b'"', start_index + len('<meta property="og:image" content="'))
        gh_repo_preview.append(content[start_index + len('<meta property="og:image" content="'):end_index].decode('utf-8'))
    gh_repo_stargazers_count.append(repo["stargazers_count"])

# Get fresh list of repositories
gh_repo_fresh = list(map(lambda owner, name, preview, star_count:(owner,name,preview,star_count), gh_repo_owners, gh_repo_names, gh_repo_preview, gh_repo_stargazers_count))

# Get current list of repositories
cursor.execute("SELECT owner, name FROM github_repositories")
gh_repo_current = set(cursor.fetchall())

# Get and delete differing entries
gh_repo_diff = gh_repo_current - set(gh_repo_fresh)
cursor.executemany("DELETE FROM github_repositories where owner=? AND name=?", (gh_repo_diff))

# Insert new entries
cursor.executemany("""INSERT OR IGNORE INTO github_repositories (owner, name, preview, stargazers_count)
                   VALUES (?, ?, ?, ?)""", [(repo[0], repo[1], repo[2], repo[3]) for repo in gh_repo_fresh])

# Commit changes
connection.commit()

# Get repository star information
#gh_repo_star_payload = http.request(GITHUB_REPO_URL)

# Close Connection
connection.close()
