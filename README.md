## A fast, concurrent CLI tool for bulk-checking website statuses - written in Go

### Description:
Tool used to very quickly check the status of multiple website concurrently. Developed as part of my initiative to learn Go by building projects I find interesting.

### Features:
- Reads URLs from a provided text file
- Processes all URLs concurrently and checks if the website is up or not
- Generates the report for all of the URLs in report.txt
- Automatically handles time-outs in case that websites are frozen

### Usage:
1. Compile binary:
```bash
go build -o url-monitor
```

2.  Call the binary using the filename that you wish to analyse:
```bash
./url-monitor targets.txt 
```

3. The program will output in results.txt the status of the websites you provided:
```
https://www.sample.net/range#train - DOWN
https://facebook.com - UP
http://www.sample.net/?silver=fang&horse=design#bait - UP
https://google.com - UP
https://www.sample.net/wave#scene - DOWN
https://www.sample.net/?rabbit=measure&crack=fuel - UP
https://www.sample.net/trains?lunchroom=fuel - DOWN
https://www.sample.net/?sticks=payment&eye=trains - UP
```

### What I learned:
- Goroutines and Concurency
- Channels - To follow Go's concurrency idiom:  Do not communicate by sharing memory. Instead share memory by communicating.
- WaitGroups & the Orchestrator pattern (To prevent deadlocks)
- HTTP Client timeouts

### Future Improvments:
Learn about Worker Pools and implement a worker pool in this project so that I don't DDoS my router again by running this with 100k links