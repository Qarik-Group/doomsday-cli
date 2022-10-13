# Doomsday CLI

[Doomsday](https://github.com/doomsday-project/doomsday) is a server which can be configured to track certificates from different storage backends (Vault, Credhub, Pivotal Ops Manager, or actual websites) and provide a tidy view into when certificates will expire. Doomsday CLI provides the option to get all required information about certificates expiration in the JSON format and apply filters to get more detailed results.

## Installation

Installation process is very simple.
* Check available [release versions](https://github.com/starkandwayne/doomsday-cli/releases)
* Download binary for your operating system and architecture and save it as `doomsday` file
* Open your terminal console e.g. iTerm and open directory where the binary was downloaded
* Add execution permission to `doomsday` file and move it to `/usr/local/bin/` directory with the following command
   ```
   chmod +x ./doomsday && mv ./doomsday /usr/local/bin/
   ```

## Using the CLI

### Command line global options
In the Doomsday CLI, command line global options are parameters you can use to override the default server configuration and user authentication settings, or environment variable settings.

You can use the following command line options to override the default configuration settings.
```
--server, -s <string>
  Doomsday server URL

--username, -u <string>
  Username for server authentication

--password, -p <string>
  Password for server authentication
```

### Environment variables
The following examples show how you can configure environment variables for the doomsday CLI instead specyfing command line option flags.

```
$ export DOOMSDAY_SERVER=https://127.0.0.1:443
$ export DOOMSDAY_USER=doomsday
$ export DOOMSDAY_PASSWORD=password
```

### Getting help
You can get help with any command when using the Doomsday CLI. To do so, simply type `help` at the end of a command name.

For example, the following command displays help for the general Doomsday CLI options and the available top-level commands.
```
$ doomsday help
```
The following command displays the available certificates specific commands.
```
$ doomsday certificates help
```

### Command structure
The CLI uses a multipart structure on the command line that must be specified in this order:
1. The base call to the `doomsday` program.
2. The top-level command, which typically corresponds to a Doomsday server API endpoint supported by the CLI.
3. The subcommand that specifies which operation to perform.
4. General CLI options or parameters required by the operation.
```
$ doomsday <command> <subcommand> [options and parameters]
```
Parameters can take different types of input values, such as numbers or strings.

### Command examples
Here are some of the example commands used to interact with Doomsday server.
* Get all certificates and their corresponding data
   ```
   $ doomsday certificates
   ```
* Get all `expired` certificates and their corresponding data
   ```
   $ doomsday certificates expired
   ```
* Get all certificates that will expire in the next 30 days
   ```
   $ doomsday certificates willexpire --days 30
   ```
* Refresh certificates cache
   ```
   $ doomsday refresh
   ```
* Get Doomsday server version and enabled authentication method
   ```
   $ doomsday info
   ```
### Output
All commands will return output in the JSON format.