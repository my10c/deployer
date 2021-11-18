# deployer

The blue - green deployer API server

```
usage: deployer [-h|--help] [-c|--configFile "<value>"] [-R|--runAsRoot]
                [-S|--showconfig] [-v|--version] [-i|--info] [-l|--logFile
                "<value>"] [-M|--logMaxSize "<value>"] [-B|--logMaxBackups
                "<value>"] [-A|--logMaxAge "<value>"]

                A blue - green deployment API server

Arguments:

  -h  --help           Print help information
  -c  --configFile     Path to the configuration file. Default:
                       /etc/deployer/config.json
  -R  --runAsRoot      Do run as root. Default: false
  -S  --showconfig     Show configuration example. Default: false
  -v  --version        Show version. Default: false
  -i  --info           Show information. Default: false
  -l  --logFile        Path to the log file. Default: /var/log/deployer.log
  -M  --logMaxSize     Max size of the log file (MB). Default: 512
  -B  --logMaxBackups  Max log file count. Default: 28
  -A  --logMaxAge      Max days to keep a log file. Default: 30
```

### Installation
- install the binary in any locatiom, recommeded under /usr/local/sbin
- create the configuration file, defaul should be /etc/deployer/config.json
- adjust the config file and adjust the API name and script to be execute by the API call
  make sure the script are protected an have the execute bit enabled
- recommend to run as root, the -R flag


