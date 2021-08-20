## Usage
`env-mapper [TO_ENVIRONMENT_VARIABLE_NAME:FROM_ENVIRONMENT_VARIABLE_NAME]... -- <command> [argument]...`

Executes `command` (required) with one or more `argument`s (optional) in the environment with
`TO_ENVIRONMENT_VARIABLE_NAME` set to the value of `FROM_ENVIRONMENT_VARIABLE_NAME`.

### Example
`env-mapper RUN_AS:USER WORK_DIRECTORY:PWD -- /usr/bin/env -v`

Will run `/usr/bin/env` with argument `-v`
The environment variable `RUN_AS` will be set to the value of environment variable `USER` or empty value if `USER` is unset.
`WORK_DIRECTORY` will be set to the value of `PWD` if any.

The command is equivalent to bourne shell:

`RUN_AS="${USER}" WORK_DIRECTORY="${PWD}" /usr/bin/env -v`

and is designed to be used on the systems (primarily lightweight containers) where the shell is not part of the system image to remap the environment variables injected by runtime to the names expected by the programs.

 https://github.com/asecurityteam/env-mapper for more documentation.
