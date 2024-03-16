package main

import (
	"context"
	"testing"
)

const pingEvent = `{
  "zen": "Practicality beats purity.",
  "hook_id": 466966724,
  "hook": {
    "type": "Repository",
    "id": 466966724,
    "name": "web",
    "active": true,
    "events": [
      "*"
    ],
    "config": {
      "content_type": "json",
      "insecure_ssl": "0",
      "url": "https://api.emitly.dev/v1/webhook?listenerId=fn_d4ed02609e29c9826eaad0744cbf0e56&apikey=tsxKKZgpDx7CM3yQq3pUYvXH4yp3oCe7KxaZtmqi"
    },
    "updated_at": "2024-03-16T15:31:32Z",
    "created_at": "2024-03-16T15:31:32Z",
    "url": "https://api.github.com/repos/Quix44/hackathon-example/hooks/466966724",
    "test_url": "https://api.github.com/repos/Quix44/hackathon-example/hooks/466966724/test",
    "ping_url": "https://api.github.com/repos/Quix44/hackathon-example/hooks/466966724/pings",
    "deliveries_url": "https://api.github.com/repos/Quix44/hackathon-example/hooks/466966724/deliveries",
    "last_response": {
      "code": null,
      "status": "unused",
      "message": null
    }
  },
  "repository": {
    "id": 772923854,
    "node_id": "R_kgDOLhHhzg",
    "name": "hackathon-example",
    "full_name": "Quix44/hackathon-example",
    "private": false,
    "owner": {
      "login": "Quix44",
      "id": 96739828,
      "node_id": "U_kgDOBcQh9A",
      "avatar_url": "https://avatars.githubusercontent.com/u/96739828?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/Quix44",
      "html_url": "https://github.com/Quix44",
      "followers_url": "https://api.github.com/users/Quix44/followers",
      "following_url": "https://api.github.com/users/Quix44/following{/other_user}",
      "gists_url": "https://api.github.com/users/Quix44/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/Quix44/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/Quix44/subscriptions",
      "organizations_url": "https://api.github.com/users/Quix44/orgs",
      "repos_url": "https://api.github.com/users/Quix44/repos",
      "events_url": "https://api.github.com/users/Quix44/events{/privacy}",
      "received_events_url": "https://api.github.com/users/Quix44/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/Quix44/hackathon-example",
    "description": "Gitpaid example",
    "fork": false,
    "url": "https://api.github.com/repos/Quix44/hackathon-example",
    "forks_url": "https://api.github.com/repos/Quix44/hackathon-example/forks",
    "keys_url": "https://api.github.com/repos/Quix44/hackathon-example/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/Quix44/hackathon-example/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/Quix44/hackathon-example/teams",
    "hooks_url": "https://api.github.com/repos/Quix44/hackathon-example/hooks",
    "issue_events_url": "https://api.github.com/repos/Quix44/hackathon-example/issues/events{/number}",
    "events_url": "https://api.github.com/repos/Quix44/hackathon-example/events",
    "assignees_url": "https://api.github.com/repos/Quix44/hackathon-example/assignees{/user}",
    "branches_url": "https://api.github.com/repos/Quix44/hackathon-example/branches{/branch}",
    "tags_url": "https://api.github.com/repos/Quix44/hackathon-example/tags",
    "blobs_url": "https://api.github.com/repos/Quix44/hackathon-example/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/Quix44/hackathon-example/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/Quix44/hackathon-example/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/Quix44/hackathon-example/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/Quix44/hackathon-example/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/Quix44/hackathon-example/languages",
    "stargazers_url": "https://api.github.com/repos/Quix44/hackathon-example/stargazers",
    "contributors_url": "https://api.github.com/repos/Quix44/hackathon-example/contributors",
    "subscribers_url": "https://api.github.com/repos/Quix44/hackathon-example/subscribers",
    "subscription_url": "https://api.github.com/repos/Quix44/hackathon-example/subscription",
    "commits_url": "https://api.github.com/repos/Quix44/hackathon-example/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/Quix44/hackathon-example/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/Quix44/hackathon-example/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/Quix44/hackathon-example/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/Quix44/hackathon-example/contents/{+path}",
    "compare_url": "https://api.github.com/repos/Quix44/hackathon-example/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/Quix44/hackathon-example/merges",
    "archive_url": "https://api.github.com/repos/Quix44/hackathon-example/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/Quix44/hackathon-example/downloads",
    "issues_url": "https://api.github.com/repos/Quix44/hackathon-example/issues{/number}",
    "pulls_url": "https://api.github.com/repos/Quix44/hackathon-example/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/Quix44/hackathon-example/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/Quix44/hackathon-example/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/Quix44/hackathon-example/labels{/name}",
    "releases_url": "https://api.github.com/repos/Quix44/hackathon-example/releases{/id}",
    "deployments_url": "https://api.github.com/repos/Quix44/hackathon-example/deployments",
    "created_at": "2024-03-16T08:51:18Z",
    "updated_at": "2024-03-16T08:58:45Z",
    "pushed_at": "2024-03-16T13:51:56Z",
    "git_url": "git://github.com/Quix44/hackathon-example.git",
    "ssh_url": "git@github.com:Quix44/hackathon-example.git",
    "clone_url": "https://github.com/Quix44/hackathon-example.git",
    "svn_url": "https://github.com/Quix44/hackathon-example",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": "Python",
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "has_discussions": false,
    "forks_count": 1,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 3,
    "license": null,
    "allow_forking": true,
    "is_template": false,
    "web_commit_signoff_required": false,
    "topics": [

    ],
    "visibility": "public",
    "forks": 1,
    "open_issues": 3,
    "watchers": 0,
    "default_branch": "main"
  },
  "sender": {
    "login": "Quix44",
    "id": 96739828,
    "node_id": "U_kgDOBcQh9A",
    "avatar_url": "https://avatars.githubusercontent.com/u/96739828?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/Quix44",
    "html_url": "https://github.com/Quix44",
    "followers_url": "https://api.github.com/users/Quix44/followers",
    "following_url": "https://api.github.com/users/Quix44/following{/other_user}",
    "gists_url": "https://api.github.com/users/Quix44/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/Quix44/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/Quix44/subscriptions",
    "organizations_url": "https://api.github.com/users/Quix44/orgs",
    "repos_url": "https://api.github.com/users/Quix44/repos",
    "events_url": "https://api.github.com/users/Quix44/events{/privacy}",
    "received_events_url": "https://api.github.com/users/Quix44/received_events",
    "type": "User",
    "site_admin": false
  }
}`

const closedIssue = `{
  "action": "closed",
  "issue": {
    "url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6",
    "repository_url": "https://api.github.com/repos/danagain/stunning-computing-machine",
    "labels_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6/labels{/name}",
    "comments_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6/comments",
    "events_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6/events",
    "html_url": "https://github.com/danagain/stunning-computing-machine/issues/6",
    "id": 2189684328,
    "node_id": "I_kwDOLg9pH86Cg-5o",
    "number": 6,
    "title": "test label 2",
    "user": {
      "login": "danagain",
      "id": 18300720,
      "node_id": "MDQ6VXNlcjE4MzAwNzIw",
      "avatar_url": "https://avatars.githubusercontent.com/u/18300720?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/danagain",
      "html_url": "https://github.com/danagain",
      "followers_url": "https://api.github.com/users/danagain/followers",
      "following_url": "https://api.github.com/users/danagain/following{/other_user}",
      "gists_url": "https://api.github.com/users/danagain/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/danagain/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/danagain/subscriptions",
      "organizations_url": "https://api.github.com/users/danagain/orgs",
      "repos_url": "https://api.github.com/users/danagain/repos",
      "events_url": "https://api.github.com/users/danagain/events{/privacy}",
      "received_events_url": "https://api.github.com/users/danagain/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "id": 6698004172,
        "node_id": "LA_kwDOLg9pH88AAAABjztuzA",
        "url": "https://api.github.com/repos/danagain/stunning-computing-machine/labels/100%20usdc",
        "name": "100 usdc",
        "color": "1d76db",
        "default": false,
        "description": ""
      }
    ],
    "state": "closed",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2024-03-16T01:33:44Z",
    "updated_at": "2024-03-16T01:37:20Z",
    "closed_at": "2024-03-16T01:37:20Z",
    "author_association": "OWNER",
    "active_lock_reason": null,
    "body": null,
    "reactions": {
      "url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6/reactions",
      "total_count": 0,
      "+1": 0,
      "-1": 0,
      "laugh": 0,
      "hooray": 0,
      "confused": 0,
      "heart": 0,
      "rocket": 0,
      "eyes": 0
    },
    "timeline_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/6/timeline",
    "performed_via_github_app": null,
    "state_reason": "completed"
  },
  "repository": {
    "id": 772761887,
    "node_id": "R_kgDOLg9pHw",
    "name": "stunning-computing-machine",
    "full_name": "danagain/stunning-computing-machine",
    "private": false,
    "owner": {
      "login": "danagain",
      "id": 18300720,
      "node_id": "MDQ6VXNlcjE4MzAwNzIw",
      "avatar_url": "https://avatars.githubusercontent.com/u/18300720?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/danagain",
      "html_url": "https://github.com/danagain",
      "followers_url": "https://api.github.com/users/danagain/followers",
      "following_url": "https://api.github.com/users/danagain/following{/other_user}",
      "gists_url": "https://api.github.com/users/danagain/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/danagain/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/danagain/subscriptions",
      "organizations_url": "https://api.github.com/users/danagain/orgs",
      "repos_url": "https://api.github.com/users/danagain/repos",
      "events_url": "https://api.github.com/users/danagain/events{/privacy}",
      "received_events_url": "https://api.github.com/users/danagain/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/danagain/stunning-computing-machine",
    "description": null,
    "fork": false,
    "url": "https://api.github.com/repos/danagain/stunning-computing-machine",
    "forks_url": "https://api.github.com/repos/danagain/stunning-computing-machine/forks",
    "keys_url": "https://api.github.com/repos/danagain/stunning-computing-machine/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/danagain/stunning-computing-machine/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/danagain/stunning-computing-machine/teams",
    "hooks_url": "https://api.github.com/repos/danagain/stunning-computing-machine/hooks",
    "issue_events_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/events{/number}",
    "events_url": "https://api.github.com/repos/danagain/stunning-computing-machine/events",
    "assignees_url": "https://api.github.com/repos/danagain/stunning-computing-machine/assignees{/user}",
    "branches_url": "https://api.github.com/repos/danagain/stunning-computing-machine/branches{/branch}",
    "tags_url": "https://api.github.com/repos/danagain/stunning-computing-machine/tags",
    "blobs_url": "https://api.github.com/repos/danagain/stunning-computing-machine/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/danagain/stunning-computing-machine/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/danagain/stunning-computing-machine/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/danagain/stunning-computing-machine/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/danagain/stunning-computing-machine/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/danagain/stunning-computing-machine/languages",
    "stargazers_url": "https://api.github.com/repos/danagain/stunning-computing-machine/stargazers",
    "contributors_url": "https://api.github.com/repos/danagain/stunning-computing-machine/contributors",
    "subscribers_url": "https://api.github.com/repos/danagain/stunning-computing-machine/subscribers",
    "subscription_url": "https://api.github.com/repos/danagain/stunning-computing-machine/subscription",
    "commits_url": "https://api.github.com/repos/danagain/stunning-computing-machine/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/danagain/stunning-computing-machine/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/danagain/stunning-computing-machine/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/danagain/stunning-computing-machine/contents/{+path}",
    "compare_url": "https://api.github.com/repos/danagain/stunning-computing-machine/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/danagain/stunning-computing-machine/merges",
    "archive_url": "https://api.github.com/repos/danagain/stunning-computing-machine/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/danagain/stunning-computing-machine/downloads",
    "issues_url": "https://api.github.com/repos/danagain/stunning-computing-machine/issues{/number}",
    "pulls_url": "https://api.github.com/repos/danagain/stunning-computing-machine/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/danagain/stunning-computing-machine/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/danagain/stunning-computing-machine/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/danagain/stunning-computing-machine/labels{/name}",
    "releases_url": "https://api.github.com/repos/danagain/stunning-computing-machine/releases{/id}",
    "deployments_url": "https://api.github.com/repos/danagain/stunning-computing-machine/deployments",
    "created_at": "2024-03-15T21:02:27Z",
    "updated_at": "2024-03-15T21:02:27Z",
    "pushed_at": "2024-03-16T01:37:19Z",
    "git_url": "git://github.com/danagain/stunning-computing-machine.git",
    "ssh_url": "git@github.com:danagain/stunning-computing-machine.git",
    "clone_url": "https://github.com/danagain/stunning-computing-machine.git",
    "svn_url": "https://github.com/danagain/stunning-computing-machine",
    "homepage": null,
    "size": 3,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "has_discussions": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 1,
    "license": null,
    "allow_forking": true,
    "is_template": false,
    "web_commit_signoff_required": false,
    "topics": [

    ],
    "visibility": "public",
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "main"
  },
  "sender": {
    "login": "danagain",
    "id": 18300720,
    "node_id": "MDQ6VXNlcjE4MzAwNzIw",
    "avatar_url": "https://avatars.githubusercontent.com/u/18300720?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/danagain",
    "html_url": "https://github.com/danagain",
    "followers_url": "https://api.github.com/users/danagain/followers",
    "following_url": "https://api.github.com/users/danagain/following{/other_user}",
    "gists_url": "https://api.github.com/users/danagain/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/danagain/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/danagain/subscriptions",
    "organizations_url": "https://api.github.com/users/danagain/orgs",
    "repos_url": "https://api.github.com/users/danagain/repos",
    "events_url": "https://api.github.com/users/danagain/events{/privacy}",
    "received_events_url": "https://api.github.com/users/danagain/received_events",
    "type": "User",
    "site_admin": false
  }
}`

func TestPingEvent(t *testing.T) {
	ctx := context.Background()
	req := LambdaEvent{
		Event: pingEvent, // Use the test payload you provided
		Headers: map[string]string{
			"X-GitHub-Event": "ping", // Set the event type to "star"
		},
	}

	// Call your function
	resp, err := HandleRequest(ctx, req)

	// Check for errors
	if err != nil {
		t.Errorf("HandleRequest() error = %v", err)
		return
	}

	// Perform your test: check if the response status code is 200
	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusCode=200, got %v", resp.StatusCode)
	}
}

func TestClosedIssue(t *testing.T) {
	ctx := context.Background()
	req := LambdaEvent{
		Event: closedIssue, // Use the test payload you provided
		Headers: map[string]string{
			"X-GitHub-Event": "issues", // Set the event type to "star"
		},
	}

	// Call your function
	resp, err := HandleRequest(ctx, req)

	// Check for errors
	if err != nil {
		t.Errorf("HandleRequest() error = %v", err)
		return
	}

	// Perform your test: check if the response status code is 200
	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusCode=200, got %v", resp.StatusCode)
	}

}
