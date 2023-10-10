# TooGoodToGo Notifier

A simple application to notify you via email when items become available on too good to go.

## Prerequisites

This script leverages `msmtp` and was designed for use on a Raspberry Pi.

## Getting Started

To initialise the script run the application with the `i` flag

```shell
./tgtg-notifier -i <your account email>
```

This will attempt to authenticate you and then make a configuration containing your favorite items from Too Good To Go.
Feel free to adjust the configuration in order to specify which items you wish to receive notifications from.

To perform a check simply run:
```shell
./tgtg-notifier
```
Please ensure you have initialised one time before trying to perform a check.
Notifications only occur once per day, per item.

## Configuring msmtp 

By default, the configuration looks for a msmtp account called `gmail`
To adjust this, simply change the `email_config` section of the configuration.

## Recommendation

Configure a crontab task to run the script in preferable intervals.

Example of executing every 15 minutes:
```shell
*/15 * * * * cd <location of script> && <script> > <location of script>/logfile.log 2>&1
```
ensure you change the values above to suit your needs.

