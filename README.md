# project templates

Repository with templates for different kinds of projects.

## Setup

Clone repo and copy the `newp` (new project) script to a diretory in your PATH.

```sh
git clone git@github.com:jesperkha/project-templates.git

# Set templates path
sed -i "s|^TEMPLATES_DIR=.*|TEMPLATES_DIR=$(pwd)|" newp

# Put script in PATH
ln -sf "$(pwd)/newp" ~/.local/bin/newp
```

## Use

Run `newp` to create a new project.
