Svein's log
============
## 08/05/2019
#### Worked on
* Thesis
  * Requirements chapter
* Updating logs

## 07/05/2019
#### Worked on
* Thesis
  * Development Process chapter
* Updating logs
  
## 06/05/2019
#### Worked on
* Thesis
  * Development Process chapter

## 03/05/2019
#### Worked on
* Thesis
  * Development Process chapter

## 02/05/2019
#### Worked on
* Thesis
  * Development Process chapter

## 01/05/2019
#### Worked on
* Thesis
  * Development Process chapter

## 30/04/2019
#### Worked on
* Thesis
  * Development Process chapter

## 29/04/2019
#### Worked on
* Fixed filtering in choices, on the Form Builder

## 27/04/2019
#### Worked on
* Added more views for group-feature

## 26/04/2019
#### Worked on
* Worked on another project, other course 

## 25/04/2019
#### Worked on
* Worked on another project, other course

## 24/04/2019
#### Worked on
* Worked on another project, other course 

## 23/04/2019
#### Worked on
* Groups back-/front-end
  * Creating/joining/leaving groups
  * Functions for future implementations

## 22/04/2019
#### Worked on
* Group back-end

## 17/04/2019
#### Worked on
* Removed some more deprecated code
* Fixed error on updating assignments
* Group-based back-/front-end
* Fixed old bug in deleting submissions
  * Renamed functions for convenience

## 15/04/2019
#### Worked on
* Removed some deprecated code
* Added new templates for group-delivery
* Added code for implementing groups
* Simple refactoring, usage of services

## 12/04/2019
#### Worked on
* Group-delivery
  * Added groups to the database
  * Let user create, join and leave group
#### Bad
* Big task with little time to do (group-delivery)
#### Good
* Added some missing commenting from before
#### Decisions
* 

## 11/04/2019
#### Worked on
* Added a new handler/view/page for reviews done by a user (admin-view)

## 10/04/2019
#### Worked on
* Added a new plugin for checking if user had been reviewed
* Added current user to all views
* Added ability to update reviews after submission
* Merged with master

## 09/04/2019
#### Worked on
* Testing exporting data to Excel-file
* Refactoring some code
* Merged with another branch

## 08/04/2019
#### Worked on
* Fetching raw data about review scores
* Processing review-score data
* Displaying review-data

## 04/04/2019
#### Worked on
* Merged with master branch
* Fixed imports after renaming of the repository
* Fixed import path for config-file
#### Good
* New laptop worked fine
#### Bad
* Previous laptop broke, so tested out a new one

## 03/04/2019
#### Worked on
* Implemented Sprintf for templates
* Added link for url-fields on rendering for some users who could not open the url's
* Merged with another branch
#### Decisions
* Shorten down float numbers to two decimals for better readability

## 02/04/2019
#### Worked on
* Fixed a bug on the Form Builder
* Made some QoL changes to the design on submission-tables
* Removed label on comments-title (Feedback from students)
* Worked statistics for assignments, calculating average score and standard deviation
  * Displaying statistics for the course-submissions
  * Calculating radio-selection linear
  * Created own functions for statistics outside handler
#### Decisions
* Created internal lib for calculating standard deviation and average score for reviews

## 01/04/2019
#### Worked on
* Merged with the master branch

## 29/03/2019
#### Worked on
* Fixing TODO's
* Updated code
  * More informative log messages
* Adding more security to the application
* Added selection for SimpleMDE (QoL)
* Removed deprecated functions
#### Decisions
* Redirect all attempts to go to `/admin*` to `/`

## 24/03/2019
#### Worked on
* Implemented default weight and field type to the form builder
* Implemented deleting submissions by admin
* Implemented Bootstrap-table extension
* Fixed bugs for beta
* Added default weight and field to form-builder
* Made submissions withdrawable
* Implemented feedback to student after submission
#### Good
* Deployed beta for testing with real students and a real assignment
#### Bad
* Wrong query-string overwrote all submissions 

## 22/03/2019
#### Worked on
* Finished refactoring of service implementation
* Implemented weights display for reviews
#### Good
* Refactoring done

## 21/03/2019
#### Worked on
* Continue: Refactoring
#### Bad
* Not full group, so many hours of work in a few days

## 20/03/2019
#### Worked on
* Refactoring
  * Adding service/repository pattern 
#### Good
* Works good, allow us to be less reliable of SQL queries
#### Bad
* Takes a lot of time, not full group, so did 15 hours today
#### Decisions
* Refactoring DAL (data access layer) away from controllers

## 13/03/2019
#### Worked on
Fixed things from feedback from the pull request
#### Good
* One step closer to a good alpha
#### Bad
#### Decisions
* Changed time formatting to norwegian time

## 11/03/2019
#### Worked on
Finished of  the alpha work for reviews, displaying, submitting them.
#### Good
* Got the pull request out for the rest of the group, early on monday morning

# Week 10 report
#### Worked on
* Review features
* Was gone 3 days for a work interview, lost does days for production
#### Bad
* Lost 3 days of programming due to traveling


## 08/03/2019
#### Worked on
* Created a few new view plugins for helping with the review-form
  * Increment
  * Split Choices

## 07/03/2019
#### Worked on
* Generating the review form
* Remodelling the `user_reviews` table

## 04/03/2019
#### Worked on
* Generating the review form

## 03/03/2019
#### Worked on
* Updating reviews
* Creating reviews
#### Good
* Found bug in Javascript-framework, but was easily fixed 

## 02/03/2019
#### Worked on
* Updating submissions
* Creating reviews
* Tests for the view-plugins
* Parsing JSON to Javascript Form
#### Good
* Javascript framework for dynamic forms is working great, and easily changeable

## 01/03/2019
#### Worked on
* Nothing, was sick

## 28/02/2019
#### Worked on
* Assignment page
* View plugins
#### Good
* View plugins works great

## 27/02/2019
#### Worked on
* QoL (Navbar)

## 26/02/2019
#### Worked on
* Assignment update/view
* Setup OpenStack VM
#### Good
* Parsing `time.Time` to `datetime-local` from Go to HTML
* OpenStack: first configuration worked as a charm

## 25/02/2019
#### Worked on
* Settings page
* Assignment update/view
#### Good
* Group talked about the next weeks need-to-be-done stuff
#### Bad
* Wasted some time on the settings-stuff

## Week 8 report
#### Worked on
* I have been working on the assignment and submission features for the lecturer-side of the software.
* Making the forms for creating the submission and review-forms dynamic. Doing this with Javascript, 
making a small reusable library with some configuration.
* Database for the dynamic form
#### Decisions
* The database tables were changed a few times throughout the week, but I think the last version
will be more suited for the application, as it is scalable.
#### Good
* Got the story done, and feel the way it ended up, will be good for the rest of the application,
making it scalable for other types of submissions, and assignments.
#### Bad
* Used to much time on this story
* Had problems with the database, cause of old relations, making it hard to merge new tables together with old once. 

## 22/02/2019
* Did some QoL on the form submits
* Fixed some bugs of redirecting and form-validation
* Researched best practice for Go

#### Research examples
Simple example with `Courses`

Go-code:
```go
package research

type Course struct {
	ID int
	Data map[string]string
}

// Current usage
type Courses struct {
	Items []Courses
}

// Researched usage
type Courses []Courses
```

HTML-code:
```html
<!-- Current usage -->
{{range .Courses.Items}}
    <div class="col">...</div>
{{end}}

<!-- Researched usage -->
{{range .Courses}}
    <div class="col">...</div>
{{end}}
```

## 20/02/2019
* Created a working tables for the assignments, submissions, forms and fields.
* Did pull request on a bigger card, but group found a lot of bugs, got most of them fixed
* Found a new way of doing the `forms` and `fields` tables:
#### `forms`
| id | prefix | name | description | created |
| --- | --- | --- | --- | --- |
| PK | prefix for fields (HTML ) | Display name | Description .. | TIMESTAMP |

#### `fields`
| id | form_id  | type | name | label | description | priority | weight | choices |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| PK | FK | type of field (text, radio, checkbox) | Name of field (HTML) | Label (Display name) | Description | Order number | Weight for grading | Choices (Array, split by ',')

## 19/02/2019
* Short day of working. Got the basic framework for the dynamic form together

## 18/02/2019
* Talked with product owner about the database-design for the custom forms, and found a solution together that will not be too hard to implement.
  * Later found a more generic way to design the database, with a `forms` and `fields` table:
  
#### Forms-table
| id  | name | description | created | prefix |
| --- | ---- | ----------- | ------- | ------ |
| PK  | Display name | Description .. | TIMESTAMP | prefix for fields (HTML)
#### Fields-table
| id | data | order | form_id |
| --- | --- | --- | --- |
| PK | JSON | needed? | FK for `forms` |

## Week 7 Report
This week I have been looking into data structure for the dynamic review form, and where and what should take care for the input/output for the form. From the research done, it seams like Javascript will be the best choice of creating and parsing data to strings, with JSON, as Javascript is well-equipped with JSON-functions.

The database is also a challenge, because of the dynamic form, the database has to be design for a agile software, that needs to be flexible today, as well in a few years. I think I have found a good solution for the desgin of the database, but it is hard to implement cause of the auto-generated schema we are using from MySQL Workbench, but I think we need to rewrite the database-schema, and look more into normalization of relational databases, to make it flexible enough for the requirement-specs.

Have also been working on the last touches of the restructuring of the project, making it more flexible with a MVC-architecture. Making less files, and all the files in a folder, does the same type of tasks.  

## 15/02/2019
* Updated teh time.Time convert-function from datetime-local (HTML).
* Looked at the database for assignments, 

## 14/02/2019
* Worked on the design of assignments with the peer reviews, thinking about the data-structure, and how to store it, and where it should be written/read.

## 13/02/2019
* Worked on the assignments page for admins, looking at data-structure, and an easy approach for designing the form.

## 12/02/2019
* Worked on restructuring the folder-/file-structure for the project. Making it a MVC-structure.

## 11/02/2019
* Started the planning for the project restructure, for a more efficient development later on in the project.

# Week 6 Report
I have worked mainly with Bootstrap to get a concise design on every page of the application. Templating has also taken up quite a bit of time, as I did not have much knowledge of this before, and had to read about this, and talk to the other team members about this issue. It was resolved in a straightforward way, but afterwards I see that it could have been solved even better, with the creation of features to simply enter simple parameters to be able to change the whole page, as well as help reuse code.

I have also worked part with the dynamic form to be used for peer review. Since I have a good deal of knowledge from before with dynamic front-end programming, this task went pretty smoothly. Met some challenges with the prioritization of fields, but found a solution in the end that makes it easy for both me and the person who will set up the form. Will look later on this, but it works pretty well as it does now. Only thing missing is adding a button to remove items from the form.

The group had a discussion on the database structure and form structure of the application, and found that we could reduce the number of fields in the database, using Markdown as word processing, which means that user has great freedom in relation to text and content that can added to the subject page.

## 08/02/2019
* Worked more on the assignments-form
* Had some problems with ordering on the form, but found an easier solution that works perfectly fine

## 07/02/2019
* Added navbar to every pag
* Worked on the assignments-form, for peer-reviews
* Refactored some of the previous code, cause of navbar on every page

## 06/02/2019
* Added another function for loading JSON-data from file
  * Dummy-data for reuse, displaying data on site
* Created form for creating new courses
* Created am flexible error handler

## 05/02/2019
 * Made some simple data-structures for page-data, and made a more agile template for the page-title, menu and content.
 * Made a simple function for loading data from JSON-file
   * Menu-items

## 04/02/2019
Worked on nested templates, and design for the website, both for the main site, and the dashboard for admin.
Created one more test for the HTTP-requests, checking for response body-size.

## 01/02/2019
Worked on the project plan:
 * Rewrote some phrases
 * QA

## 31/01/2019
Worked on the project plan:
 * Risk Analysis
 * Main division of the project

## 30/01/2019
Worked on the project plan:
 * Subject area
 * Limits
 * Gantt-implementation

## 28/01/2019
Worked on the project plan:
 * Technology, Business, Project Group
 * Risk Analysis