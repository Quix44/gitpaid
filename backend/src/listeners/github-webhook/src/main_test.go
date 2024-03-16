package main

import (
	"context"
	"testing"
)

const pingEvent = `{
  "zen": "Avoid administrative distraction.",
  "hook_id": 466808924,
  "hook": {
    "type": "Repository",
    "id": 466808924,
    "name": "web",
    "active": true,
    "events": [
      "*"
    ],
    "config": {
      "content_type": "json",
      "insecure_ssl": "0",
      "url": "https://api.emitly.dev/v1/webhook?listenerId=fn_3361a5ade3850b114c2ac90b243b643e&apikey=tsxKKZgpDx7CM3yQq3pUYvXH4yp3oCe7KxaZtmqi"
    },
    "updated_at": "2024-03-15T22:23:50Z",
    "created_at": "2024-03-15T22:23:50Z",
    "url": "https://api.github.com/repos/danagain/stunning-computing-machine/hooks/466808924",
    "test_url": "https://api.github.com/repos/danagain/stunning-computing-machine/hooks/466808924/test",
    "ping_url": "https://api.github.com/repos/danagain/stunning-computing-machine/hooks/466808924/pings",
    "deliveries_url": "https://api.github.com/repos/danagain/stunning-computing-machine/hooks/466808924/deliveries",
    "last_response": {
      "code": null,
      "status": "unused",
      "message": null
    }
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
    "pushed_at": "2024-03-15T21:02:27Z",
    "git_url": "git://github.com/danagain/stunning-computing-machine.git",
    "ssh_url": "git@github.com:danagain/stunning-computing-machine.git",
    "clone_url": "https://github.com/danagain/stunning-computing-machine.git",
    "svn_url": "https://github.com/danagain/stunning-computing-machine",
    "homepage": null,
    "size": 0,
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
    "open_issues_count": 0,
    "license": null,
    "allow_forking": true,
    "is_template": false,
    "web_commit_signoff_required": false,
    "topics": [

    ],
    "visibility": "public",
    "forks": 0,
    "open_issues": 0,
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
