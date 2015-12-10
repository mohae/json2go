json2go
=======

json2go is a CLI application that generates Go type definitions from JSON.  The type may be one of the following:

    * struct
    * map[string]T
    * map[string][]T

By default, a struct will be generated.

Any objects in the source JSON will result in their own struct.  Any values that are null will have their type be `interface{}`; the type cannot be determined on null values.

If the source JSON is an array of objects, the first element in the array will be used to generate the definition(s).  Any objects within the JSON will result in additional embedded struct types.

Keys with underscores, `_`, are converted to MixedCase.  Keys starting with characters that are invalid for Go variable names have those characters discarded, unless they are a number, `0-9`, which are converted to their word equivalents. All fields are exported and the JSON field tag for the field is generated using the original JSON key value.

By default, json2go will read the JSON from `stdin` and write it to `stdout`.  Both a source file and destination file can be specified.  When the output destination is a file, the JSON used to generate the struct definition can also be written to a file.  The filename will be the same as the Go output file except it will have the `.json` extension.

The default package name for the generated Go source is `main`, this can be overridden using either the `-pkg` or `-p` flags.

The generated source can include the import statement for `encoding/json` by using either the `-addimport` or `-a` flag.

## Flags

    Flag | Short | Default | Description  
    :---|:---|:---|:---  
    -name | -n |   | The name of the type: required.
    -input | -i | stdin | The JSON input source.
    -output | -o | stdout | The Go srouce code output destination.
    -writejson | -w | false | Write the source JSON to file; only valid when the output is a file.
    -pkg | -p | main | The name of the package.
    -addimport | -a | false | Add import statement for 'encoding/json'.
    -maptype | -m | false | Interpret the JSON as a map type instead of a struct type.
    -structname | -s | Struct | The name of the struct; only used in conjunction with -maptype.
    -help | -h | false | Print the help text; 'help' is also valid.

##  Usage:

Compile:

    go get github.com/mohae/json2go
    cd github.com/mohae/json2go/cmd/json2go
    go build -o $GOPATH/bin/json2go

Verify:
  
    json2go -h  
    

## Example 1

This example gets the JSON from a remote source and pipes it into `json2go`; generating both the Go source code file and a file with the JSON used to generate the struct definitions.

### Command

    curl -s https://api.github.com/repos/mohae/json2go | json2go -o github.go -w -a -n repo

#### Generated `github.go`

```
package main

import (
	"encoding/json"
)

type Repo struct {
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CloneURL         string      `json:"clone_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	CreatedAt        string      `json:"created_at"`
	DefaultBranch    string      `json:"default_branch"`
	Description      string      `json:"description"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	Fork             bool        `json:"fork"`
	Forks            int         `json:"forks"`
	ForksCount       int         `json:"forks_count"`
	ForksURL         string      `json:"forks_url"`
	FullName         string      `json:"full_name"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	HasDownloads     bool        `json:"has_downloads"`
	HasIssues        bool        `json:"has_issues"`
	HasPages         bool        `json:"has_pages"`
	HasWiki          bool        `json:"has_wiki"`
	Homepage         string      `json:"homepage"`
	HooksURL         string      `json:"hooks_url"`
	HTMLURL          string      `json:"html_url"`
	ID               int         `json:"id"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	Language         string      `json:"language"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	MirrorURL        interface{} `json:"mirror_url"`
	Name             string      `json:"name"`
	NetworkCount     int         `json:"network_count"`
	NotificationsURL string      `json:"notifications_url"`
	OpenIssues       int         `json:"open_issues"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	Owner            `json:"owner"`
	Private          bool   `json:"private"`
	PullsURL         string `json:"pulls_url"`
	PushedAt         string `json:"pushed_at"`
	ReleasesURL      string `json:"releases_url"`
	Size             int    `json:"size"`
	SSHURL           string `json:"ssh_url"`
	StargazersCount  int    `json:"stargazers_count"`
	StargazersURL    string `json:"stargazers_url"`
	StatusesURL      string `json:"statuses_url"`
	SubscribersCount int    `json:"subscribers_count"`
	SubscribersURL   string `json:"subscribers_url"`
	SubscriptionURL  string `json:"subscription_url"`
	SvnURL           string `json:"svn_url"`
	TagsURL          string `json:"tags_url"`
	TeamsURL         string `json:"teams_url"`
	TreesURL         string `json:"trees_url"`
	UpdatedAt        string `json:"updated_at"`
	URL              string `json:"url"`
	Watchers         int    `json:"watchers"`
	WatchersCount    int    `json:"watchers_count"`
}

type Owner struct {
	AvatarURL         string `json:"avatar_url"`
	EventsURL         string `json:"events_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	HTMLURL           string `json:"html_url"`
	ID                int    `json:"id"`
	Login             string `json:"login"`
	OrganizationsURL  string `json:"organizations_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	URL               string `json:"url"`
}
```

#### Source JSON written to `github.json`

```
{
  "id": 47099645,
  "name": "json2go",
  "full_name": "mohae/json2go",
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
  "html_url": "https://github.com/mohae/json2go",
  "description": "Generate Go struct definitions from JSON",
  "fork": false,
  "url": "https://api.github.com/repos/mohae/json2go",
  "forks_url": "https://api.github.com/repos/mohae/json2go/forks",
  "keys_url": "https://api.github.com/repos/mohae/json2go/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/mohae/json2go/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/mohae/json2go/teams",
  "hooks_url": "https://api.github.com/repos/mohae/json2go/hooks",
  "issue_events_url": "https://api.github.com/repos/mohae/json2go/issues/events{/number}",
  "events_url": "https://api.github.com/repos/mohae/json2go/events",
  "assignees_url": "https://api.github.com/repos/mohae/json2go/assignees{/user}",
  "branches_url": "https://api.github.com/repos/mohae/json2go/branches{/branch}",
  "tags_url": "https://api.github.com/repos/mohae/json2go/tags",
  "blobs_url": "https://api.github.com/repos/mohae/json2go/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/mohae/json2go/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/mohae/json2go/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/mohae/json2go/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/mohae/json2go/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/mohae/json2go/languages",
  "stargazers_url": "https://api.github.com/repos/mohae/json2go/stargazers",
  "contributors_url": "https://api.github.com/repos/mohae/json2go/contributors",
  "subscribers_url": "https://api.github.com/repos/mohae/json2go/subscribers",
  "subscription_url": "https://api.github.com/repos/mohae/json2go/subscription",
  "commits_url": "https://api.github.com/repos/mohae/json2go/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/mohae/json2go/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/mohae/json2go/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/mohae/json2go/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/mohae/json2go/contents/{+path}",
  "compare_url": "https://api.github.com/repos/mohae/json2go/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/mohae/json2go/merges",
  "archive_url": "https://api.github.com/repos/mohae/json2go/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/mohae/json2go/downloads",
  "issues_url": "https://api.github.com/repos/mohae/json2go/issues{/number}",
  "pulls_url": "https://api.github.com/repos/mohae/json2go/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/mohae/json2go/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/mohae/json2go/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/mohae/json2go/labels{/name}",
  "releases_url": "https://api.github.com/repos/mohae/json2go/releases{/id}",
  "created_at": "2015-11-30T06:31:18Z",
  "updated_at": "2015-12-09T21:19:26Z",
  "pushed_at": "2015-12-09T21:18:44Z",
  "git_url": "git://github.com/mohae/json2go.git",
  "ssh_url": "git@github.com:mohae/json2go.git",
  "clone_url": "https://github.com/mohae/json2go.git",
  "svn_url": "https://github.com/mohae/json2go",
  "homepage": "",
  "size": 1198,
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
### Example 2:

This example results in a map[string]T

### Command:

    json2struct -i -m hockey.json -o hockey.go -n team -s player

#### hockey.json

```
{
  "Blackhawks": [
    {
      "name": "Tony Esposito",
      "number": 35,
      "position": "Goal Tender"
    },
    {
      "name": "Stan Mikita",
      "number": 21,
      "position": "Center"
    }
  ]
}
```

#### Generated hockey.go  

```
package main

type Team map[string][]Player

type Player struct {
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Position string `json:"position"`
}
```

## Example 3:

This example uses json of much greater complexity in a local file.

### Command:

    json2struct -i weather.json -o weather.go -n weather

#### Generated weather.go
The generated Go source code for `weather.json`:

```
package main

type Weather struct {
	HourlyForecasts []HourlyForecast `json:"hourly_forecast"`
	Response        `json:"response"`
}

type HourlyForecast struct {
	FCTTIME   `json:"FCTTIME"`
	Condition string `json:"condition"`
	Dewpoint  `json:"dewpoint"`
	Fctcode   string `json:"fctcode"`
	Feelslike `json:"feelslike"`
	Heatindex `json:"heatindex"`
	Humidity  string `json:"humidity"`
	Icon      string `json:"icon"`
	IconURL   string `json:"icon_url"`
	Mslp      `json:"mslp"`
	Pop       string `json:"pop"`
	Qpf       `json:"qpf"`
	Sky       string `json:"sky"`
	Snow      `json:"snow"`
	Temp      `json:"temp"`
	Uvi       string `json:"uvi"`
	Wdir      `json:"wdir"`
	Windchill `json:"windchill"`
	Wspd      `json:"wspd"`
	Wx        string `json:"wx"`
}

type Response struct {
	Features       `json:"features"`
	TermsofService string `json:"termsofService"`
	Version        string `json:"version"`
}

type FCTTIME struct {
	UTCDATE                string `json:"UTCDATE"`
	Age                    string `json:"age"`
	Ampm                   string `json:"ampm"`
	Civil                  string `json:"civil"`
	Epoch                  string `json:"epoch"`
	Hour                   string `json:"hour"`
	HourPadded             string `json:"hour_padded"`
	Isdst                  string `json:"isdst"`
	Mday                   string `json:"mday"`
	MdayPadded             string `json:"mday_padded"`
	Min                    string `json:"min"`
	MinUnpadded            string `json:"min_unpadded"`
	Mon                    string `json:"mon"`
	MonAbbrev              string `json:"mon_abbrev"`
	MonPadded              string `json:"mon_padded"`
	MonthName              string `json:"month_name"`
	MonthNameAbbrev        string `json:"month_name_abbrev"`
	Pretty                 string `json:"pretty"`
	Sec                    string `json:"sec"`
	Tz                     string `json:"tz"`
	WeekdayName            string `json:"weekday_name"`
	WeekdayNameAbbrev      string `json:"weekday_name_abbrev"`
	WeekdayNameNight       string `json:"weekday_name_night"`
	WeekdayNameNightUnlang string `json:"weekday_name_night_unlang"`
	WeekdayNameUnlang      string `json:"weekday_name_unlang"`
	Yday                   string `json:"yday"`
	Year                   string `json:"year"`
}

type Dewpoint struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Feelslike struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Heatindex struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Mslp struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Qpf struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Snow struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Temp struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Wdir struct {
	Degrees string `json:"degrees"`
	Dir     string `json:"dir"`
}

type Windchill struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Wspd struct {
	English string `json:"english"`
	Metric  string `json:"metric"`
}

type Features struct {
	Hourly int `json:"hourly"`
}
```

#### Contents of weather.json
The source JSON used to generate the Go struct definitions.

 ```
    {  
   "response":{  
      "version":"0.1",
      "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
      "features":{  
         "hourly":1
      }
   },
   "hourly_forecast":[  
      {  
         "FCTTIME":{  
            "hour":"17",
            "hour_padded":"17",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448920800",
            "pretty":"5:00 PM EST on November 30, 2015",
            "civil":"5:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"23",
            "metric":"-5"
         },
         "condition":"Mostly Cloudy",
         "icon":"mostlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_mostlycloudy.gif",
         "fctcode":"3",
         "sky":"71",
         "wspd":{  
            "english":"4",
            "metric":"6"
         },
         "wdir":{  
            "dir":"SSE",
            "degrees":"158"
         },
         "wx":"Mostly Cloudy",
         "uvi":"0",
         "humidity":"59",
         "windchill":{  
            "english":"32",
            "metric":"0"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"32",
            "metric":"0"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"0",
         "mslp":{  
            "english":"30.42",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"18",
            "hour_padded":"18",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448924400",
            "pretty":"6:00 PM EST on November 30, 2015",
            "civil":"6:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"35",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"24",
            "metric":"-4"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"59",
         "wspd":{  
            "english":"5",
            "metric":"8"
         },
         "wdir":{  
            "dir":"S",
            "degrees":"172"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"62",
         "windchill":{  
            "english":"31",
            "metric":"-0"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"31",
            "metric":"-0"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"0",
         "mslp":{  
            "english":"30.43",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"19",
            "hour_padded":"19",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448928000",
            "pretty":"7:00 PM EST on November 30, 2015",
            "civil":"7:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"25",
            "metric":"-4"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"49",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"S",
            "degrees":"184"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"63",
         "windchill":{  
            "english":"30",
            "metric":"-1"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"30",
            "metric":"-1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"0",
         "mslp":{  
            "english":"30.42",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"20",
            "hour_padded":"20",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448931600",
            "pretty":"8:00 PM EST on November 30, 2015",
            "civil":"8:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"25",
            "metric":"-4"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"32",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"193"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"65",
         "windchill":{  
            "english":"30",
            "metric":"-1"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"30",
            "metric":"-1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"0",
         "mslp":{  
            "english":"30.42",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"21",
            "hour_padded":"21",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448935200",
            "pretty":"9:00 PM EST on November 30, 2015",
            "civil":"9:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"26",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"33",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"203"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"66",
         "windchill":{  
            "english":"31",
            "metric":"-1"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"31",
            "metric":"-1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"1",
         "mslp":{  
            "english":"30.42",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"22",
            "hour_padded":"22",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448938800",
            "pretty":"10:00 PM EST on November 30, 2015",
            "civil":"10:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"26",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"30",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"213"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"66",
         "windchill":{  
            "english":"31",
            "metric":"-1"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"31",
            "metric":"-1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"1",
         "mslp":{  
            "english":"30.41",
            "metric":"1030"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"23",
            "hour_padded":"23",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"11",
            "mon_padded":"11",
            "mon_abbrev":"Nov",
            "mday":"30",
            "mday_padded":"30",
            "yday":"333",
            "isdst":"0",
            "epoch":"1448942400",
            "pretty":"11:00 PM EST on November 30, 2015",
            "civil":"11:00 PM",
            "month_name":"November",
            "month_name_abbrev":"Nov",
            "weekday_name":"Monday",
            "weekday_name_night":"Monday Night",
            "weekday_name_abbrev":"Mon",
            "weekday_name_unlang":"Monday",
            "weekday_name_night_unlang":"Monday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"35",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"26",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"31",
         "wspd":{  
            "english":"5",
            "metric":"8"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"207"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"69",
         "windchill":{  
            "english":"31",
            "metric":"-0"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"31",
            "metric":"-0"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.4",
            "metric":"1029"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"0",
            "hour_padded":"00",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448946000",
            "pretty":"12:00 AM EST on December 01, 2015",
            "civil":"12:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"26",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"33",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"212"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"71",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.4",
            "metric":"1029"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"1",
            "hour_padded":"01",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448949600",
            "pretty":"1:00 AM EST on December 01, 2015",
            "civil":"1:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"26",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"35",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"201"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"73",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.39",
            "metric":"1029"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"2",
            "hour_padded":"02",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448953200",
            "pretty":"2:00 AM EST on December 01, 2015",
            "civil":"2:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"27",
            "metric":"-3"
         },
         "condition":"Clear",
         "icon":"clear",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
         "fctcode":"1",
         "sky":"24",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"202"
         },
         "wx":"Mostly Clear",
         "uvi":"0",
         "humidity":"75",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.39",
            "metric":"1029"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"3",
            "hour_padded":"03",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448956800",
            "pretty":"3:00 AM EST on December 01, 2015",
            "civil":"3:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"27",
            "metric":"-3"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"33",
         "wspd":{  
            "english":"2",
            "metric":"3"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"213"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"76",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.37",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"4",
            "hour_padded":"04",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448960400",
            "pretty":"4:00 AM EST on December 01, 2015",
            "civil":"4:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"28",
            "metric":"-2"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"46",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"208"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"78",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"3",
         "mslp":{  
            "english":"30.37",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"5",
            "hour_padded":"05",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448964000",
            "pretty":"5:00 AM EST on December 01, 2015",
            "civil":"5:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"33",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"28",
            "metric":"-2"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_partlycloudy.gif",
         "fctcode":"2",
         "sky":"45",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"203"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"79",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"33",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"3",
         "mslp":{  
            "english":"30.36",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"6",
            "hour_padded":"06",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448967600",
            "pretty":"6:00 AM EST on December 01, 2015",
            "civil":"6:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"28",
            "metric":"-2"
         },
         "condition":"Clear",
         "icon":"clear",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_clear.gif",
         "fctcode":"1",
         "sky":"29",
         "wspd":{  
            "english":"2",
            "metric":"3"
         },
         "wdir":{  
            "dir":"SSW",
            "degrees":"195"
         },
         "wx":"Mostly Clear",
         "uvi":"0",
         "humidity":"80",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"3",
         "mslp":{  
            "english":"30.36",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"7",
            "hour_padded":"07",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448971200",
            "pretty":"7:00 AM EST on December 01, 2015",
            "civil":"7:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"34",
            "metric":"1"
         },
         "dewpoint":{  
            "english":"29",
            "metric":"-2"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/partlycloudy.gif",
         "fctcode":"2",
         "sky":"37",
         "wspd":{  
            "english":"1",
            "metric":"2"
         },
         "wdir":{  
            "dir":"SSE",
            "degrees":"156"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"82",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"34",
            "metric":"1"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.35",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"8",
            "hour_padded":"08",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448974800",
            "pretty":"8:00 AM EST on December 01, 2015",
            "civil":"8:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"36",
            "metric":"2"
         },
         "dewpoint":{  
            "english":"30",
            "metric":"-1"
         },
         "condition":"Partly Cloudy",
         "icon":"partlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/partlycloudy.gif",
         "fctcode":"2",
         "sky":"38",
         "wspd":{  
            "english":"2",
            "metric":"3"
         },
         "wdir":{  
            "dir":"SE",
            "degrees":"129"
         },
         "wx":"Partly Cloudy",
         "uvi":"0",
         "humidity":"79",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"36",
            "metric":"2"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.36",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"9",
            "hour_padded":"09",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448978400",
            "pretty":"9:00 AM EST on December 01, 2015",
            "civil":"9:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"39",
            "metric":"4"
         },
         "dewpoint":{  
            "english":"32",
            "metric":"0"
         },
         "condition":"Mostly Cloudy",
         "icon":"mostlycloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/mostlycloudy.gif",
         "fctcode":"3",
         "sky":"65",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"113"
         },
         "wx":"Mostly Cloudy",
         "uvi":"0",
         "humidity":"74",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"39",
            "metric":"4"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.36",
            "metric":"1028"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"10",
            "hour_padded":"10",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448982000",
            "pretty":"10:00 AM EST on December 01, 2015",
            "civil":"10:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"42",
            "metric":"6"
         },
         "dewpoint":{  
            "english":"33",
            "metric":"1"
         },
         "condition":"Overcast",
         "icon":"cloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/cloudy.gif",
         "fctcode":"4",
         "sky":"80",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"115"
         },
         "wx":"Cloudy",
         "uvi":"1",
         "humidity":"70",
         "windchill":{  
            "english":"39",
            "metric":"4"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"39",
            "metric":"4"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"2",
         "mslp":{  
            "english":"30.34",
            "metric":"1027"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"11",
            "hour_padded":"11",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448985600",
            "pretty":"11:00 AM EST on December 01, 2015",
            "civil":"11:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"33",
            "metric":"1"
         },
         "condition":"Overcast",
         "icon":"cloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/cloudy.gif",
         "fctcode":"4",
         "sky":"86",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"SE",
            "degrees":"128"
         },
         "wx":"Cloudy",
         "uvi":"1",
         "humidity":"66",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"1",
         "mslp":{  
            "english":"30.32",
            "metric":"1027"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"12",
            "hour_padded":"12",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448989200",
            "pretty":"12:00 PM EST on December 01, 2015",
            "civil":"12:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"32",
            "metric":"0"
         },
         "condition":"Overcast",
         "icon":"cloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/cloudy.gif",
         "fctcode":"4",
         "sky":"93",
         "wspd":{  
            "english":"8",
            "metric":"13"
         },
         "wdir":{  
            "dir":"SE",
            "degrees":"124"
         },
         "wx":"Cloudy",
         "uvi":"1",
         "humidity":"62",
         "windchill":{  
            "english":"40",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"40",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"1",
         "mslp":{  
            "english":"30.3",
            "metric":"1026"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"13",
            "hour_padded":"13",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448992800",
            "pretty":"1:00 PM EST on December 01, 2015",
            "civil":"1:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"32",
            "metric":"0"
         },
         "condition":"Overcast",
         "icon":"cloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/cloudy.gif",
         "fctcode":"4",
         "sky":"97",
         "wspd":{  
            "english":"8",
            "metric":"13"
         },
         "wdir":{  
            "dir":"SE",
            "degrees":"137"
         },
         "wx":"Cloudy",
         "uvi":"1",
         "humidity":"62",
         "windchill":{  
            "english":"40",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"40",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"3",
         "mslp":{  
            "english":"30.28",
            "metric":"1025"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"14",
            "hour_padded":"14",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1448996400",
            "pretty":"2:00 PM EST on December 01, 2015",
            "civil":"2:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"32",
            "metric":"0"
         },
         "condition":"Overcast",
         "icon":"cloudy",
         "icon_url":"http://icons.wxug.com/i/c/k/cloudy.gif",
         "fctcode":"4",
         "sky":"98",
         "wspd":{  
            "english":"8",
            "metric":"13"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"122"
         },
         "wx":"Cloudy",
         "uvi":"0",
         "humidity":"62",
         "windchill":{  
            "english":"40",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"40",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"14",
         "mslp":{  
            "english":"30.26",
            "metric":"1025"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"15",
            "hour_padded":"15",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449000000",
            "pretty":"3:00 PM EST on December 01, 2015",
            "civil":"3:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"33",
            "metric":"1"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"8",
            "metric":"13"
         },
         "wdir":{  
            "dir":"SE",
            "degrees":"127"
         },
         "wx":"Few Showers",
         "uvi":"0",
         "humidity":"64",
         "windchill":{  
            "english":"40",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"40",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.0",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"32",
         "mslp":{  
            "english":"30.23",
            "metric":"1024"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"16",
            "hour_padded":"16",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449003600",
            "pretty":"4:00 PM EST on December 01, 2015",
            "civil":"4:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"34",
            "metric":"1"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"8",
            "metric":"13"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"119"
         },
         "wx":"Showers",
         "uvi":"0",
         "humidity":"68",
         "windchill":{  
            "english":"40",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"40",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"41",
         "mslp":{  
            "english":"30.21",
            "metric":"1023"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"17",
            "hour_padded":"17",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449007200",
            "pretty":"5:00 PM EST on December 01, 2015",
            "civil":"5:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"35",
            "metric":"2"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"113"
         },
         "wx":"Showers",
         "uvi":"0",
         "humidity":"71",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"53",
         "mslp":{  
            "english":"30.19",
            "metric":"1022"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"18",
            "hour_padded":"18",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449010800",
            "pretty":"6:00 PM EST on December 01, 2015",
            "civil":"6:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"36",
            "metric":"2"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"99",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"ESE",
            "degrees":"109"
         },
         "wx":"Showers",
         "uvi":"0",
         "humidity":"76",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"53",
         "mslp":{  
            "english":"30.18",
            "metric":"1022"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"19",
            "hour_padded":"19",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449014400",
            "pretty":"7:00 PM EST on December 01, 2015",
            "civil":"7:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"37",
            "metric":"3"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"98",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"E",
            "degrees":"102"
         },
         "wx":"Showers",
         "uvi":"0",
         "humidity":"79",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.03",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"49",
         "mslp":{  
            "english":"30.16",
            "metric":"1021"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"20",
            "hour_padded":"20",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449018000",
            "pretty":"8:00 PM EST on December 01, 2015",
            "civil":"8:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"39",
            "metric":"4"
         },
         "condition":"Rain",
         "icon":"rain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
         "fctcode":"13",
         "sky":"100",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"E",
            "degrees":"93"
         },
         "wx":"Rain",
         "uvi":"0",
         "humidity":"82",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.03",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"77",
         "mslp":{  
            "english":"30.14",
            "metric":"1021"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"21",
            "hour_padded":"21",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449021600",
            "pretty":"9:00 PM EST on December 01, 2015",
            "civil":"9:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"44",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"40",
            "metric":"4"
         },
         "condition":"Rain",
         "icon":"rain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_rain.gif",
         "fctcode":"13",
         "sky":"100",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"E",
            "degrees":"83"
         },
         "wx":"Rain",
         "uvi":"0",
         "humidity":"84",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.03",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"82",
         "mslp":{  
            "english":"30.1",
            "metric":"1019"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"22",
            "hour_padded":"22",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449025200",
            "pretty":"10:00 PM EST on December 01, 2015",
            "civil":"10:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"40",
            "metric":"4"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"ENE",
            "degrees":"69"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"85",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"83",
         "mslp":{  
            "english":"30.07",
            "metric":"1018"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"23",
            "hour_padded":"23",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"1",
            "mday_padded":"01",
            "yday":"334",
            "isdst":"0",
            "epoch":"1449028800",
            "pretty":"11:00 PM EST on December 01, 2015",
            "civil":"11:00 PM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Tuesday",
            "weekday_name_night":"Tuesday Night",
            "weekday_name_abbrev":"Tue",
            "weekday_name_unlang":"Tuesday",
            "weekday_name_night_unlang":"Tuesday Night",
            "ampm":"PM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"42",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"6",
            "metric":"10"
         },
         "wdir":{  
            "dir":"ENE",
            "degrees":"67"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"89",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"65",
         "mslp":{  
            "english":"30.04",
            "metric":"1017"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"0",
            "hour_padded":"00",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"2",
            "mday_padded":"02",
            "yday":"335",
            "isdst":"0",
            "epoch":"1449032400",
            "pretty":"12:00 AM EST on December 02, 2015",
            "civil":"12:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Wednesday",
            "weekday_name_night":"Wednesday Night",
            "weekday_name_abbrev":"Wed",
            "weekday_name_unlang":"Wednesday",
            "weekday_name_night_unlang":"Wednesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"42",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"ENE",
            "degrees":"60"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"91",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"63",
         "mslp":{  
            "english":"30.0",
            "metric":"1016"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"1",
            "hour_padded":"01",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"2",
            "mday_padded":"02",
            "yday":"335",
            "isdst":"0",
            "epoch":"1449036000",
            "pretty":"1:00 AM EST on December 02, 2015",
            "civil":"1:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Wednesday",
            "weekday_name_night":"Wednesday Night",
            "weekday_name_abbrev":"Wed",
            "weekday_name_unlang":"Wednesday",
            "weekday_name_night_unlang":"Wednesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"43",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"NE",
            "degrees":"57"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"93",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"64",
         "mslp":{  
            "english":"29.98",
            "metric":"1015"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"2",
            "hour_padded":"02",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"2",
            "mday_padded":"02",
            "yday":"335",
            "isdst":"0",
            "epoch":"1449039600",
            "pretty":"2:00 AM EST on December 02, 2015",
            "civil":"2:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Wednesday",
            "weekday_name_night":"Wednesday Night",
            "weekday_name_abbrev":"Wed",
            "weekday_name_unlang":"Wednesday",
            "weekday_name_night_unlang":"Wednesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"43",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"7",
            "metric":"11"
         },
         "wdir":{  
            "dir":"NE",
            "degrees":"48"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"93",
         "windchill":{  
            "english":"41",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"41",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"62",
         "mslp":{  
            "english":"29.95",
            "metric":"1014"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"3",
            "hour_padded":"03",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"2",
            "mday_padded":"02",
            "yday":"335",
            "isdst":"0",
            "epoch":"1449043200",
            "pretty":"3:00 AM EST on December 02, 2015",
            "civil":"3:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Wednesday",
            "weekday_name_night":"Wednesday Night",
            "weekday_name_abbrev":"Wed",
            "weekday_name_unlang":"Wednesday",
            "weekday_name_night_unlang":"Wednesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"45",
            "metric":"7"
         },
         "dewpoint":{  
            "english":"43",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"100",
         "wspd":{  
            "english":"5",
            "metric":"8"
         },
         "wdir":{  
            "dir":"NNE",
            "degrees":"16"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"91",
         "windchill":{  
            "english":"42",
            "metric":"5"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"42",
            "metric":"5"
         },
         "qpf":{  
            "english":"0.02",
            "metric":"1"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"62",
         "mslp":{  
            "english":"29.94",
            "metric":"1014"
         }
      },
      {  
         "FCTTIME":{  
            "hour":"4",
            "hour_padded":"04",
            "min":"00",
            "min_unpadded":"0",
            "sec":"0",
            "year":"2015",
            "mon":"12",
            "mon_padded":"12",
            "mon_abbrev":"Dec",
            "mday":"2",
            "mday_padded":"02",
            "yday":"335",
            "isdst":"0",
            "epoch":"1449046800",
            "pretty":"4:00 AM EST on December 02, 2015",
            "civil":"4:00 AM",
            "month_name":"December",
            "month_name_abbrev":"Dec",
            "weekday_name":"Wednesday",
            "weekday_name_night":"Wednesday Night",
            "weekday_name_abbrev":"Wed",
            "weekday_name_unlang":"Wednesday",
            "weekday_name_night_unlang":"Wednesday Night",
            "ampm":"AM",
            "tz":"",
            "age":"",
            "UTCDATE":""
         },
         "temp":{  
            "english":"46",
            "metric":"8"
         },
         "dewpoint":{  
            "english":"43",
            "metric":"6"
         },
         "condition":"Chance of Rain",
         "icon":"chancerain",
         "icon_url":"http://icons.wxug.com/i/c/k/nt_chancerain.gif",
         "fctcode":"12",
         "sky":"99",
         "wspd":{  
            "english":"3",
            "metric":"5"
         },
         "wdir":{  
            "dir":"NNW",
            "degrees":"341"
         },
         "wx":"Light Rain",
         "uvi":"0",
         "humidity":"89",
         "windchill":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "heatindex":{  
            "english":"-9999",
            "metric":"-9999"
         },
         "feelslike":{  
            "english":"46",
            "metric":"8"
         },
         "qpf":{  
            "english":"0.01",
            "metric":"0"
         },
         "snow":{  
            "english":"0.0",
            "metric":"0"
         },
         "pop":"62",
         "mslp":{  
            "english":"29.92",
            "metric":"1013"
         }
      }
   ]
}
```
