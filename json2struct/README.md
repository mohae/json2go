json2struct
===========

json2struct is a CLI app that generates Go struct definitions from JSON.  Any objects in the source JSON will result in their own struct.  Any values that are null will have their type be `interface{}`; the type cannot be determined on null values.

By default, json2struct will read the JSON from `stdin` and write it to `stdout`.  Both a file source and file destination can be specified.  When the output destination is a file, the JSON used to generate the struct definition can also be written out to the same destination; the filename will be the same as the struct definition(s) except it will have the `.json` extension, instead fo the `.go` extension.

The default package name for the generated Go source is `main`, this can be overridden using the `-pkg` or `-p` flags.

The generated source can include the import statement for `encoding/json` by using the `-import` or `-m` flag.

## Flags

```
Flag | Short | Default | Description
---- | ----- | ------- | -----------
name | n | | the name of the struct
pkg | p | main | the name of the package
input | i | stdin | the input source
output | o | stdout | the output destination
writejson | w | false | write the source JSON to file; only applicable if output is not stdout
import | i | false | add import statement for 'encoding/json'

```
### Example

    curl -s https://api.github.com/repos/mohae/json2struct | json2struct -o github.go -w -m -n Github

Results in `github.go`:

```
package main

import (
	"encoding/json"
)

type Github struct {
	ArchiveUrl string `json:"archive_url"`
	AssigneesUrl string `json:"assignees_url"`
	BlobsUrl string `json:"blobs_url"`
	BranchesUrl string `json:"branches_url"`
	CloneUrl string `json:"clone_url"`
	CollaboratorsUrl string `json:"collaborators_url"`
	CommentsUrl string `json:"comments_url"`
	CommitsUrl string `json:"commits_url"`
	CompareUrl string `json:"compare_url"`
	ContentsUrl string `json:"contents_url"`
	ContributorsUrl string `json:"contributors_url"`
	CreatedAt string `json:"created_at"`
	DefaultBranch string `json:"default_branch"`
	Description string `json:"description"`
	DownloadsUrl string `json:"downloads_url"`
	EventsUrl string `json:"events_url"`
	Fork bool `json:"fork"`
	Forks int `json:"forks"`
	ForksCount int `json:"forks_count"`
	ForksUrl string `json:"forks_url"`
	FullName string `json:"full_name"`
	GitCommitsUrl string `json:"git_commits_url"`
	GitRefsUrl string `json:"git_refs_url"`
	GitTagsUrl string `json:"git_tags_url"`
	GitUrl string `json:"git_url"`
	HasDownloads bool `json:"has_downloads"`
	HasIssues bool `json:"has_issues"`
	HasPages bool `json:"has_pages"`
	HasWiki bool `json:"has_wiki"`
	Homepage interface{} `json:"homepage"`
	HooksUrl string `json:"hooks_url"`
	HtmlUrl string `json:"html_url"`
	Id int `json:"id"`
	IssueCommentUrl string `json:"issue_comment_url"`
	IssueEventsUrl string `json:"issue_events_url"`
	IssuesUrl string `json:"issues_url"`
	KeysUrl string `json:"keys_url"`
	LabelsUrl string `json:"labels_url"`
	Language string `json:"language"`
	LanguagesUrl string `json:"languages_url"`
	MergesUrl string `json:"merges_url"`
	MilestonesUrl string `json:"milestones_url"`
	MirrorUrl interface{} `json:"mirror_url"`
	Name string `json:"name"`
	NetworkCount int `json:"network_count"`
	NotificationsUrl string `json:"notifications_url"`
	OpenIssues int `json:"open_issues"`
	OpenIssuesCount int `json:"open_issues_count"`
	Owner `json:"owner"`
	Private bool `json:"private"`
	PullsUrl string `json:"pulls_url"`
	PushedAt string `json:"pushed_at"`
	ReleasesUrl string `json:"releases_url"`
	Size int `json:"size"`
	SshUrl string `json:"ssh_url"`
	StargazersCount int `json:"stargazers_count"`
	StargazersUrl string `json:"stargazers_url"`
	StatusesUrl string `json:"statuses_url"`
	SubscribersCount int `json:"subscribers_count"`
	SubscribersUrl string `json:"subscribers_url"`
	SubscriptionUrl string `json:"subscription_url"`
	SvnUrl string `json:"svn_url"`
	TagsUrl string `json:"tags_url"`
	TeamsUrl string `json:"teams_url"`
	TreesUrl string `json:"trees_url"`
	UpdatedAt string `json:"updated_at"`
	Url string `json:"url"`
	Watchers int `json:"watchers"`
	WatchersCount int `json:"watchers_count"`
}
type Owner struct {
	AvatarUrl string `json:"avatar_url"`
	EventsUrl string `json:"events_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	GravatarId string `json:"gravatar_id"`
	HtmlUrl string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	OrganizationsUrl string `json:"organizations_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	ReposUrl string `json:"repos_url"`
	SiteAdmin bool `json:"site_admin"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	Type string `json:"type"`
	Url string `json:"url"`
}

```

and github.json:

```
{
  "id": 47099645,
  "name": "json2struct",
  "full_name": "mohae/json2struct",
  "owner": {
    "login": "mohae",
    "id": 2699987,
    "avatar_url": "https://avatars.githubusercontent.com/u/2699987?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/mohae",
    "html_url": "https://github.com/mohae",
    "followers_url": "https://api.github.com/users/mohae/followers",
    "following_url": "https://api.github.com/users/mohae/following{/other_user}",
    "gists_url": "https://api.github.com/users/mohae/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/mohae/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/mohae/subscriptions",
    "organizations_url": "https://api.github.com/users/mohae/orgs",
    "repos_url": "https://api.github.com/users/mohae/repos",
    "events_url": "https://api.github.com/users/mohae/events{/privacy}",
    "received_events_url": "https://api.github.com/users/mohae/received_events",
    "type": "User",
    "site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/mohae/json2struct",
  "description": "generate Go struct definitions from JSON",
  "fork": false,
  "url": "https://api.github.com/repos/mohae/json2struct",
  "forks_url": "https://api.github.com/repos/mohae/json2struct/forks",
  "keys_url": "https://api.github.com/repos/mohae/json2struct/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/mohae/json2struct/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/mohae/json2struct/teams",
  "hooks_url": "https://api.github.com/repos/mohae/json2struct/hooks",
  "issue_events_url": "https://api.github.com/repos/mohae/json2struct/issues/events{/number}",
  "events_url": "https://api.github.com/repos/mohae/json2struct/events",
  "assignees_url": "https://api.github.com/repos/mohae/json2struct/assignees{/user}",
  "branches_url": "https://api.github.com/repos/mohae/json2struct/branches{/branch}",
  "tags_url": "https://api.github.com/repos/mohae/json2struct/tags",
  "blobs_url": "https://api.github.com/repos/mohae/json2struct/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/mohae/json2struct/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/mohae/json2struct/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/mohae/json2struct/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/mohae/json2struct/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/mohae/json2struct/languages",
  "stargazers_url": "https://api.github.com/repos/mohae/json2struct/stargazers",
  "contributors_url": "https://api.github.com/repos/mohae/json2struct/contributors",
  "subscribers_url": "https://api.github.com/repos/mohae/json2struct/subscribers",
  "subscription_url": "https://api.github.com/repos/mohae/json2struct/subscription",
  "commits_url": "https://api.github.com/repos/mohae/json2struct/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/mohae/json2struct/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/mohae/json2struct/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/mohae/json2struct/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/mohae/json2struct/contents/{+path}",
  "compare_url": "https://api.github.com/repos/mohae/json2struct/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/mohae/json2struct/merges",
  "archive_url": "https://api.github.com/repos/mohae/json2struct/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/mohae/json2struct/downloads",
  "issues_url": "https://api.github.com/repos/mohae/json2struct/issues{/number}",
  "pulls_url": "https://api.github.com/repos/mohae/json2struct/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/mohae/json2struct/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/mohae/json2struct/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/mohae/json2struct/labels{/name}",
  "releases_url": "https://api.github.com/repos/mohae/json2struct/releases{/id}",
  "created_at": "2015-11-30T06:31:18Z",
  "updated_at": "2015-11-30T06:32:11Z",
  "pushed_at": "2015-11-30T06:32:10Z",
  "git_url": "git://github.com/mohae/json2struct.git",
  "ssh_url": "git@github.com:mohae/json2struct.git",
  "clone_url": "https://github.com/mohae/json2struct.git",
  "svn_url": "https://github.com/mohae/json2struct",
  "homepage": null,
  "size": 13,
  "stargazers_count": 0,
  "watchers_count": 0,
  "language": "Go",
  "has_issues": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 0,
  "mirror_url": null,
  "open_issues_count": 0,
  "forks": 0,
  "open_issues": 0,
  "watchers": 0,
  "default_branch": "master",
  "network_count": 0,
  "subscribers_count": 1
}

```
