## Alfred Sublime Text Workflow
This is an [Alfred workflow](https://www.alfredapp.com/) designed to quickly list common workspaces and open
them in sublime text

## Installation
Simply download the [AlfredIsengard.alfredworkflow](AlfredSublimeOpen.alfredworkflow) file in this repository and open it in Alfred

## Usage
- Specify which directories to search by setting the workflow variable `WORKSPACE_DIRS`. This can be a comma
separated list of full paths to directories.
- In your alfred spotlight, run `subl <Name of workspace>` to open your workspace 

## Requirements
* [fzf](https://github.com/junegunn/fzf)