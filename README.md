# auto-report

Auto daily report for HITSZ Students.

## Usage

### Report Once

**Golang**

```shell
go run main.go -u your-studentID -p your-password -e your-email
```

**Docker**

```shell
docker run --rm rocketeerli/auto-report -u your-studentID -p your-password -e your-email
```

### Daily Auto Report 

**Crontab**

1. You need a server.
2. Adjust your local timezone.
   ```shell
   timedatectl set-timezone "Asia/Shanghai"
   ```
4. Use `crontab -e` and append the following line:
   ```cron
   30 11 * * * curl -L -o /tmp/auto-report https://github.com/hstable/auto-report/releases/latest/download/auto-report_linux_amd64 && chmod +x /tmp/auto-report && /tmp/auto-report -u your-studentID -p your-password -e your-email
   ```
By default, this program will run at 11:30 a.m. everyday.

**Github Action**

1. Fork this repository from here.

2. Add the repository secrets `STUDENTID`, `PASSWORD` and `EMAIL`  in your own repository's <a href="../../settings/secrets">Settings-Secrets</a>,  which also can be found by `Settings`-> `Secrets` -> `New repository secret`. (Repository secrets are invisible for others.)

3. Add `report.yml` file to `.github/workflows`, and write the following content to this file:

   ```yaml
   name: Auto Report
   on: 
     schedule:
       - cron: "30 3 * * *"
     push:
       branches: [main]
   
   jobs:
   
     report:
       name: auto report
       runs-on: ubuntu-latest
       steps:
       - uses: actions/checkout@v2
   
       - name: Build
         run: |
           curl -L -o auto-report https://github.com/hstable/auto-report/releases/latest/download/auto-report_linux_amd64
           chmod +x auto-report
           
       - name: Run
         env:
           STUDENTID: ${{ secrets.STUDENTID }}
           PASSWORD: ${{ secrets.PASSWORD }}
           EMAIL: ${{ secrets.EMAIL }}
         run: |
           if [[ -z $EMAIL ]]
           then ./auto-report -u $STUDENTID -p $PASSWORD
           else ./auto-report -u $STUDENTID -p $PASSWORD -e $EMAIL
           fi
   ```
   
   The time cron above is an UTC time, which is 8 hours slower than Beijing time zone.
   
   By default, this program will run at 11:30 a.m. everyday.

## License

auto-report is licensed by an MIT license as can be found in the LICENSE file.

