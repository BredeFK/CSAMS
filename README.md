# ~NTNU-Bachelor-Management-System-For-CS-Assignments~
# Computer Science Assignment Management System

## Setup guides
- In openstack, create instance and security group.
- Security group should open port 80 (for http access) and possibly 8080 if you want to access adminer for database management
- In a fresh linux instance, install Docker and Git
- Pull this repo
- Edit the Makefile to change envvars for passwords and email authentication (do not commit these changes to git)
- run 'sudo make new' to apply envvars and build the service

### Importing old database
- If this is a fresh instance then you will have to open port 8080 and access adminer
- Then you want to run SQL query and copy paste the old database file there.

### Exporting database
- Access adminer through port 8080
- The should be a export button somewhere
- Export all table data, do not include the tables themselves. you want this either as sql text in browser or as a downloadable txt file.
- Save it somewhere safe

### Service crash recover
- If the server crashes for some unknown reason, run 'sudo make restart'
- Verify that the database has been attached
- if not, run 'sudo make stop' and 'sudo make run'

## Links
* [Github](https://github.com/JohanAanesen/CSAMS)
* [Trello](https://trello.com/bachelor531)
* [ShareLatex](https://www.overleaf.com/project/5c3491a162ba3128fda8c11d)
* [Discord](https://discord.gg/rZ4zg2R)
* [Google Drive](https://drive.google.com/drive/folders/1kiQiBj12zrn45q6QOfXefrzgNb4fZhyW?usp=sharing)
* [Toggl](https://toggl.com)
* [UML](https://www.lucidchart.com/invitations/accept/421b3f38-581e-4790-80f7-3d43604a717c)

## Members
| Role | Name | E-mail |
| -------- | -------- | ------- |
| Project Leader | Johan Aanesen | johanaan@stud.ntnu.no |
| Member | Brede Fritjof Klausen | bredefk@stud.ntnu.no |
| Member | Svein Are Danielsen | sveiad@stud.ntnu.no |

## Project Owner
* [Christopher Frantz](https://www.ntnu.no/ansatte/christopher.frantz)

## Supervisor
* [Ivar Farup](https://www.ntnu.no/ansatte/ivar.farup)

## OBS - Timezones!
This project is in the norwegian timezone!
### SQL
Uses the Go function below combined with `ConvertTimeStampToString(time.Time)`
### GO
[`webservice/shared/util/time.go`](https://github.com/JohanAanesen/CSAMS/blob/master/webservice/shared/util/time.go)
```Go
func GetTimeInNorwegian() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Oslo")

	return time.Now().In(loc)
}
```
### Javascript
[`webservice/static/js/time.js`](https://github.com/JohanAanesen/CSAMS/blob/master/webservice/static/js/time.js)
```Js
function getTimeInNorwegian() {
    let norTime = new Date().toLocaleString("no-no", {timeZone: "Europe/Oslo"});
    return new Date(norTime);
}
```


## Commit messages guideline
Start with the number `#42` where the number is the number to the `card` you're working on Trello followed by:
* update	- updated function/file
* add	- added function/file
* fix	- fixed a bug
* remove	- removed a function/file
* test	- added a test for a function
* refactor - refactored the code-structure, moved a file or changed the visual representation of code.

Aka #CARDNR ACTION BRIEF_EXPLANATION

Examples:
* "#2 add database.sql initial table file"
* "#8 fix changing password bug"
* "#14 update weekly report for week 2"

## Internal Peer Review Checklist
- [ ] Compiles
- [ ] Testing is sufficient and passing
- [ ] Branch completes the assigned tasks (from trello card)
- [ ] Code is well commented and tidy
- [ ] Is there something that could be done better? propose update
- [ ] Leave feedback as comment on the pull request :)

## Go Module
* See [quick start](https://github.com/golang/go/wiki/Modules#quick-start) for Modules

