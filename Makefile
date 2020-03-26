PROJECT_NAME := AlfredSublimeOpen
WORKFLOW_NAMESPACE := "com.sanjams.alfred.sublimeopen"
WORKFLOW_DIR := /Users/$(USER)/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows
CURRENT_WORKFLOW_DIRS := $(shell grep -l "$(WORKFLOW_NAMESPACE)" $(WORKFLOW_DIR)/*/info.plist | sed -e 's/ /\\ /g')
UUID := $(shell uuidgen)
PWD := $(shell pwd)
CUR_DIR := $(shell basename $(shell pwd))
SRC_FILES := $(shell find src -name "*.go")

build:
	go build -o out/open_workspace src/open_workspace.go

alfred: build
	rm $(PROJECT_NAME).alfredworkflow || true
	zip -r -X $(PROJECT_NAME).alfredworkflow *

clear-current-workflows:
	for f in $(CURRENT_WORKFLOW_DIRS); do \
		echo $$f | sed -e 's/ /\\ /g' | xargs dirname | xargs -I {} rm -rf "{}" ; \
	done
	ls -l $(WORKFLOW_DIR)

develop: clear-current-workflows
	ln -s $(PWD) $(WORKFLOW_DIR)/user.workflow.$(UUID)

test: clear-current-workflows alfred
	open $(PROJECT_NAME).alfredworkflow

