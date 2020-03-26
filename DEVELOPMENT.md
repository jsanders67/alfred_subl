## Building
To build the go src files, simply run `make` in the project directory

## Compiling Alfredworkflow
To compile your changes into an alfred workflow, simply run `make alfred`

## Developing
If you are working on the workflow and actively making changes to the source code, you will want to 
run `make develop`. This creates a symlink from the project directory to the alfred workflows directory. Then all that needs to be done to make changes is to run `make build` and the binaries alfred uses to
execute the workflows will be picked up automatically

## Testing
To test, run `make test`. This will remove any existing workflows, compile the workflow from the current directory and install it into alfread. You should be prompted to import the workflow in the alfred console. This is more of an end-to-end tests of what experience people will have when installing the workflow on their own. 

## Contributing
To contribute, compile the workflow and commit it to git. Be sure to run `make test` and try `isen refresh` and `isen <test query>` to test your workflow locally before commiting.