# Go Snippetbox

## Structure
- `cmd` directory contains the application-specific code for the executable applications in the project. Like a web application.
- `internal` directory contains the non-application-specific code. Like validation helpers and the SQL models.
- `ui` directory contains the user-interface assets.
    - `html` directory contains the HTML templates.
    - `static` directory contains CSS and images files

## Avoid directory listings
Add a blank `index.html` file to the root of static files directory and all sub-directories.